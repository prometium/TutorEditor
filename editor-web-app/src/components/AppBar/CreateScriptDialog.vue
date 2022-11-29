<template>
  <v-dialog v-model="dialog" width="600">
    <v-card>
      <v-card-title class="text-h5">
        Создание обучающей программы
      </v-card-title>
      <v-card-text style="max-height: 300px">
        <v-text-field v-model="name" label="Название обучающей программы" />
        <v-file-input
          v-model="files"
          :label="`Выбрать архив от ${
            radioGroup === '1' ? 'перехватчика' : 'редактора'
          }`"
          accept=".zip"
        />
        <v-radio-group v-model="radioGroup">
          <v-radio label="От перехватчика" value="1" />
          <v-radio label="От редактора" value="2" />
        </v-radio-group>
      </v-card-text>
      <v-divider />
      <v-card-actions>
        <v-spacer />
        <v-btn @click="dialog = false" variant="text"> Отменить </v-btn>
        <v-btn
          @click="handleCreate"
          :loading="isLoading"
          variant="text"
          color="primary"
        >
          Создать
        </v-btn>
      </v-card-actions>
    </v-card>
  </v-dialog>
</template>

<script lang="ts">
import { mapActions, mapState } from "pinia";
import { useStore } from "@/store";
import { createScript, createScriptV2 } from "@/common/requests";

export default {
  name: "CreateScriptDialog",
  props: ["modelValue"],
  data() {
    return {
      name: "",
      files: [] as File[],
      radioGroup: "1",
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
    handleCreate() {
      if (!this.files.length) return;

      const formData = new FormData();
      formData.append("script", this.files[0]);
      formData.append("name", this.name);

      const action = this.radioGroup === "1" ? createScript : createScriptV2;

      this.isLoading = true;
      action(formData)
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
        this.name = "";
        this.files = [];
        this.radioGroup = "1";
      }
    },
  },
};
</script>
