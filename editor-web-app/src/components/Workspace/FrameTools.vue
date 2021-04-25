<template>
  <div class="frame-tools">
    <v-btn elevation="1" icon>
      <v-icon>mdi-arrow-up</v-icon>
    </v-btn>
    <v-btn elevation="1" icon>
      <v-icon>mdi-arrow-down</v-icon>
    </v-btn>
    <v-btn elevation="1" icon>
      <v-icon>mdi-plus</v-icon>
    </v-btn>
    <v-btn @click="handleDelete" elevation="1" icon>
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
      "prevAction"
    ])
  },
  methods: {
    ...mapActions({
      updateFrames: ActionTypes.UPDATE_SCRIPT
    }),
    ...mapMutations({
      selectFrame: MutationTypes.SELECT_FRAME
    }),
    async handleDelete() {
      if (this.path.length < 3) return;

      const hasPrev = this.prevFrame && this.prevAction;
      const prevFrameUid = this.prevFrame?.uid;
      const nextFrameUid = this.currentAction.nextFrame?.uid;

      await this.updateFrames({
        script: hasPrev
          ? {}
          : {
              firstFrame: { uid: this.currentAction.nextFrame.uid }
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
        frameIdsToDel: [this.currentFrame?.uid]
      });

      this.selectFrame(
        nextFrameUid ? nextFrameUid : prevFrameUid ? prevFrameUid : null
      );
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
