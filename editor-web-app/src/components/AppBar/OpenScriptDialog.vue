<template>
  <v-dialog v-model="dialog" width="600">
    <template v-slot:activator="activator">
      <slot name="activator" v-bind="activator" />
    </template>
    <v-card>
      <v-card-title class="headline lighten-2">
        Обучающие программы
      </v-card-title>
      <v-card-text style="max-height: 300px; overflow-y: auto">
        <v-radio-group v-model="selectedScriptId" column>
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
        <v-btn @click="dialog = false" text> Отменить </v-btn>
        <v-btn @click="handleOpen" :loading="isLoading" text color="primary">
          Открыть
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
  name: "OpenScriptDialog",
  props: ["value"],
  data() {
    return {
      selectedScriptId: "",
      isLoading: false
    };
  },
  computed: {
    ...mapState(["scriptsInfo"]),
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
      loadScript: ActionTypes.LOAD_SCRIPT,
      loadScriptsInfo: ActionTypes.LOAD_SCRIPTS_INFO
    }),
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
