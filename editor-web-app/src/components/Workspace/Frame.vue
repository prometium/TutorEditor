<template>
  <div v-if="currentFrame" class="frame" tabindex="0" @blur="handleBlur">
    <div v-show="currentFrame.uid" class="frame__img-wrapper">
      <template v-if="currentFrame.pictureLink">
        <img
          ref="img"
          :src="`${API_ROOT}/images/${
            currentFrame.pictureLink
          }#${new Date().getTime()}`"
          :alt="currentFrame.uid"
          class="frame__img"
        />
        <div v-show="showDragMoveArea" ref="resizeDrag" class="resize-drag" />
      </template>
      <ImageUploader />
    </div>
  </div>
</template>

<script lang="ts">
import Vue from "vue";
import { mapActions, mapGetters } from "vuex";
import interact from "interactjs";
import { ActionGroup } from "@/common/constants";
import { API_ROOT } from "@/common/requests";
import { ActionTypes } from "@/store/action-types";
import ImageUploader from "./ImageUploader.vue";

export default Vue.extend({
  name: "Frame",
  components: {
    ImageUploader
  },
  data() {
    return {
      API_ROOT
    };
  },
  mounted() {
    const resizeDrag = this.$refs.resizeDrag as HTMLElement;
    if (resizeDrag) {
      this.initInteract(resizeDrag);
      window.addEventListener("resize", this.onResize);
    }
  },
  beforeDestroy() {
    window.removeEventListener("resize", this.onResize);
  },
  computed: {
    ...mapGetters(["currentFrame", "currentAction", "currentActionGroup"]),
    showDragMoveArea(): boolean {
      return this.currentActionGroup === ActionGroup.Mouse;
    }
  },
  watch: {
    currentAction(action) {
      const img = this.$refs.img as HTMLImageElement;
      const resizeDrag = this.$refs.resizeDrag as HTMLDivElement;

      if (!img || !resizeDrag || !this.showDragMoveArea) return;

      const width = action.xRight - action.xLeft;
      const height = action.yRight - action.yLeft;

      resizeDrag.setAttribute("data-fixed-x", action.xLeft);
      resizeDrag.setAttribute("data-fixed-y", action.yLeft);
      resizeDrag.setAttribute("data-fixed-width", String(width));
      resizeDrag.setAttribute("data-fixed-height", String(height));

      setTimeout(() => {
        this.onResize();
      }, 100);
    }
  },
  methods: {
    ...mapActions({
      updateFrames: ActionTypes.UPDATE_SCRIPT
    }),
    onResize() {
      const img = this.$refs.img as HTMLImageElement;
      const resizeDrag = this.$refs.resizeDrag as HTMLDivElement;

      const scale = img.clientWidth / img.naturalWidth || 1;

      resizeDrag.setAttribute("data-scale", String(scale));
      const x = parseFloat(resizeDrag.getAttribute("data-fixed-x") || "") || 0;
      const y = parseFloat(resizeDrag.getAttribute("data-fixed-y") || "") || 0;
      const width =
        parseInt(resizeDrag.getAttribute("data-fixed-width") || "") || 0;
      const height =
        parseInt(resizeDrag.getAttribute("data-fixed-height") || "") || 0;

      const scaledX = x * scale;
      const scaledY = y * scale;

      resizeDrag.style.transform = `translate(${scaledX}px, ${scaledY}px)`;
      resizeDrag.style.width = `${width * scale}px`;
      resizeDrag.style.height = `${height * scale}px`;

      resizeDrag.setAttribute("data-x", String(scaledX));
      resizeDrag.setAttribute("data-y", String(scaledY));
    },
    initInteract(selector: HTMLElement) {
      interact(selector)
        .draggable({
          onend(event) {
            const target = event.target;
            const x = parseFloat(target.getAttribute("data-x")) || 0;
            const y = parseFloat(target.getAttribute("data-y")) || 0;
            const scale = parseFloat(target.getAttribute("data-scale")) || 1;

            target.setAttribute("data-fixed-x", x / scale);
            target.setAttribute("data-fixed-y", y / scale);
          }
        })
        .resizable({
          // resize from all edges and corners
          edges: { left: true, right: true, bottom: true, top: true },

          listeners: {
            move(event) {
              const target = event.target;
              let x = parseFloat(target.getAttribute("data-x")) || 0;
              let y = parseFloat(target.getAttribute("data-y")) || 0;
              const scale = parseFloat(target.getAttribute("data-scale")) || 1;

              // update the element's style
              target.style.width = `${event.rect.width}px`;
              target.style.height = `${event.rect.height}px`;

              // translate when resizing from top or left edges
              x += event.deltaRect.left;
              y += event.deltaRect.top;

              target.style.transform = `translate(${x}px, ${y}px)`;

              target.setAttribute("data-x", x);
              target.setAttribute("data-y", y);
              target.setAttribute("data-fixed-x", x / scale);
              target.setAttribute("data-fixed-y", y / scale);
              target.setAttribute("data-fixed-width", event.rect.width / scale);
              target.setAttribute(
                "data-fixed-height",
                event.rect.height / scale
              );
            }
          },
          modifiers: [
            // keep the edges inside the parent
            interact.modifiers.restrictEdges({
              outer: "parent"
            }),

            // minimum size
            interact.modifiers.restrictSize({
              min: { width: 5, height: 5 }
            })
          ],

          inertia: true
        })
        .draggable({
          listeners: { move: this.dragMoveListener },
          inertia: true,
          modifiers: [
            interact.modifiers.restrictRect({
              restriction: "parent",
              endOnly: true
            })
          ]
        });
    },
    dragMoveListener(event: DragEvent & { dx: number; dy: number }) {
      const target = event.target as HTMLDivElement;

      // keep the dragged position in the data-x/data-y attributes
      const x =
        (parseFloat(target.getAttribute("data-x") || "") || 0) + event.dx;
      const y =
        (parseFloat(target.getAttribute("data-y") || "") || 0) + event.dy;

      // translate the element
      target.style.webkitTransform = target.style.transform =
        "translate(" + x + "px, " + y + "px)";

      // update the posiion attributes
      target.setAttribute("data-x", String(x));
      target.setAttribute("data-y", String(y));
    },
    handleBlur() {
      const resizeDrag = this.$refs.resizeDrag as HTMLDivElement;
      if (!resizeDrag) return;

      const x = parseFloat(resizeDrag.getAttribute("data-fixed-x") || "") || 0;
      const y = parseFloat(resizeDrag.getAttribute("data-fixed-y") || "") || 0;
      const scale =
        parseFloat(resizeDrag.getAttribute("data-scale") || "") || 1;

      const rect = resizeDrag.getBoundingClientRect();

      this.updateFrames({
        frames: [
          {
            uid: this.currentFrame.uid,
            actions: [
              {
                uid: this.currentAction.uid,
                xLeft: x,
                xRight: x + rect.width / scale,
                yLeft: y,
                yRight: y + rect.height / scale
              }
            ]
          }
        ]
      });
    }
  }
});
</script>

<style scoped lang="scss">
.frame {
  display: flex;
  align-items: center;
  justify-content: center;
  overflow-y: auto;
  outline: none;
}

.frame__img-wrapper {
  position: relative;
  flex: 1;
}

.frame__img {
  max-height: 100%;
  max-width: 100%;
}

@keyframes shine {
  from {
    background-color: rgba(33, 33, 33, 70%);
  }
  to {
    background-color: rgba(207, 216, 220, 70%);
  }
}

.resize-drag {
  position: absolute;
  top: 0;
  left: 0;
  padding: 5px;
  touch-action: none;
  animation: 4s ease-in-out infinite alternate shine;
}
</style>
