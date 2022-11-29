<template>
  <div class="editor">
    <AppBar />
    <v-icon icon="home" />
    <Workspace />
  </div>
</template>

<style scoped lang="scss">
.editor {
  display: flex;
  flex-direction: column;
  height: 100%;
}
</style>

<script lang="ts">
import AppBar from "@/components/AppBar/index.vue";
import Workspace from "@/components/Workspace/index.vue";

export default {
  name: "App",
  components: {
    AppBar,
    Workspace,
  },
};
</script>
