<template>
  <div class="frame">
    <div ref="resizeDrag" class="resize-drag">
      <img
        :src="frame.pictureLink + '#' + new Date().getTime()"
        :alt="frame.uid"
        class="frame__img"
      />
    </div>
  </div>
</template>

<script>
import Vue from "vue";
import { mapState } from "vuex";
import interact from "interactjs";

export default Vue.extend({
  name: "Frame",
  computed: {
    ...mapState(["frame"])
  },
  mounted() {
    let resizeDrag = this.$refs.resizeDrag;
    this.initInteract(resizeDrag);
  },
  methods: {
    initInteract(selector) {
      interact(selector)
        .resizable({
          // resize from all edges and corners
          edges: { left: true, right: true, bottom: true, top: true },

          listeners: {
            move(event) {
              var target = event.target;
              var x = parseFloat(target.getAttribute("data-x")) || 0;
              var y = parseFloat(target.getAttribute("data-y")) || 0;

              // update the element's style
              target.style.width = event.rect.width + "px";
              target.style.height = event.rect.height + "px";

              // translate when resizing from top or left edges
              x += event.deltaRect.left;
              y += event.deltaRect.top;

              target.style.webkitTransform = target.style.transform =
                "translate(" + x + "px," + y + "px)";

              target.setAttribute("data-x", x);
              target.setAttribute("data-y", y);
              target.textContent =
                Math.round(event.rect.width) +
                "\u00D7" +
                Math.round(event.rect.height);
            }
          },
          modifiers: [
            // keep the edges inside the parent
            interact.modifiers.restrictEdges({
              outer: "parent"
            }),

            // minimum size
            interact.modifiers.restrictSize({
              min: { width: 100, height: 50 }
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

.frame__img {
  max-height: 100%;
  max-width: 100%;
}

.resize-drag {
  width: 120px;
  border-radius: 8px;
  padding: 20px;
  margin: 1rem;
  background-color: #29e;
  color: white;
  font-size: 20px;
  font-family: sans-serif;
  touch-action: none;
}
</style>
