<template>
  <v-dialog v-model="dialog" width="600">
    <v-card>
      <v-card-title class="text-h5">
        Создание копии обучающей программы
      </v-card-title>
      <v-card-text style="max-height: 300px">
        <v-text-field
          v-model="name"
          label="Название копии обучающей программы"
        />
      </v-card-text>
      <v-divider />
      <v-card-actions>
        <v-spacer />
        <v-btn @click="dialog = false" variant="text"> Отменить </v-btn>
        <v-btn
          @click="handleCopy"
          :loading="isLoading"
          variant="text"
          color="primary"
        >
          Копировать
        </v-btn>
      </v-card-actions>
    </v-card>
  </v-dialog>
</template>

<script lang="ts">
import { mapActions, mapState } from "pinia";
import { useStore } from "@/store";
import { copyScript } from "@/common/requests";

export default {
  name: "CopyScriptDialog",
  props: ["modelValue"],
  data() {
    return {
      name: "",
      isLoading: false,
    };
  },
  computed: {
    ...mapState(useStore, ["script", "scriptsInfo"]),
    dialog: {
      get(): boolean {
        return this.modelValue;
      },
      set(value: boolean) {
        this.$emit("update:modelValue", value);
      },
    },
  },
  methods: {
    ...mapActions(useStore, ["loadScript"]),
    handleCopy() {
      this.isLoading = true;
      copyScript({
        ...this.script,
        frames: Object.values(this.script.frameByUid),
        name: this.name,
      })
        .then((data) => {
          this.loadScript(data.uid);
          this.$router.push({ path: "/", query: { scriptUid: data.uid } });
          this.dialog = false;
        })
        .catch(console.error)
        .finally(() => {
          this.isLoading = false;
        });
    },
  },
  watch: {
    dialog(value) {
      if (value) {
        this.name = this.script.name + " - копия";
      }
    },
  },
};
</script>
