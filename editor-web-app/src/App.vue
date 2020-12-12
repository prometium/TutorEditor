<template>
  <v-app>
    <AppBar />
    <main class="workspace">
      <Frame class="workspace__frame" />
      <FramePreviews class="workspace__frame-previews" />
      <FrameTools class="workspace__frame-tools" />
    </main>
  </v-app>
</template>

<script lang="ts">
import Vue from "vue";
import { mapState } from "vuex";
import store from "@/store";
import AppBar from "@/components/AppBar.vue";
import Frame from "@/components/Frame.vue";
import FramePreviews from "@/components/FramePreviews.vue";
import FrameTools from "@/components/FrameTools.vue";

export default Vue.extend({
  name: "App",
  components: {
    AppBar,
    Frame,
    FramePreviews,
    FrameTools
  },
  store,
  data() {
    return {
      scriptUid: ""
    };
  },
  computed: {
    ...mapState(["frame"])
  }
});
</script>

<style lang="scss">
html {
  overflow-y: auto;
}

html,
body {
  height: 100%;
}

#app {
  display: flex;
  flex-direction: column;
  height: 100%;
}

.workspace {
  display: grid;
  grid-template-areas:
    "previews frame"
    "tools frame ";
  grid-template-rows: 1fr auto;
  grid-template-columns: minmax(120px, 12%) 1fr;
  height: 100%;
  min-height: 0px;
  padding-top: 4px;
}

.workspace__frame {
  grid-area: frame;
}

.workspace__frame-previews {
  grid-area: previews;
}

.workspace__frame-tools {
  grid-area: tools;
}

.workspace__action {
  grid-area: action;
}
</style>
