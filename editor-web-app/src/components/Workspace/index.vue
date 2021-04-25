<template>
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
</template>

<script lang="ts">
import Vue from "vue";
import { mapActions, mapState } from "vuex";
import store from "@/store";
import { ActionTypes } from "@/store/action-types";
import Frame from "./Frame.vue";
import FramePreviews from "./FramePreviews.vue";
import FrameTools from "./FrameTools.vue";

export default Vue.extend({
  name: "App",
  components: {
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
  mounted() {
    this.loadScript(this.$route.query.scriptUid);
  },
  computed: {
    ...mapState(["script", "frame"]),
    showSidePanel(): boolean {
      return !!this.script.uid;
    }
  },
  methods: {
    ...mapActions({
      loadScript: ActionTypes.LOAD_SCRIPT
    })
  }
});
</script>

<style scoped lang="scss">
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
