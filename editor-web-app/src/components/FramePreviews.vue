<template>
  <aside class="frame-previews">
    <img
      v-for="(pathItem, index) in path"
      :src="script.frameByUid[pathItem.frameUid].pictureLink"
      :key="pathItem.frameUid"
      v-on:click="setFrame(pathItem.frameUid)"
      :alt="`Кадр ${index}`"
      loading="lazy"
      :class="[
        'frame-previews__img',
        { active: pathItem.frameUid == frame.uid }
      ]"
    />
  </aside>
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
      setFrame: MutationTypes.SET_FRAME
    })
  }
});
</script>

<style scoped lang="scss">
.frame-previews {
  overflow-y: auto;
}

.frame-previews__img {
  width: 100%;
  outline-offset: -3px;

  &:hover {
    outline: 3px solid var(--palette-default);
  }

  &.active {
    outline: 3px solid var(--palette-primary);
  }
}
</style>
