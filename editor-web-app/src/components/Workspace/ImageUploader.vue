<template>
  <v-file-input
    v-model="files"
    label="Выбрать новое изображение"
    accept=".png"
    class="file-input"
  />
</template>

<script lang="ts">
import { mapState, mapActions } from "pinia";
import { useStore } from "@/store";
import { addImage } from "@/common/requests";

export default {
  name: "ImageUploader",
  data() {
    return {
      files: [] as File[],
    };
  },
  computed: {
    ...mapState(useStore, ["currentFrame", "script", "scriptsInfo"]),
  },
  methods: {
    ...mapActions(useStore, ["updateScript"]),
  },
  watch: {
    files(value: File[]) {
      if (value.length) return;

      const formData = new FormData();
      formData.append("image", value[0]);

      addImage(formData)
        .then(({ link }) => {
          if (this.currentFrame) {
            this.updateScript({
              frames: [
                {
                  uid: this.currentFrame.uid,
                  pictureLink: link,
                },
              ],
            });
          }
        })
        .finally(() => {
          this.files = [];
        });
    },
  },
};
</script>

<style scoped lang="scss">
.file-input {
  margin: 0 auto;
  max-width: 300px;
}
</style>
