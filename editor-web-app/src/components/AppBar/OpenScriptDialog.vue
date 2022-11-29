<template>
  <v-dialog v-model="dialog" width="600">
    <template v-slot:activator="activator">
      <slot name="activator" v-bind="activator" />
    </template>
    <v-card>
      <v-card-title class="text-h5"> Обучающие программы </v-card-title>
      <v-card-text style="max-height: 300px; overflow-y: auto">
        <v-radio-group v-model="selectedScriptId">
          <v-radio
            v-for="scriptInfo in scriptsInfo"
            :key="scriptInfo.uid"
            :label="scriptInfo.name"
            :value="scriptInfo.uid"
          />
        </v-radio-group>
      </v-card-text>
      <v-divider />
      <v-card-actions>
        <v-spacer />
        <v-btn @click="dialog = false"> Отменить </v-btn>
        <v-btn @click="handleOpen" :loading="isLoading" color="primary">
          Открыть
        </v-btn>
      </v-card-actions>
    </v-card>
  </v-dialog>
</template>

<script lang="ts">
import { mapState, mapActions } from "pinia";
import { useStore } from "@/store";

export default {
  name: "OpenScriptDialog",
  props: ["modelValue"],
  data() {
    return {
      selectedScriptId: "",
      isLoading: false,
    };
  },
  computed: {
    ...mapState(useStore, ["scriptsInfo", "script", "scriptsInfo"]),
    dialog: {
      get(): boolean {
        return this.modelValue;
      },
      set(value: boolean) {
        this.$emit("update:modelValue", value);
      },
    },
  },
  methods: {
    ...mapActions(useStore, ["loadScript", "loadScriptsInfo"]),
    handleOpen() {
      this.isLoading = true;
      this.loadScript(this.selectedScriptId)
        .then(() => {
          this.dialog = false;
          this.$router
            .push({ path: "/", query: { scriptUid: this.selectedScriptId } })
            .catch(() => {
              /* ignore */
            });
        })
        .catch(console.error)
        .finally(() => {
          this.isLoading = false;
        });
    },
  },
  watch: {
    dialog(value) {
      if (value) {
        this.loadScriptsInfo();
      }
    },
  },
};
</script>
