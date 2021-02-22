<template>
  <v-sheet tag="header" class="header" elevation="4">
    <div class="menubar-container">
      <div class="script-title text-h6">{{ script.name || "..." }}</div>
      <div class="menubar">
        <v-menu>
          <template v-slot:activator="{ on, attrs }">
            <v-btn v-bind="attrs" v-on="{ ...on }" small text elevation="0">
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
    <div class="toolbar">
      <v-btn @click="toggleExpansion" elevation="1" icon
        ><v-icon>{{
          expanded ? "mdi-chevron-up" : "mdi-chevron-down"
        }}</v-icon></v-btn
      >
      <v-container fluid class="toolbar__container">
        <v-row align="center">
          <v-col class="d-flex" cols="12" md="6">
            <v-text-field
              v-model="frame.taskText"
              label="Текст задания"
              dense
            ></v-text-field>
          </v-col>
          <v-col class="d-flex" cols="12" md="6">
            <v-text-field
              v-model="frame.hintText"
              label="Текст подсказки"
              dense
            ></v-text-field>
          </v-col>
        </v-row>
        <v-row v-show="expanded">
          <v-col class="d-flex">
            <v-select :items="[]" label="Действие" dense></v-select>
          </v-col>
        </v-row>
      </v-container>
    </div>
  </v-sheet>
</template>

<script lang="ts">
import Vue from "vue";
import { mapState } from "vuex";
import OpenScriptDialogButton from "./OpenScriptDialogButton.vue";
import CreateScriptDialogButton from "./CreateScriptDialogButton.vue";

export default Vue.extend({
  name: "AppBar",
  components: {
    OpenScriptDialogButton,
    CreateScriptDialogButton
  },
  data() {
    return {
      expanded: false
    };
  },
  methods: {
    toggleExpansion: function () {
      this.expanded = !this.expanded;
    }
  },
  computed: {
    ...mapState(["frame", "script"])
  }
});
</script>

<style scoped lang="scss">
.header {
  position: relative;
  padding: 4px 16px 0;
}

.menubar-container {
  display: grid;
  grid-template-areas:
    "script-title user-button"
    "menubar user-button";
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
