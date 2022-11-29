<template>
  <div class="frame" tabindex="0" @blur="handleBlur">
    <div class="frame__img-wrapper">
      <img
        v-if="currentFrame && currentFrame.pictureLink"
        ref="img"
        :src="pictureLink"
        :alt="`Кадр ${currentFrame.uid}`"
        class="frame__img"
        @load="handleImgLoad"
      />
      <div
        v-show="showDragMoveArea"
        ref="resizeDragRef"
        class="resize-drag-area"
      />
      <div
        v-for="limit in 2"
        v-show="showDraggingPath"
        ref="draggingPathLimitRefs"
        :key="'dragging-path-limit' + limit"
        :class="`dragging-path-limit-${limit}`"
        @mouseover="activeSwitchPicture = null"
        @mouseleave="activeSwitchPicture = null"
      />
      <div
        v-for="(draggingPathItem, index) in draggingPath"
        :key="'dragging-path-line' + index"
        :style="draggingPathItem.style"
        class="dragging-path-line"
        @mouseover="activeSwitchPicture = draggingPathItem"
        @mouseleave="activeSwitchPicture = null"
      />
      <ImageUploader />
    </div>
  </div>
</template>

<script lang="ts">
import { mapActions, mapState } from "pinia";
import { useStore } from "@/store";
import interact from "interactjs";
import { ActionGroup, ActionType } from "@/common/constants";
import ImageUploader from "./ImageUploader.vue";
import type { SwitchPicture } from "@/common/types";

export default {
  name: "FrameImg",
  components: {
    ImageUploader,
  },
  data() {
    return {
      scale: 1,
      activeSwitchPicture: null as Record<string, unknown> | null,
      timerId: undefined as ReturnType<typeof setTimeout> | undefined,
    };
  },
  mounted() {
    const resizeDragRef = this.$refs.resizeDragRef as HTMLElement;
    this.initInteract(resizeDragRef);
    window.addEventListener("resize", this.onResize);

    const draggingPathLimitRefs = this.$refs.draggingPathLimitRefs as
      | HTMLElement[];
    draggingPathLimitRefs.forEach((ref) => {
      this.initInteract(ref);
      window.addEventListener("resize", this.onResize);
    });
  },
  beforeUnmount() {
    clearTimeout(this.timerId);
    window.removeEventListener("resize", this.onResize);
  },
  computed: {
    ...mapState(useStore, [
      "currentFrame",
      "currentAction",
      "currentActionGroup",
      "script",
      "scriptsInfo",
    ]),
    showDragMoveArea(): boolean {
      return this.currentActionGroup === ActionGroup.Mouse;
    },
    showDraggingPath(): boolean {
      return this.currentAction?.actionType === ActionType.Drag;
    },
    pictureLink(): string {
      return `${import.meta.env.VITE_S3_URL || ""}/${
        import.meta.env.VITE_S3_BUCKET_NAME || "editor"
      }/${
        this.activeSwitchPicture
          ? this.activeSwitchPicture.pictureLink
          : this.currentFrame?.pictureLink
      }`;
    },
    draggingPath(): Record<string, any>[] | null {
      if (this.currentAction?.actionType === ActionType.Drag) {
        return (this.currentAction.switchPictures || []).map(
          (item: SwitchPicture, index: number, array: SwitchPicture[]) => {
            const nextItem = array[index + 1];
            return {
              style: {
                left: item.x * this.scale + "px",
                top: item.y * this.scale + "px",
                width: nextItem
                  ? Math.abs(nextItem.x - item.x) * this.scale + "px"
                  : undefined,
                height: nextItem
                  ? Math.abs(nextItem.y - item.y) * this.scale + "px"
                  : undefined,
                transform: nextItem
                  ? "rotate(" +
                    (Math.atan2(nextItem.y - item.y, nextItem.x - item.x) *
                      180) /
                      Math.PI +
                    "deg)"
                  : undefined,
              },
              pictureLink: item.pictureLink,
            };
          }
        );
      }
      return null;
    },
  },
  watch: {
    currentAction: {
      handler(action) {
        // Макрозадача (таймер), чтобы успели инициализоваться рефы
        this.timerId = setTimeout(() => {
          const resizeDragRef = this.$refs.resizeDragRef as HTMLDivElement;
          const draggingPathLimitRefs = this.$refs
            .draggingPathLimitRefs as HTMLElement[];

          [
            {
              left: action.xLeft,
              top: action.yLeft,
              width: Math.abs(action.xRight - action.xLeft),
              height: Math.abs(action.yRight - action.yLeft),
              ref: resizeDragRef,
            },
            {
              left: action.startXLeft,
              top: action.startYLeft,
              width: Math.abs(action.startXRight - action.startXLeft),
              height: Math.abs(action.startYRight - action.startYLeft),
              ref: draggingPathLimitRefs[0],
            },
            {
              left: action.finishXLeft,
              top: action.finishYLeft,
              width: Math.abs(action.finishXRight - action.finishXLeft),
              height: Math.abs(action.finishYRight - action.finishYLeft),
              ref: draggingPathLimitRefs[1],
            },
          ].forEach((item) => {
            item.ref.setAttribute("data-fixed-x", item.left);
            item.ref.setAttribute("data-fixed-y", item.top);
            item.ref.setAttribute("data-fixed-width", String(item.width));
            item.ref.setAttribute("data-fixed-height", String(item.height));
          });

          this.onResize();

          this.activeSwitchPicture = null;
        }, 0);
      },
      immediate: true,
    },
  },
  methods: {
    ...mapActions(useStore, ["updateScript"]),
    onResize() {
      const img = this.$refs.img as HTMLImageElement | undefined;

      if (img) {
        this.scale = img.clientWidth / img.naturalWidth || 1;
        const resizeDragRef = this.$refs.resizeDragRef as HTMLDivElement;
        const draggingPathLimitRefs = this.$refs
          .draggingPathLimitRefs as HTMLElement[];

        [resizeDragRef, ...draggingPathLimitRefs].forEach((ref) => {
          ref.setAttribute("data-scale", String(this.scale));
          const x = parseFloat(ref.getAttribute("data-fixed-x") || "") || 0;
          const y = parseFloat(ref.getAttribute("data-fixed-y") || "") || 0;
          const width =
            parseInt(ref.getAttribute("data-fixed-width") || "") || 0;
          const height =
            parseInt(ref.getAttribute("data-fixed-height") || "") || 0;

          const scaledX = x * this.scale;
          const scaledY = y * this.scale;

          ref.style.transform = `translate(${scaledX}px, ${scaledY}px)`;
          ref.style.width = `${width * this.scale}px`;
          ref.style.height = `${height * this.scale}px`;

          ref.setAttribute("data-x", String(scaledX));
          ref.setAttribute("data-y", String(scaledY));
        });
      }
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
          },
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
            },
          },
          modifiers: [
            // keep the edges inside the parent
            interact.modifiers.restrictEdges({
              outer: "parent",
            }),

            // minimum size
            interact.modifiers.restrictSize({
              min: { width: 5, height: 5 },
            }),
          ],

          inertia: true,
        })
        .draggable({
          listeners: { move: this.dragMoveListener },
          inertia: true,
          modifiers: [
            interact.modifiers.restrictRect({
              restriction: "parent",
              endOnly: true,
            }),
          ],
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
      if (this.showDragMoveArea) {
        const resizeDragRef = this.$refs.resizeDragRef as HTMLDivElement;

        const x =
          parseFloat(resizeDragRef.getAttribute("data-fixed-x") || "") || 0;
        const y =
          parseFloat(resizeDragRef.getAttribute("data-fixed-y") || "") || 0;
        const scale =
          parseFloat(resizeDragRef.getAttribute("data-scale") || "") || 1;
        const rect = resizeDragRef.getBoundingClientRect();

        this.updateScript({
          frames: [
            {
              uid: this.currentFrame?.uid || "",
              actions: [
                {
                  uid: this.currentAction?.uid || "",
                  xLeft: x,
                  xRight: x + rect.width / scale,
                  yLeft: y,
                  yRight: y + rect.height / scale,
                },
              ],
            },
          ],
        });
      }

      if (this.showDraggingPath) {
        const draggingPathLimitRefs = this.$refs
          .draggingPathLimitRefs as HTMLElement[];

        const limitItems = draggingPathLimitRefs.map((ref) => ({
          x: parseFloat(ref.getAttribute("data-fixed-x") || "") || 0,
          y: parseFloat(ref.getAttribute("data-fixed-y") || "") || 0,
          scale: parseFloat(ref.getAttribute("data-scale") || "") || 1,
          rect: ref.getBoundingClientRect(),
        }));

        this.updateScript({
          frames: [
            {
              uid: this.currentFrame?.uid || "",
              actions: [
                {
                  uid: this.currentAction?.uid || "",
                  startXLeft: limitItems[0].x,
                  startYLeft: limitItems[0].y,
                  startXRight:
                    limitItems[0].x +
                    limitItems[0].rect.width / limitItems[0].scale,
                  startYRight:
                    limitItems[0].y +
                    limitItems[0].rect.height / limitItems[0].scale,
                  finishXLeft: limitItems[1].x,
                  finishYLeft: limitItems[1].y,
                  finishXRight:
                    limitItems[1].x +
                    limitItems[1].rect.width / limitItems[1].scale,
                  finishYRight:
                    limitItems[1].y +
                    limitItems[1].rect.height / limitItems[1].scale,
                },
              ],
            },
          ],
        });
      }
    },
    handleImgLoad() {
      const img = this.$refs.img as HTMLImageElement;
      this.scale = img.clientWidth / img.naturalWidth || 1;
    },
  },
};
</script>

<style scoped lang="scss">
.frame {
  display: flex;
  justify-content: center;
  overflow-y: auto;
  outline: none;
}

.frame__img-wrapper {
  margin: auto;
  position: relative;
  flex: 1;
}

.frame__img {
  max-height: 100%;
  max-width: 100%;
  user-select: none;
}

@keyframes shine {
  0% {
    background-color: rgba(244, 67, 54, 70%);
  }

  10% {
    background-color: rgba(233, 30, 99, 70%);
  }

  20% {
    background-color: rgba(156, 39, 176, 70%);
  }

  30% {
    background-color: rgba(103, 58, 183, 70%);
  }

  40% {
    background-color: rgba(63, 81, 181, 70%);
  }

  50% {
    background-color: rgba(33, 150, 243, 70%);
  }

  60% {
    background-color: rgba(0, 150, 136, 70%);
  }

  70% {
    background-color: rgba(76, 175, 80, 70%);
  }

  80% {
    background-color: rgba(255, 235, 59, 70%);
  }

  90% {
    background-color: rgba(255, 152, 0, 70%);
  }

  100% {
    background-color: rgba(255, 87, 34, 70%);
  }
}

.resize-drag-area {
  position: absolute;
  top: 0;
  left: 0;
  padding: 5px;
  touch-action: none;
  animation: 10s ease-in-out infinite alternate shine;
}

.dragging-path-line {
  position: absolute;
  padding: 1px;
  background-color: rgba(158, 158, 158, 50%);
  cursor: pointer;

  &:hover {
    animation: 10s ease-in-out infinite alternate shine;
  }
}

.dragging-path-limit-1,
.dragging-path-limit-2 {
  position: absolute;
  top: 0;
  left: 0;
  padding: 1px;
  cursor: pointer;
}

.dragging-path-limit-1 {
  animation: 10s ease-in-out infinite alternate shine;
}

.dragging-path-limit-2 {
  animation: 4s ease-in-out infinite alternate shine;
}
</style>
