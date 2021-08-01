<template>
  <v-sheet tag="header" class="header" elevation="4">
    <div class="menubar-container">
      <div class="script-title text-h6">{{ script.name || "..." }}</div>
      <div class="menubar">
        <v-menu>
          <template v-slot:activator="{ on, attrs }">
            <v-btn v-bind="attrs" v-on="on" small text elevation="0">
              Файл
            </v-btn>
          </template>
          <v-list>
            <v-list-item @click="createScriptDialog = true">
              <v-list-item-title>Создать</v-list-item-title>
            </v-list-item>
            <v-list-item @click="openScriptDialog = true">
              <v-list-item-title>Открыть</v-list-item-title>
            </v-list-item>
            <v-list-item @click="handleDownloadScriptArchive">
              <v-list-item-title>Скачать архив</v-list-item-title>
            </v-list-item>
          </v-list>
        </v-menu>
        <CreateScriptDialog v-model="createScriptDialog" />
        <OpenScriptDialog v-model="openScriptDialog" />
        <v-menu>
          <template v-slot:activator="{ on, attrs }">
            <v-btn v-bind="attrs" v-on="on" small text elevation="0">
              Редактирование
            </v-btn>
          </template>
          <v-list>
            <v-list-item @click="addBranchingDialog = true">
              <v-list-item-title>Добавить ветвление</v-list-item-title>
            </v-list-item>
            <v-list-item
              @click="handleRemoveBranch"
              :disabled="isBranchRemovingDisabled"
            >
              <v-list-item-title>Удалить ветку</v-list-item-title>
            </v-list-item>
          </v-list>
        </v-menu>
        <AddBranchingDialog v-model="addBranchingDialog" />
      </div>
      <v-btn icon large class="user-button">
        <v-icon>mdi-account-circle</v-icon>
      </v-btn>
    </div>
    <Toolbar v-if="showToolbar" />
  </v-sheet>
</template>

<script lang="ts">
import Vue from "vue";
import { mapState, mapGetters, mapActions, mapMutations } from "vuex";
import { downloadScriptArchive } from "@/common/requests";
import { Frame } from "@/common/types";
import { ActionTypes } from "@/store/action-types";
import { MutationTypes } from "@/store/mutation-types";
import OpenScriptDialog from "./OpenScriptDialog.vue";
import CreateScriptDialog from "./CreateScriptDialog.vue";
import AddBranchingDialog from "./AddBranchingDialog.vue";
import Toolbar from "./Toolbar/index.vue";

export default Vue.extend({
  name: "AppBar",
  props: ["value"],
  data() {
    return {
      createScriptDialog: false,
      openScriptDialog: false,
      addBranchingDialog: false
    };
  },
  components: {
    OpenScriptDialog,
    CreateScriptDialog,
    AddBranchingDialog,
    Toolbar
  },
  computed: {
    ...mapState(["script"]),
    ...mapGetters([
      "currentFrame",
      "path",
      "currentAction",
      "currentPathItem",
      "currentPathItemIndex"
    ]),
    showToolbar(): boolean {
      return !!this.currentFrame;
    },
    isBranchRemovingDisabled(): boolean {
      return (this.currentFrame?.actions?.length || 0) <= 1;
    }
  },
  methods: {
    ...mapActions({
      updateScript: ActionTypes.UPDATE_SCRIPT
    }),
    ...mapMutations({
      configurePath: MutationTypes.CONFIGURE_PATH
    }),
    handleDownloadScriptArchive() {
      downloadScriptArchive(this.script.uid).then(result => {
        const url = window.URL.createObjectURL(result);
        const link = document.createElement("a");
        link.href = url;
        link.download = "Script";
        link.click();
        window.URL.revokeObjectURL(url);
      });
    },
    async handleRemoveBranch() {
      const countByNextFrameUid = {} as Record<string, number>;
      Object.values(this.script.frameByUid as Record<string, Frame>).forEach(
        frame => {
          frame?.actions?.forEach(action => {
            if (action.nextFrame) {
              countByNextFrameUid[action.nextFrame.uid] =
                (countByNextFrameUid[action.nextFrame.uid] || 0) + 1;
            }
          });
        }
      );

      const lastFramePathItemIndex = this.path.length - 1;

      const frameIdsToDel = [];
      let pathItemIndex = this.currentPathItemIndex + 1;
      for (
        ;
        pathItemIndex < lastFramePathItemIndex &&
        countByNextFrameUid[this.path[pathItemIndex].frameUid] <= 1;
        pathItemIndex++
      ) {
        frameIdsToDel.push(this.path[pathItemIndex].frameUid);
      }

      const actionIdsToDel = [this.currentAction.uid];

      this.configurePath({
        frameUid: this.currentPathItem.frameUid,
        branchNum: Math.max(this.currentPathItem.branchNum - 1, 0)
      });

      this.updateScript({
        actionIdsToDel,
        frameIdsToDel
      });
    }
  }
});
</script>

<style scoped lang="scss">
.header {
  position: relative;
  padding: 4px 16px;
}

.menubar-container {
  display: grid;
  grid-template-areas:
    "script-title user-button"
    "menubar user-button";
  padding-bottom: 8px;
}

.script-title {
  grid-area: script-title;
  padding-left: 12px;
}

.menubar {
  grid-area: menubar;
}

.user-button {
  grid-area: user-button;
  align-self: center;
  justify-self: end;
}

.toolbar {
  display: flex;
  padding: 4px 12px 0;
}

.toolbar__container {
  padding-left: 16px;
  padding-top: 0;
  padding-bottom: 0;
}
</style>
