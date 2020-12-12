<template>
  <Header />
  <Toolbar />
  <main class="main">
    <Frame class="frame" />
    <FramePreviews class="frame-previews" />
    <FrameTools class="frame-tools" />
  </main>
</template>

<script lang="ts">
import { defineComponent } from "vue";
import store from "@/store";
import Header from "@/components/Header.vue";
import Toolbar from "@/components/Toolbar.vue";
import Frame from "@/components/Frame.vue";
import FramePreviews from "@/components/FramePreviews.vue";
import FrameTools from "@/components/FrameTools.vue";

export default defineComponent({
  name: "App",
  components: {
    Header,
    Toolbar,
    Frame,
    FramePreviews,
    FrameTools
  },
  store
});
</script>

<style lang="scss">
#app {
  display: flex;
  flex-direction: column;
  height: 100%;
}

.main {
  display: grid;
  grid-template-areas:
    "frame previews"
    "frame tools";
  grid-template-rows: 1fr auto;
  grid-template-columns: 1fr minmax(120px, 12%);
  height: 100%;
  min-height: 0px;
}

.frame {
  grid-area: frame;
}

.frame-previews {
  grid-area: previews;
}

.frame-tools {
  grid-area: tools;
}
</style>
