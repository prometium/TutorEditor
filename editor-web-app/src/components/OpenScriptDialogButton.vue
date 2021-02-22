<template>
  <v-dialog v-model="dialog" width="500">
    <template v-slot:activator="{ on, attrs }">
      <v-list-item v-bind="attrs" v-on="on" @click="loadScriptsInfo">
        <v-list-item-title>Открыть</v-list-item-title>
      </v-list-item>
    </template>
    <v-card>
      <v-card-title class="headline lighten-2">
        Обучающие программы
      </v-card-title>
      <v-card-text style="max-height: 300px">
        <v-radio-group v-model="radioGroup" column>
          <v-radio
            v-for="scriptInfo in scriptsInfo"
            :key="scriptInfo.uid"
            :label="scriptInfo.name"
            :value="scriptInfo.uid"
          ></v-radio>
        </v-radio-group>
      </v-card-text>
      <v-divider></v-divider>
      <v-card-actions>
        <v-spacer></v-spacer>
        <v-btn @click="dialog = false" text> Отменить </v-btn>
        <v-btn @click="handleOpen" text color="primary"> Открыть </v-btn>
      </v-card-actions>
    </v-card>
  </v-dialog>
</template>

<script lang="ts">
import Vue from "vue";
import { mapState, mapActions } from "vuex";
import { ActionTypes } from "@/store/action-types";

export default Vue.extend({
  name: "OpenScriptDialogButton",
  data() {
    return {
      dialog: false,
      radioGroup: ""
    };
  },
  methods: {
    ...mapActions({
      loadScriptsInfo: ActionTypes.LOAD_SCRIPTS_INFO,
      loadScript: ActionTypes.LOAD_SCRIPT
    }),
    handleOpen() {
      this.dialog = false;
      this.loadScript(this.radioGroup);
    }
  },
  computed: {
    ...mapState(["scriptsInfo"])
  }
});
</script>
