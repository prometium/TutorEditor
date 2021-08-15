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
        <v-select
          v-model="firstFrameIndex"
          :items="indexedFirstFrames"
          :item-text="convertIndexedFrameToText"
          :item-value="convertIndexedFrameToValue"
          label="Кадр начала ветвления"
          dense
        />
        <v-select
          v-model="firstConnectedFrameIndex"
          :items="indexedFramesOfSelectedScript"
          :item-text="convertIndexedFrameToText"
          :item-value="convertIndexedFrameToValue"
          label="Первый присоединенный кадр"
          :disabled="!this.selectedScriptUid"
          dense
        />
        <v-select
          v-model="lastConnectedFrameIndex"
          :items="indexedFramesOfSelectedScript"
          :item-text="convertIndexedFrameToText"
          :item-value="convertIndexedFrameToValue"
          label="Последний присоединенный кадр"
          :disabled="
            !this.selectedScriptUid || firstConnectedFrameIndex == null
          "
          dense
        />
        <v-select
          v-model="lastFrameIndex"
          :items="indexedLastFrames"
          :item-text="convertIndexedFrameToText"
          :item-value="convertIndexedFrameToValue"
          label="Кадр окончания ветвления"
          :disabled="firstFrameIndex == null"
          dense
        />
      </v-card-text>
      <v-divider />
      <v-card-actions>
        <v-spacer />
        <v-btn @click="dialog = false" text> Отменить </v-btn>
        <v-btn
          @click="handleAdd"
          text
          color="primary"
          :disabled="isAddDisabled"
        >
          Добавить
        </v-btn>
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

type IndexedFrame = Frame & { index: number };

export default Vue.extend({
  name: "AddBranchingDialog",
  props: ["value"],
  data() {
    return {
      selectedScriptUid: null as string | null,
      selectedScript: null as Script | null,
      firstFrameIndex: null as number | null,
      firstConnectedFrameIndex: null as number | null,
      lastConnectedFrameIndex: null as number | null,
      lastFrameIndex: null as number | null
    };
  },
  computed: {
    ...mapState(["script", "scriptsInfo"]),
    ...mapGetters(["path", "currentPathItemIndex"]),
    dialog: {
      get(): boolean {
        return this.value;
      },
      set(value: boolean) {
        this.$emit("input", value);
      }
    },
    indexedFirstFrames(): IndexedFrame[] {
      return this.path.map((pathItem: PathItem, index: number) => ({
        index,
        ...this.script.frameByUid[pathItem.frameUid]
      }));
    },
    indexedLastFrames(): IndexedFrame[] {
      if (this.firstFrameIndex == null) return this.indexedFirstFrames;
      return this.indexedFirstFrames.slice(this.firstFrameIndex + 1);
    },
    frameOfSelectedScriptByUid(): Record<string, Frame> {
      const frameByUid: Record<string, Frame> = {};
      this.selectedScript?.frames.forEach(frame => {
        frameByUid[frame.uid] = frame;
      });
      return frameByUid;
    },
    pathOfSelectedScript(): PathItem[] {
      if (!this.selectedScript) return [];

      const path = configurePath(
        this.selectedScript.firstFrame,
        this.frameOfSelectedScriptByUid
      );
      return path;
    },
    indexedFramesOfSelectedScript(): IndexedFrame[] {
      return this.pathOfSelectedScript.map(
        (pathItem: PathItem, index: number) => ({
          index,
          ...this.frameOfSelectedScriptByUid[pathItem.frameUid]
        })
      );
    },
    isAddDisabled(): boolean {
      return (
        this.firstFrameIndex == null ||
        this.firstConnectedFrameIndex == null ||
        this.lastConnectedFrameIndex == null ||
        this.lastFrameIndex == null ||
        !this.selectedScript
      );
    }
  },
  methods: {
    ...mapActions({
      loadScript: ActionTypes.LOAD_SCRIPT,
      loadScriptsInfo: ActionTypes.LOAD_SCRIPTS_INFO,
      updateScript: ActionTypes.UPDATE_SCRIPT
    }),
    convertScriptInfoToText(scriptInfo: ScriptInfo) {
      return scriptInfo.name;
    },
    convertScriptInfoToValue(scriptInfo: ScriptInfo) {
      return scriptInfo.uid;
    },
    convertIndexedFrameToText(frame: IndexedFrame) {
      return `${frame.index}. ${frame.hintText || ""}`;
    },
    convertIndexedFrameToValue(frame: IndexedFrame) {
      return frame.index;
    },
    async handleAdd() {
      if (
        this.firstFrameIndex == null ||
        this.firstConnectedFrameIndex == null ||
        this.lastConnectedFrameIndex == null ||
        this.lastFrameIndex == null ||
        !this.selectedScript
      )
        return;

      const preparedLastFrame =
        this.script.frameByUid[this.path[this.lastFrameIndex].frameUid];

      const framesToConnect = this.pathOfSelectedScript
        .slice(this.firstConnectedFrameIndex, this.lastConnectedFrameIndex + 1)
        .map((pathItem: PathItem, index) => {
          const frame = this.frameOfSelectedScriptByUid[pathItem.frameUid];

          const action = frame.actions?.[0] || {
            uid: `_:framesToConnectAction${index}`,
            nextFrame: { uid: "" }
          };

          return {
            ...frame,
            uid: `_:${frame.uid}`,
            actions: frame.actions
              ? [
                  {
                    ...action,
                    uid: `_:${action.uid}`,
                    nextFrame: { uid: `_:${action.nextFrame?.uid}` }
                  },
                  ...frame.actions.slice(1)
                ]
              : frame.actions || [action]
          };
        });

      const nextFrame =
        framesToConnect[framesToConnect.length - 1].actions[0].nextFrame;
      if (nextFrame) {
        nextFrame.uid = preparedLastFrame.uid;
      }

      const firstFrame =
        this.script.frameByUid[this.path[this.firstFrameIndex].frameUid];

      const preparedFirstFrame = {
        ...firstFrame,
        actions: [
          ...firstFrame.actions,
          {
            uid: "_:firstFrameAction",
            nextFrame: { uid: framesToConnect[0].uid }
          }
        ]
      };

      await this.updateScript({
        frames: [preparedFirstFrame, ...framesToConnect, preparedLastFrame]
      }).then(() => {
        this.dialog = false;
      });

      this.selectedScriptUid = null;
      this.selectedScript = null;
      this.firstFrameIndex = null;
      this.firstConnectedFrameIndex = null;
      this.lastConnectedFrameIndex = null;
      this.lastFrameIndex = null;
    }
  },
  watch: {
    dialog(value) {
      if (value) {
        this.loadScriptsInfo();
        this.firstFrameIndex = this.currentPathItemIndex;
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
