<template>
  <v-sheet tag="header" class="header" elevation="4">
    <div class="menubar-container">
      <div class="script-title text-h6">{{ script.name || "..." }}</div>
      <div class="menubar">
        <v-menu close-on-click>
          <template v-slot:activator="{ on, attrs }">
            <v-btn v-bind="attrs" v-on="on" small text elevation="0">
              Файл
            </v-btn>
          </template>
          <v-list>
            <CreateScriptDialogButton />
            <OpenScriptDialogButton />
          </v-list>
        </v-menu>
        <v-btn small text elevation="0"> Редактирование </v-btn>
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
import { mapState } from "vuex";
import OpenScriptDialogButton from "./OpenScriptDialogButton.vue";
import CreateScriptDialogButton from "./CreateScriptDialogButton.vue";
import Toolbar from "./Toolbar.vue";

export default Vue.extend({
  name: "AppBar",
  components: {
    OpenScriptDialogButton,
    CreateScriptDialogButton,
    Toolbar
  },
  computed: {
    ...mapState(["script", "frame"]),
    showToolbar(): boolean {
      return !!this.frame;
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
