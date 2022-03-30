<template>
  <div class="frame-previews">
    <div
      v-for="(pathItem, index) in path"
      :key="pathItem.frameUid"
      class="frame-previews__item"
      @click="selectFrame(pathItem.frameUid)"
    >
      <template v-if="script.frameByUid[pathItem.frameUid].pictureLink">
        <img
          :src="getPictureLink(pathItem)"
          :alt="`Превью кадра №${index}`"
          :class="[
            'frame-previews__img',
            currentFrame && pathItem.frameUid === currentFrame.uid && 'active'
          ]"
          loading="lazy"
        />
        <div
          v-if="(script.frameByUid[pathItem.frameUid].actions || []).length > 1"
          class="frame-previews__branches"
        >
          <span
            v-for="offsetBranchNum in script.frameByUid[pathItem.frameUid]
              .actions.length"
            :key="offsetBranchNum"
            :style="{
              width: `calc(100% / ${
                script.frameByUid[pathItem.frameUid].actions.length
              })`
            }"
            :class="[
              'frame-previews__branch',
              offsetBranchNum - 1 === pathItem.branchNum && 'active'
            ]"
            @click="
              configurePath({
                frameUid: pathItem.frameUid,
                branchNum: offsetBranchNum - 1
              })
            "
          />
        </div>
      </template>
      <div
        v-else
        :class="[
          'frame-previews__img',
          'empty',
          currentFrame && pathItem.frameUid === currentFrame.uid && 'active'
        ]"
      />
    </div>
  </div>
</template>

<script lang="ts">
import Vue from "vue";
import { mapState, mapMutations, mapGetters } from "vuex";
import { MutationTypes } from "@/store/mutation-types";
import { API_ROOT } from "@/common/requests";
import { PathItem } from "@/common/types";

export default Vue.extend({
  name: "FramePreviews",
  data() {
    return {
      API_ROOT
    };
  },
  computed: {
    ...mapState(["script"]),
    ...mapGetters(["path", "currentFrame"])
  },
  methods: {
    ...mapMutations({
      selectFrame: MutationTypes.SELECT_FRAME,
      configurePath: MutationTypes.CONFIGURE_PATH
    }),
    getPictureLink(pathItem: PathItem): string {
      return `${process.env.VUE_APP_S3_URL || ""}/${
        process.env.VUE_APP_S3_BUCKET_NAME || "editor"
      }/${this.script.frameByUid[pathItem.frameUid].pictureLink}`;
    }
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

    &.empty {
      height: 60px;
      background: linear-gradient(
          rgba(255, 255, 255, 0.6) 40%,
          rgba(255, 255, 255, 0) 40%
        ),
        linear-gradient(rgba(255, 255, 255, 0.6), rgba(255, 255, 255, 0));
      background-size: 1em 1em, 100%;
      background-color: var(--v-secondary-base);
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
