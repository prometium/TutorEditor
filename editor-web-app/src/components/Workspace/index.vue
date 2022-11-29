<template>
  <main class="workspace">
    <v-sheet
      v-if="hasScript"
      elevation="2"
      tag="aside"
      class="workspace__aside"
    >
      <FramePreviews class="workspace__aside-previews" />
      <FrameTools />
    </v-sheet>
    <FrameImg v-if="hasScript" class="workspace__frame" />
  </main>
</template>

<script lang="ts">
import { mapActions, mapState } from "pinia";
import { useStore } from "@/store";

import FrameImg from "./FrameImg.vue";
import FramePreviews from "./FramePreviews.vue";
import FrameTools from "./FrameTools.vue";

export default {
  name: "App",
  components: {
    FrameImg,
    FramePreviews,
    FrameTools,
  },
  data() {
    return {
      scriptUid: "",
    };
  },
  mounted() {
    if (this.$route.query.scriptUid) {
      this.loadScript(String(this.$route.query.scriptUid));
    }
  },
  computed: {
    ...mapState(useStore, ["script", "scriptsInfo"]),
    hasScript(): boolean {
      return !!this.script.uid;
    },
  },
  methods: {
    ...mapActions(useStore, ["loadScript"]),
  },
};
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
