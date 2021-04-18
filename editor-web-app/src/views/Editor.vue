<template>
  <div class="editor">
    <AppBar />
    <Workspace />
  </div>
</template>

<script lang="ts">
import Vue from "vue";
import AppBar from "@/components/AppBar/index.vue";
import Workspace from "@/components/Workspace/index.vue";

export default Vue.extend({
  name: "App",
  components: {
    AppBar,
    Workspace
  }
});
</script>

<style scoped lang="scss">
.editor {
  display: flex;
  flex-direction: column;
  height: 100%;
}
</style>
