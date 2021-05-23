<template>
  <v-dialog v-model="dialog" width="600">
    <template v-slot:activator="activator">
      <slot name="activator" v-bind="activator" />
    </template>
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
        <v-btn @click="handleCreate" text color="primary"> Создать </v-btn>
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
  data() {
    return {
      dialog: false,
      loading: false,
      name: "",
      file: null as File | null,
      radioGroup: "1"
    };
  },
  computed: {
    ...mapState(["scriptsInfo"])
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

      action(formData)
        .then(data => {
          this.loadScript(data.uid);
          this.$router.push({ path: "/", query: { scriptUid: data.uid } });
          this.dialog = false;
        })
        .catch(console.error); // TODO: error handling
    }
  }
});
</script>
