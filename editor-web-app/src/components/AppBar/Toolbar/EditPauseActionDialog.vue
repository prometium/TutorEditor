<template>
  <v-dialog v-model="dialog" width="600">
    <template v-slot:activator="activator">
      <slot name="activator" v-bind="activator" />
    </template>
    <v-card>
      <v-card-title class="headline lighten-2">
        Редактирование действия
      </v-card-title>
      <v-card-text style="max-height: 300px">
        <v-text-field v-model="currentDuration" label="Длительность (мс)" />
      </v-card-text>
      <v-divider />
      <v-card-actions>
        <v-spacer />
        <v-btn @click="handleDiscard" text> Отменить </v-btn>
        <v-btn @click="handleSubmit" text color="primary"> Сохранить </v-btn>
      </v-card-actions>
    </v-card>
  </v-dialog>
</template>

<script lang="ts">
import Vue from "vue";
import { mapState, mapActions } from "vuex";
import { ActionTypes } from "@/store/action-types";

export default Vue.extend({
  name: "EditPauseActionDialog",
  props: {
    frameUid: { type: String, required: true },
    action: { type: Object, required: true }
  },
  data() {
    return {
      dialog: false
    };
  },
  methods: {
    ...mapActions({
      updateScript: ActionTypes.UPDATE_SCRIPT
    }),
    handleDiscard() {
      this.dialog = false;
    },
    handleSubmit() {
      this.dialog = false;
    }
  },
  computed: {
    ...mapState(["currentAction"]),
    currentDuration: {
      get(): number {
        return this.currentAction?.actionType;
      },
      set(newValue: number) {
        //
      }
    }
  }
});
</script>
