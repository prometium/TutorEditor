<template>
  <v-sheet tag="header" class="header" elevation="4">
    <div class="menubar-container">
      <div class="script-title text-h6">{{ script.name || "..." }}</div>
      <div class="menubar">
        <v-menu>
          <template v-slot:activator="{ on: menu, attrs }">
            <v-btn v-bind="attrs" v-on="{ ...menu }" small text elevation="0">
              Файл
            </v-btn>
          </template>
          <v-list>
            <v-list-item v-for="(item, index) in items" :key="index">
              <v-list-item-title>{{ item.title }}</v-list-item-title>
            </v-list-item>
          </v-list>
        </v-menu>
        <v-btn v-bind="attrs" v-on="{ ...menu }" small text elevation="0">
          Редактирование
        </v-btn>
      </div>
      <v-btn icon large class="user-button">
        <v-icon>mdi-account-circle</v-icon>
      </v-btn>
    </div>
    <div class="toolbar">
      <v-btn @click="toggle" elevation="1" icon
        ><v-icon>{{ open ? "mdi-arrow-up" : "mdi-arrow-down" }}</v-icon></v-btn
      >
      <v-container fluid class="toolbar__container">
        <v-row align="center">
          <v-col class="d-flex" md="6">
            <v-text-field
              v-model="frame.taskText"
              label="Текст задания"
              dense
            ></v-text-field>
          </v-col>
          <v-col class="d-flex" md="6">
            <v-text-field
              v-model="frame.hintText"
              label="Текст подсказки"
              dense
            ></v-text-field>
          </v-col>
          <!-- <v-col class="d-flex" cols="1">
          <v-btn @click="toggle" elevation="1" icon
            ><v-icon>{{
              open ? "mdi-arrow-up" : "mdi-arrow-down"
            }}</v-icon></v-btn
          >
        </v-col> -->
        </v-row>
        <v-row v-show="open">
          <v-col class="d-flex">
            <v-select :items="items" label="Действие" dense></v-select>
          </v-col>
        </v-row>
      </v-container>
    </div>
  </v-sheet>
</template>

<script lang="ts">
import Vue from "vue";
import { mapState } from "vuex";
import store from "@/store";

export default Vue.extend({
  name: "AppBar",
  store,
  data() {
    return {
      open: false,
      scriptUid: "",
      items: [
        { title: "Click Me1" },
        { title: "Click Me2" },
        { title: "Click Me3" },
        { title: "Click Me4" }
      ]
    };
  },
  methods: {
    toggle: function() {
      this.open = !this.open;
    }
  },
  computed: {
    ...mapState(["frame", "script"])
  }
});
</script>

<style scoped lang="scss">
.header {
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
  align-items: flex-start;
  padding: 4px 12px 0;
}

.toolbar__container {
  padding-left: 16px;
  padding-top: 0;
  padding-bottom: 0;
}
</style>
