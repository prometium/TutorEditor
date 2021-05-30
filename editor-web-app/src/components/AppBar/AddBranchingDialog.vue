<template>
  <v-dialog v-model="dialog" width="600">
    <template v-slot:activator="activator">
      <slot name="activator" v-bind="activator" />
    </template>
    <v-card>
      <v-card-title class="headline lighten-2">
        Добавление ветвления
      </v-card-title>
      <v-card-text style="max-height: 300px; overflow-y: auto">
        <v-select
          :items="scriptsInfo"
          :item-text="convertScriptInfoToName"
          label="Присоединенный фрагмент"
        />
        <v-select :items="frameNames" label="Кадр начала ветвления" />
        <v-select :items="frameNames" label="Кадр окончания ветвления" />
      </v-card-text>
      <v-divider />
      <v-card-actions>
        <v-spacer />
        <v-btn @click="dialog = false" text> Отменить </v-btn>
        <v-btn @click="handleAdd" text color="primary"> Добавить </v-btn>
      </v-card-actions>
    </v-card>
  </v-dialog>
</template>

<script lang="ts">
import Vue from "vue";
import { mapState, mapActions, mapGetters } from "vuex";
import { ActionTypes } from "@/store/action-types";
import { PathItem, ScriptInfo } from "@/common/types";

export default Vue.extend({
  name: "AddBranchingDialog",
  props: ["value"],
  data() {
    return {
      radioGroup: ""
    };
  },
  computed: {
    ...mapState(["script", "scriptsInfo"]),
    ...mapGetters(["path"]),
    dialog: {
      get(): boolean {
        return this.value;
      },
      set(value) {
        this.$emit("input", value);
      }
    },
    scriptNames(): string[] {
      return this.scriptsInfo.map((scriptInfo: ScriptInfo) => scriptInfo.name);
    },
    frameNames(): string[] {
      return this.path.map(
        (pathItem: PathItem, index: number) =>
          `${index}. ${
            this.script.frameByUid[pathItem.frameUid]?.hintText || ""
          }`
      );
    }
  },
  methods: {
    ...mapActions({
      loadScript: ActionTypes.LOAD_SCRIPT,
      loadScriptsInfo: ActionTypes.LOAD_SCRIPTS_INFO
    }),
    convertScriptInfoToName(scriptInfo: ScriptInfo) {
      return scriptInfo.name;
    },
    handleAdd() {
      console.log("TODO");
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
