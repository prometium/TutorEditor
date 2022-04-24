<template>
  <v-dialog v-model="dialog" width="600">
    <template v-slot:activator="activator">
      <slot name="activator" v-bind="activator" />
    </template>
    <v-card>
      <v-card-title class="headline lighten-2">
        Удаление обучающих программ
      </v-card-title>
      <v-card-text style="max-height: 300px; overflow-y: auto">
        <v-list-item-group v-model="selectedScriptIds" multiple>
          <v-list-item
            v-for="scriptInfo in scriptsInfo"
            :key="scriptInfo.uid"
            :value="scriptInfo.uid"
            dense
            :disabled="script && scriptInfo.uid === script.uid"
          >
            <template v-slot:default="{ active }">
              <v-list-item-action>
                <v-checkbox :input-value="active" />
              </v-list-item-action>
              <v-list-item-content>
                <v-list-item-title>{{ scriptInfo.name }}</v-list-item-title>
              </v-list-item-content>
            </template>
          </v-list-item>
        </v-list-item-group>
      </v-card-text>
      <v-divider />
      <v-card-actions>
        <v-spacer />
        <v-btn @click="dialog = false" text> Отменить </v-btn>
        <v-btn @click="handleDelete" :loading="isLoading" text color="primary">
          Удалить
        </v-btn>
      </v-card-actions>
    </v-card>
  </v-dialog>
</template>

<script lang="ts">
import Vue from "vue";
import { mapState, mapActions } from "vuex";
import { ActionTypes } from "@/store/action-types";

export default Vue.extend({
  name: "DeleteScriptDialog",
  props: ["value"],
  data() {
    return {
      selectedScriptIds: [],
      isLoading: false
    };
  },
  computed: {
    ...mapState(["scriptsInfo", "script"]),
    dialog: {
      get(): boolean {
        return this.value;
      },
      set(value: boolean) {
        this.$emit("input", value);
      }
    }
  },
  methods: {
    ...mapActions({
      deleteScript: ActionTypes.DELETE_SCRIPT,
      loadScriptsInfo: ActionTypes.LOAD_SCRIPTS_INFO
    }),
    async handleDelete() {
      this.isLoading = true;
      try {
        await Promise.all(
          this.selectedScriptIds.map(id => this.deleteScript(id))
        );
        this.dialog = false;
      } catch (error) {
        console.error(error);
      } finally {
        this.isLoading = false;
      }
    }
  },
  watch: {
    dialog(value) {
      if (value) {
        this.loadScriptsInfo();
      }
    }
  }
});
</script>
