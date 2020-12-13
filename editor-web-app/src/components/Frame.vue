<template>
  <div class="frame">
    <img
      :src="frame.pictureLink + '#' + new Date().getTime()"
      :alt="frame.uid"
      class="frame__img"
    />
  </div>
</template>

<script lang="ts">
import Vue from "vue";
import { mapState } from "vuex";

export default Vue.extend({
  name: "Frame",
  computed: {
    ...mapState(["frame"])
  }
});
</script>

<style scoped lang="scss">
.frame {
  display: flex;
  align-items: center;
  justify-content: center;
  overflow-y: auto;
}

.frame__img {
  max-height: 100%;
  max-width: 100%;
}
</style>
