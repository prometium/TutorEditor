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
        <v-text-field
          type="number"
          min="10"
          step="10"
          max="10000"
          v-model.number="currentDuration"
          label="Длительность (мс)"
        />
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
import { mapGetters, mapActions } from "vuex";
import { ActionTypes } from "@/store/action-types";

export default Vue.extend({
  name: "EditPauseActionDialog",
  props: {
    frameUid: { type: String, required: true },
    action: { type: Object, required: true }
  },
  data() {
    return {
      dialog: false,
      currentDuration: null
    };
  },
  methods: {
    ...mapActions({
      updateScript: ActionTypes.UPDATE_SCRIPT
    }),
    handleDiscard() {
      this.dialog = false;
    },
    async handleSubmit() {
      await this.updateScript({
        actionIdsToDel: [this.action.uid],
        frames: [
          {
            uid: this.frameUid,
            actions: [
              {
                uid: this.action.uid,
                actionType: this.action.actionType,
                nextFrame: this.action.nextFrame,
                duration: this.currentDuration
              }
            ]
          }
        ]
      });
      this.dialog = false;
    }
  },
  computed: {
    ...mapGetters(["currentAction"])
  },
  watch: {
    action: {
      immediate: true,
      handler(value) {
        this.currentDuration = value.duration;
      }
    }
  }
});
</script>
