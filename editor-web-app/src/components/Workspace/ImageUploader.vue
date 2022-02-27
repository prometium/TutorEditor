<template>
  <v-file-input
    v-model="file"
    label="Выбрать новое изображение"
    accept=".png"
    class="file-input"
  />
</template>

<script lang="ts">
import Vue from "vue";
import { mapGetters, mapActions } from "vuex";
import { ActionTypes } from "@/store/action-types";
import { addImage } from "@/common/requests";

export default Vue.extend({
  name: "Frame",
  data() {
    return {
      file: null as File | null
    };
  },
  computed: {
    ...mapGetters(["currentFrame"])
  },
  methods: {
    ...mapActions({
      updateScript: ActionTypes.UPDATE_SCRIPT
    })
  },
  watch: {
    file(value) {
      if (!value) return;

      const formData = new FormData();
      formData.append("image", value);

      addImage(formData)
        .then(({ link }) => {
          this.updateScript({
            frames: [
              {
                uid: this.currentFrame.uid,
                pictureLink: link
              }
            ]
          });
        })
        .finally(() => {
          this.file = null;
        });
    }
  }
});
</script>

<style scoped lang="scss">
.file-input {
  margin: 0 auto;
  max-width: 300px;
}
</style>
