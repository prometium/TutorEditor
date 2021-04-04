<template>
  <v-app>
    <AppBar />
    <main class="workspace">
      <v-sheet
        v-if="showSidePanel"
        elevation="2"
        tag="aside"
        class="workspace__aside"
      >
        <FramePreviews class="workspace__aside-previews" />
        <FrameTools />
      </v-sheet>
      <Frame class="workspace__frame" />
    </main>
  </v-app>
</template>

<script lang="ts">
import Vue from "vue";
import { mapState } from "vuex";
import store from "@/store";
import AppBar from "@/components/AppBar/index.vue";
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
    ...mapState(["script", "frame"]),
    showSidePanel(): boolean {
      return !!this.script.uid;
    }
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
  display: flex;
  height: 100%;
  min-height: 0px;
}

.workspace__aside {
  display: flex;
  flex-direction: column;
  flex: 1;
  min-width: 144px;
  max-width: 12%;
}

.workspace__aside-previews {
  flex: 1;
}

.workspace__frame {
  flex: 1;
}
</style>
