<template>
  <v-dialog v-model="dialog" width="600">
    <template v-slot:activator="activator">
      <slot name="activator" v-bind="activator" />
    </template>
    <v-card>
      <v-card-title class="text-h5"> Удаление обучающих программ </v-card-title>
      <v-card-text style="max-height: 300px; overflow-y: auto">
        <v-list v-model="selectedScriptIds" multiple>
          <v-list-item
            v-for="scriptInfo in scriptsInfo"
            :key="scriptInfo.uid"
            :value="scriptInfo.uid"
            dense
            :disabled="script && scriptInfo.uid === script.uid"
          >
            <template v-slot:default="{ active }">
              <v-list-item-action>
                <v-checkbox :model-value="active" />
              </v-list-item-action>
              <v-list-item-title>{{ scriptInfo.name }}</v-list-item-title>
            </template>
          </v-list-item>
        </v-list>
      </v-card-text>
      <v-divider />
      <v-card-actions>
        <v-spacer />
        <v-btn @click="dialog = false" variant="text"> Отменить </v-btn>
        <v-btn
          @click="handleDelete"
          :loading="isLoading"
          variant="text"
          color="primary"
        >
          Удалить
        </v-btn>
      </v-card-actions>
    </v-card>
  </v-dialog>
</template>

<script lang="ts">
import { mapState, mapActions } from "pinia";
import { useStore } from "@/store";

export default {
  name: "DeleteScriptDialog",
  props: ["modelValue"],
  data() {
    return {
      selectedScriptIds: [],
      isLoading: false,
    };
  },
  computed: {
    ...mapState(useStore, ["scriptsInfo", "script"]),
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
    ...mapActions(useStore, ["deleteScript", "loadScriptsInfo"]),
    async handleDelete() {
      this.isLoading = true;
      try {
        await Promise.all(
          this.selectedScriptIds.map((id) => this.deleteScript(id))
        );
        this.dialog = false;
      } catch (error) {
        console.error(error);
      } finally {
        this.isLoading = false;
      }
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
