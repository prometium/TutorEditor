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
          v-model="selectedScriptUid"
          :items="scriptsInfo"
          :item-text="convertScriptInfoToText"
          :item-value="convertScriptInfoToValue"
          label="Присоединенный фрагмент"
        />
        <v-select :items="frameNames" label="Кадр начала ветвления" dense />
        <v-select
          :items="frameOfSelectedScriptNames"
          label="Первый присоединенный кадр"
          dense
        />
        <v-select
          :items="frameOfSelectedScriptNames"
          label="Последний присоединенный кадр"
          dense
        />
        <v-select :items="frameNames" label="Кадр окончания ветвления" dense />
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
import { PathItem, ScriptInfo, Script, Frame } from "@/common/types";
import { getScript } from "@/common/requests";
import { configurePath } from "@/helpers/configurePath";

export default Vue.extend({
  name: "AddBranchingDialog",
  props: ["value"],
  data() {
    return {
      selectedScriptUid: "",
      selectedScript: null as Script | null
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
      return this.path.map((pathItem: PathItem, index: number) => {
        const text = this.script.frameByUid[pathItem.frameUid]?.hintText || "";
        return `${index}. ${text}`;
      });
    },
    frameOfSelectedScriptNames(): string[] {
      if (!this.selectedScript) return [];

      const frameByUid: Record<string, Frame> = {};
      this.selectedScript.frames.forEach(frame => {
        frameByUid[frame.uid] = frame;
      });

      const path = configurePath(this.selectedScript.firstFrame, frameByUid);

      return path.map((pathItem: PathItem, index: number) => {
        const text = frameByUid[pathItem.frameUid]?.hintText || "";
        return `${index}. ${text}`;
      });
    }
  },
  methods: {
    ...mapActions({
      loadScript: ActionTypes.LOAD_SCRIPT,
      loadScriptsInfo: ActionTypes.LOAD_SCRIPTS_INFO
    }),
    convertScriptInfoToText(scriptInfo: ScriptInfo) {
      return scriptInfo.name;
    },
    convertScriptInfoToValue(scriptInfo: ScriptInfo) {
      return scriptInfo.uid;
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
    },
    selectedScriptUid(value) {
      getScript(value).then(response => {
        this.selectedScript = response.script;
      });
    }
  }
});
</script>
