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
            <v-list-item @click="handleRemoveBranch">
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
import { mapState, mapGetters } from "vuex";
import { downloadScriptArchive } from "@/common/requests";
import { Frame, PathItem } from "@/common/types";
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
    ...mapGetters(["currentFrame", "path"]),
    showToolbar(): boolean {
      return !!this.currentFrame;
    }
  },
  methods: {
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
    handleRemoveBranch() {
      if ((this.currentFrame?.actions?.length || 0) <= 1) return;

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

      const currentFramePathIndex = this.path.findIndex(
        (pathItem: PathItem) => pathItem.frameUid === this.currentFrame.uid
      );
      const lastFramePathIndex = this.path.length - 1;

      const frameUidsToRemove = [];
      let currentPathIndex = currentFramePathIndex + 1;
      console.log(countByNextFrameUid);
      for (
        ;
        currentPathIndex < lastFramePathIndex &&
        countByNextFrameUid[this.path[currentPathIndex].frameUid] <= 1;
        currentPathIndex++
      ) {
        frameUidsToRemove.push(this.path[currentPathIndex].frameUid);
      }

      console.log(frameUidsToRemove);
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
