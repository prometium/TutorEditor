<template>
  <div class="frame-tools">
    <v-btn
      elevation="1"
      icon
      @click="handleUp"
      :disabled="!prevFrame || !prevAction"
    >
      <v-icon>mdi-arrow-up</v-icon>
    </v-btn>
    <v-btn
      elevation="1"
      icon
      @click="handleDown"
      :disabled="!nextFrame || !nextAction"
    >
      <v-icon>mdi-arrow-down</v-icon>
    </v-btn>
    <v-btn elevation="1" icon @click="handleAdd">
      <v-icon>mdi-plus</v-icon>
    </v-btn>
    <v-btn
      @click="handleDelete"
      elevation="1"
      icon
      :disabled="isDeletingDisabled"
    >
      <v-icon>mdi-delete</v-icon>
    </v-btn>
  </div>
</template>

<script lang="ts">
import { mapActions, mapState } from "pinia";
import { useStore } from "@/store";

export default {
  name: "FrameTools",
  computed: {
    ...mapState(useStore, [
      "script",
      "scriptsInfo",
      "path",
      "currentFrame",
      "currentAction",
      "prevFrame",
      "prevAction",
      "nextFrame",
      "nextAction",
    ]),
    isDeletingDisabled(): boolean {
      return (
        (this.currentFrame?.actions?.length || 0) > 1 || this.path.length <= 2
      );
    },
  },
  methods: {
    ...mapActions(useStore, ["updateScript", "selectFrame"]),
    async handleUp() {
      if (
        this.prevFrame == null ||
        this.prevAction == null ||
        this.currentFrame == null ||
        this.currentAction == null
      ) {
        return;
      }

      await this.updateScript({
        frames: [
          {
            ...this.currentFrame,
            uid: this.prevFrame.uid,
            actions: [
              {
                ...this.currentAction,
                uid: this.prevAction.uid,
                nextFrame: this.prevAction.nextFrame,
              },
            ],
          },
          {
            ...this.prevFrame,
            uid: this.currentFrame.uid,
            actions: [
              {
                ...this.prevAction,
                uid: this.currentAction.uid,
                nextFrame: this.currentAction.nextFrame,
              },
            ],
          },
        ],
      });

      this.selectFrame(this.prevFrame.uid);
    },
    async handleDown() {
      if (
        this.nextFrame == null ||
        this.nextAction == null ||
        this.currentFrame == null ||
        this.currentAction == null
      ) {
        return;
      }

      await this.updateScript({
        frames: [
          {
            ...this.currentFrame,
            uid: this.nextFrame.uid,
            actions: [
              {
                ...this.currentAction,
                uid: this.nextAction.uid,
                nextFrame: this.nextAction.nextFrame,
              },
            ],
          },
          {
            ...this.nextFrame,
            uid: this.currentFrame.uid,
            actions: [
              {
                ...this.nextAction,
                uid: this.currentAction.uid,
                nextFrame: this.currentAction.nextFrame,
              },
            ],
          },
        ],
      });

      this.selectFrame(this.nextFrame.uid);
    },
    async handleAdd() {
      if (this.currentFrame == null || this.currentAction == null) {
        return;
      }

      await this.updateScript({
        frames: [
          {
            uid: this.currentFrame.uid,
            actions: [
              {
                uid: this.currentAction.uid,
                nextFrame: { uid: "_:new1" },
              },
            ],
          },
          {
            uid: "_:new1",
            pictureLink: "",
            taskText: "",
            hintText: "",
            actions: [
              {
                uid: "_:new2",
                nextFrame: this.currentAction.nextFrame && {
                  uid: this.currentAction.nextFrame.uid,
                },
              },
            ],
          },
        ],
      });

      this.selectFrame(this.nextFrame?.uid);
    },
    async handleDelete() {
      if (this.path.length < 3) return;

      const hasPrev = !!this.prevFrame && !!this.prevAction;
      const prevFrameUid = this.prevFrame?.uid;

      if (this.currentFrame == null || this.prevAction == null) {
        return;
      }

      const nextFrameUid = this.currentAction?.nextFrame?.uid;

      const scriptForUpdate = hasPrev
        ? {}
        : {
            firstFrame: {
              uid: nextFrameUid || "",
            },
          };

      const framesForUpdate = hasPrev
        ? [
            {
              uid: prevFrameUid || "",
              actions: [
                {
                  ...this.prevAction,
                  nextFrame: nextFrameUid
                    ? {
                        uid: nextFrameUid,
                      }
                    : undefined,
                },
              ],
            },
          ]
        : [];

      await this.updateScript({
        script: scriptForUpdate,
        frames: framesForUpdate,
        frameIdsToDel: [this.currentFrame.uid],
        actionIdsToDel: this.prevAction ? [this.prevAction.uid] : undefined,
      });

      this.selectFrame(prevFrameUid || nextFrameUid);
    },
  },
};
</script>

<style scoped lang="scss">
.frame-tools {
  display: flex;
  justify-content: space-around;
  padding: 4px 0;
}
</style>
