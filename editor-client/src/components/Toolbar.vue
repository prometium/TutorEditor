<template>
  <div class="toolbar">
    <input v-model="frame.taskText" />
    <input v-model="frame.hintText" />
    <select>
      <option>Действие 1</option>
      <option>Действие 2</option>
    </select>
  </div>
</template>

<script lang="ts">
import { defineComponent } from "vue";
import { mapState } from "vuex";
import store from "@/store";

export default defineComponent({
  name: "Toolbar",
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

<style scoped lang="scss">
.toolbar {
  display: flex;
  padding: 8px 4px;
  border-top: 1px solid var(--palette-divider);
  border-bottom: 1px solid var(--palette-divider);
  background-color: var(--palette-background-paper);
}

input {
  margin: 0 8px;
}
</style>
