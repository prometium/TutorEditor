<template>
  <div v-show="showDragMoveArea" ref="resizeDrag" class="resize-drag" />
</template>

<script>
import Vue from "vue";
import { mapGetters, mapState } from "vuex";
import interact from "interactjs";
import { ActionGroup } from "@/common/constants";

export default Vue.extend({
  name: "ResizeDrag",
  mounted() {
    const resizeDrag = this.$refs.resizeDrag;
    if (resizeDrag) {
      this.initInteract(resizeDrag);
    }

    window.addEventListener("resize", this.onResize);
  },
  beforeDestroy() {
    window.removeEventListener("resize", this.onResize);
  },
  computed: {
    ...mapState(["script"]),
    ...mapGetters(["frame", "selectedAction", "selectedActionGroup"]),
    showDragMoveArea() {
      return this.selectedActionGroup === ActionGroup.Mouse;
    }
  },
  watch: {
    selectedAction(action) {
      const img = this.$refs.img;
      const resizeDrag = this.$refs.resizeDrag;

      if (!img || !resizeDrag || !this.showDragMoveArea) return;

      const width = action.xRight - action.xLeft;
      const height = action.yRight - action.yLeft;

      resizeDrag.setAttribute("data-fixed-x", action.xLeft);
      resizeDrag.setAttribute("data-fixed-y", action.yLeft);
      resizeDrag.setAttribute("data-fixed-width", width);
      resizeDrag.setAttribute("data-fixed-height", height);

      setTimeout(() => {
        this.onResize();
      }, 100);
    }
  },
  methods: {
    onResize() {
      const img = this.$refs.img;
      const resizeDrag = this.$refs.resizeDrag;
      const scale = img.clientWidth / img.naturalWidth || 1;

      resizeDrag.setAttribute("data-scale", scale);
      const x = parseFloat(resizeDrag.getAttribute("data-fixed-x")) || 0;
      const y = parseFloat(resizeDrag.getAttribute("data-fixed-y")) || 0;
      const width = resizeDrag.getAttribute("data-fixed-width") || 0;
      const height = resizeDrag.getAttribute("data-fixed-height") || 0;

      const scaledX = x * scale;
      const scaledY = y * scale;

      resizeDrag.style.transform = `translate(${scaledX}px, ${scaledY}px)`;
      resizeDrag.style.width = `${width * scale}px`;
      resizeDrag.style.height = `${height * scale}px`;

      resizeDrag.setAttribute("data-x", scaledX);
      resizeDrag.setAttribute("data-y", scaledY);
    },
    initInteract(selector) {
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
    dragMoveListener(event) {
      var target = event.target;
      // keep the dragged position in the data-x/data-y attributes
      var x = (parseFloat(target.getAttribute("data-x")) || 0) + event.dx;
      var y = (parseFloat(target.getAttribute("data-y")) || 0) + event.dy;

      // translate the element
      target.style.webkitTransform = target.style.transform =
        "translate(" + x + "px, " + y + "px)";

      // update the posiion attributes
      target.setAttribute("data-x", x);
      target.setAttribute("data-y", y);
    },
    handleBlur() {
      console.log("blur");
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
