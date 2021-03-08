<template>
  <div class="frame-previews">
    <div
      v-for="(pathItem, index) in path"
      :key="pathItem.frameUid"
      class="frame-previews__item"
    >
      <img
        :src="script.frameByUid[pathItem.frameUid].pictureLink"
        :alt="`Кадр
      ${index}`"
        :class="[
          'frame-previews__img',
          {
            active: pathItem.frameUid === frame.uid
          }
        ]"
        @click="setFrame(pathItem.frameUid)"
        loading="lazy"
      />
      <div
        v-if="
          script.frameByUid[pathItem.frameUid].actions &&
          script.frameByUid[pathItem.frameUid].actions.length > 1
        "
        class="frame-previews__branches"
      >
        <span
          v-for="offsetBranchNum in script.frameByUid[pathItem.frameUid].actions
            .length"
          :key="offsetBranchNum"
          :style="{
            width: `calc(100% / ${
              script.frameByUid[pathItem.frameUid].actions.length
            })`
          }"
          :class="[
            'frame-previews__branch',
            {
              active: offsetBranchNum - 1 === pathItem.branchNum
            }
          ]"
          @click="
            configurePath({
              frameUid: pathItem.frameUid,
              branchNum: offsetBranchNum - 1
            })
          "
        />
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import Vue from "vue";
import { mapState, mapMutations, mapGetters } from "vuex";
import { MutationTypes } from "@/store/mutation-types";

export default Vue.extend({
  name: "FramePreviews",
  computed: {
    ...mapState(["script", "frame"]),
    ...mapGetters(["path"])
  },
  methods: {
    ...mapMutations({
      setFrame: MutationTypes.SET_FRAME,
      configurePath: MutationTypes.CONFIGURE_PATH
    })
  }
});
</script>

<style scoped lang="scss">
.frame-previews {
  overflow-y: auto;

  &__item {
    display: flex;
    flex-direction: column;
    padding-bottom: 8px;
  }

  &__img {
    width: 100%;
    outline-offset: -3px;

    &:hover {
      outline: 3px solid var(--v-secondary-base);
    }

    &.active {
      outline: 3px solid var(--v-accent-base);
    }
  }

  &__branches {
    display: flex;
    height: 16px;
  }

  &__branch {
    height: 100%;
    cursor: pointer;

    &:hover {
      background-color: var(--v-secondary-base);
    }

    &.active {
      background-color: var(--v-accent-base);
    }
  }
}
</style>
