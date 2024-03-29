<template>
  <v-dialog v-model="dialog" width="600">
    <v-card>
      <v-card-title class="headline lighten-2">
        Создание обучающей программы
      </v-card-title>
      <v-card-text style="max-height: 300px">
        <v-text-field v-model="name" label="Название обучающей программы" />
        <v-file-input
          v-model="file"
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
        <v-btn @click="dialog = false" text> Отменить </v-btn>
        <v-btn @click="handleCreate" :loading="isLoading" text color="primary">
          Создать
        </v-btn>
      </v-card-actions>
    </v-card>
  </v-dialog>
</template>

<script lang="ts">
import Vue from "vue";
import { mapActions, mapState } from "vuex";
import { createScript, createScriptV2 } from "@/common/requests";
import { ActionTypes } from "@/store/action-types";

export default Vue.extend({
  name: "CreateScriptDialog",
  props: ["value"],
  data() {
    return {
      name: "",
      file: null as File | null,
      radioGroup: "1",
      isLoading: false
    };
  },
  computed: {
    ...mapState(["scriptsInfo"]),
    dialog: {
      get(): boolean {
        return this.value;
      },
      set(value: boolean) {
        this.$emit("input", value);
      }
    }
  },
  methods: {
    ...mapActions({
      loadScript: ActionTypes.LOAD_SCRIPT
    }),
    handleCreate() {
      if (!this.file) return;

      const formData = new FormData();
      formData.append("script", this.file);
      formData.append("name", this.name);

      const action = this.radioGroup === "1" ? createScript : createScriptV2;

      this.isLoading = true;
      action(formData)
        .then(data => {
          this.loadScript(data.uid);
          this.$router.push({ path: "/", query: { scriptUid: data.uid } });
          this.dialog = false;
        })
        .catch(console.error)
        .finally(() => {
          this.isLoading = false;
        });
    }
  },
  watch: {
    dialog(value) {
      if (value) {
        this.name = "";
        this.file = null;
        this.radioGroup = "1";
      }
    }
  }
});
</script>
