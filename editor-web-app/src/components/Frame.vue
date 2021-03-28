<template>
  <div class="frame">
    <div class="frame__img-wrapper">
      <img
        ref="img"
        :src="frame.pictureLink + '#' + new Date().getTime()"
        :alt="frame.uid"
        class="frame__img"
      />
      <div ref="resizeDrag" class="resize-drag" />
    </div>
  </div>
</template>

<script>
import Vue from "vue";
import { mapState } from "vuex";
import interact from "interactjs";

export default Vue.extend({
  name: "Frame",
  data() {
    return {
      scale: 1
    };
  },
  mounted() {
    const resizeDrag = this.$refs.resizeDrag;
    this.initInteract(resizeDrag);

    window.addEventListener("resize", this.onResize);
  },
  beforeDestroy() {
    window.removeEventListener("resize", this.onResize);
  },
  computed: {
    ...mapState(["frame", "script"]),
    currentBranchNum() {
      return this.script.branchNumByUid[this.frame.uid] || 0;
    },
    selectedAction() {
      return this.frame.actions && this.frame.actions[this.currentBranchNum];
    }
  },
  watch: {
    selectedAction(action) {
      this.onResize();
      const img = this.$refs.img;
      const scale = img.clientWidth / img.naturalWidth || 1;

      const resizeDrag = this.$refs.resizeDrag;
      resizeDrag.setAttribute("data-scale", scale);

      resizeDrag.style.webkitTransform = resizeDrag.style.transform = `
      translate(${action.xLeft * scale}px, ${action.yLeft * scale}px)
      `;
      resizeDrag.style.width = `${(action.xRight - action.xLeft) * scale}px`;
      resizeDrag.style.height = `${(action.yRight - action.yLeft) * scale}px`;

      resizeDrag.setAttribute("data-x", action.xLeft * scale);
      resizeDrag.setAttribute("data-y", action.yLeft * scale);
      resizeDrag.setAttribute("data-fixed-x", action.xLeft);
      resizeDrag.setAttribute("data-fixed-y", action.yLeft);
      resizeDrag.setAttribute("data-fixed-width", action.xRight - action.xLeft);
      resizeDrag.setAttribute(
        "data-fixed-height",
        action.yRight - action.yLeft
      );
    }
  },
  methods: {
    onResize() {
      const img = this.$refs.img;
      const scale = img.clientWidth / img.naturalWidth || 1;

      const resizeDrag = this.$refs.resizeDrag;
      resizeDrag.setAttribute("data-scale", scale);

      const x = parseFloat(resizeDrag.getAttribute("data-fixed-x")) || 0;
      const y = parseFloat(resizeDrag.getAttribute("data-fixed-y")) || 0;
      const width = +resizeDrag.getAttribute("data-fixed-width") || 0;
      const height = +resizeDrag.getAttribute("data-fixed-height") || 0;
      console.log(width, width * scale);
      resizeDrag.style.webkitTransform = resizeDrag.style.transform = `
      translate(${x * scale}px, ${y * scale}px)
      `;
      resizeDrag.style.width = `${width * scale}px`;
      resizeDrag.style.height = `${height * scale}px`;

      resizeDrag.setAttribute("data-x", x * scale);
      resizeDrag.setAttribute("data-y", y * scale);
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

              target.style.webkitTransform = target.style.transform = `
              translate(${x}px, ${y}px)
              `;

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
