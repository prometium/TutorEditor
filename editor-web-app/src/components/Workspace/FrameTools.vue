<template>
  <div class="frame-tools">
    <v-btn elevation="1" icon @click="handleUp">
      <v-icon>mdi-arrow-up</v-icon>
    </v-btn>
    <v-btn elevation="1" icon @click="handleDown">
      <v-icon>mdi-arrow-down</v-icon>
    </v-btn>
    <v-btn elevation="1" icon @click="handleAdd">
      <v-icon>mdi-plus</v-icon>
    </v-btn>
    <v-btn
      @click="handleDelete"
      elevation="1"
      icon
      :disabled="isDeleteDisabled"
    >
      <v-icon>mdi-delete</v-icon>
    </v-btn>
  </div>
</template>

<script lang="ts">
import Vue from "vue";
import { mapActions, mapGetters, mapMutations } from "vuex";
import { ActionTypes } from "@/store/action-types";
import { MutationTypes } from "@/store/mutation-types";

export default Vue.extend({
  name: "FrameTools",
  computed: {
    ...mapGetters([
      "path",
      "currentFrame",
      "currentAction",
      "prevFrame",
      "prevAction",
      "nextFrame",
      "nextAction"
    ]),
    isDeleteDisabled(): boolean {
      return this.currentFrame?.actions?.length > 1 || this.path.length <= 2;
    }
  },
  methods: {
    ...mapActions({
      updateScript: ActionTypes.UPDATE_SCRIPT
    }),
    ...mapMutations({
      selectFrame: MutationTypes.SELECT_FRAME
    }),
    async handleUp() {
      if (!this.prevFrame || !this.prevAction) return;

      await this.updateScript({
        frames: [
          {
            ...this.currentFrame,
            uid: this.prevFrame.uid,
            actions: [
              {
                ...this.currentAction,
                uid: this.prevAction.uid,
                nextFrame: this.prevAction.nextFrame
              }
            ]
          },
          {
            ...this.prevFrame,
            uid: this.currentFrame.uid,
            actions: [
              {
                ...this.prevAction,
                uid: this.currentAction.uid,
                nextFrame: this.currentAction.nextFrame
              }
            ]
          }
        ]
      });

      this.selectFrame(this.prevFrame.uid);
    },
    async handleDown() {
      if (!this.nextFrame || !this.nextAction) return;

      await this.updateScript({
        frames: [
          {
            ...this.currentFrame,
            uid: this.nextFrame.uid,
            actions: [
              {
                ...this.currentAction,
                uid: this.nextAction.uid,
                nextFrame: this.nextAction.nextFrame
              }
            ]
          },
          {
            ...this.nextFrame,
            uid: this.currentFrame.uid,
            actions: [
              {
                ...this.nextAction,
                uid: this.currentAction.uid,
                nextFrame: this.currentAction.nextFrame
              }
            ]
          }
        ]
      });

      this.selectFrame(this.nextFrame.uid);
    },
    async handleAdd() {
      await this.updateScript({
        frames: [
          {
            uid: this.currentFrame.uid,
            actions: [
              {
                uid: this.currentAction.uid,
                nextFrame: { uid: "_:new1" }
              }
            ]
          },
          {
            uid: "_:new1",
            actions: [
              {
                uid: "_:new2",
                nextFrame: this.currentAction.nextFrame && {
                  uid: this.currentAction.nextFrame.uid
                }
              }
            ]
          }
        ]
      });

      this.selectFrame(this.nextFrame.uid);
    },
    async handleDelete() {
      if (this.path.length < 3) return;

      const hasPrev = !!this.prevFrame && !!this.prevAction;
      const prevFrameUid = this.prevFrame?.uid;
      const nextFrameUid = this.currentAction.nextFrame?.uid;

      await this.updateScript({
        script: hasPrev
          ? {}
          : {
              firstFrame: {
                uid: this.currentAction.nextFrame.uid
              }
            },
        frames: hasPrev
          ? [
              {
                uid: prevFrameUid,
                actions: [
                  {
                    uid: this.prevAction.uid,
                    nextFrame: nextFrameUid
                      ? {
                          uid: nextFrameUid
                        }
                      : null
                  }
                ]
              }
            ]
          : [],
        frameIdsToDel: [this.currentFrame.uid]
      });

      this.selectFrame(prevFrameUid || nextFrameUid || null);
    }
  }
});
</script>

<style scoped lang="scss">
.frame-tools {
  display: flex;
  justify-content: space-around;
  padding: 4px 0;
}
</style>
