<template>
  <v-dialog v-model="dialog" width="600">
    <template v-slot:activator="{ on, attrs }">
      <v-list-item v-bind="attrs" v-on="on">
        <v-list-item-title>Создать</v-list-item-title>
      </v-list-item>
    </template>
    <v-card>
      <v-card-title class="headline lighten-2">
        Создание обучающей программы
      </v-card-title>
      <v-card-text style="max-height: 300px">
        <v-text-field v-model="name" label="Название обучающей программы">
        </v-text-field>
        <v-file-input
          label="Выбрать заготовку"
          truncate-length="15"
          v-model="file"
        />
      </v-card-text>
      <v-divider />
      <v-card-actions>
        <v-spacer></v-spacer>
        <v-btn @click="dialog = false" text> Отменить </v-btn>
        <v-btn @click="handleCreate" text color="primary"> Создать </v-btn>
      </v-card-actions>
    </v-card>
  </v-dialog>
</template>

<script lang="ts">
import Vue from "vue";
import { mapActions, mapState } from "vuex";
import { createScript } from "@/common/requests";
import { ActionTypes } from "@/store/action-types";

export default Vue.extend({
  name: "CreateScriptDialogButton",
  data() {
    return {
      dialog: false,
      loading: false,
      name: "",
      file: null as File | null
    };
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

      createScript(formData)
        .then(data => {
          this.loadScript(data.uid);
          this.dialog = false;
        })
        .catch(console.error); // TODO error
    }
  },
  computed: {
    ...mapState(["scriptsInfo"])
  }
});
</script>
