<template>
  <v-dialog v-model="dialog" width="600">
    <template v-slot:activator="activator">
      <slot name="activator" v-bind="activator" />
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
          />
        </v-radio-group>
      </v-card-text>
      <v-divider />
      <v-card-actions>
        <v-spacer />
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
  name: "OpenScriptDialog",
  data() {
    return {
      dialog: false,
      radioGroup: ""
    };
  },
  methods: {
    ...mapActions({
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
