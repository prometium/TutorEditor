<template>
  <v-dialog v-model="dialog" width="600">
    <v-card>
      <v-card-title class="headline lighten-2">
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
        <v-btn @click="dialog = false" text> Отменить </v-btn>
        <v-btn @click="handleCopy" text color="primary"> Копировать </v-btn>
      </v-card-actions>
    </v-card>
  </v-dialog>
</template>

<script lang="ts">
import Vue from "vue";
import { mapActions, mapState } from "vuex";
import { copyScript } from "@/common/requests";
import { ActionTypes } from "@/store/action-types";

export default Vue.extend({
  name: "CopyScriptDialog",
  props: ["value"],
  data() {
    return {
      name: ""
    };
  },
  computed: {
    ...mapState(["script"]),
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
    handleCopy() {
      copyScript({ ...this.script, name: this.name })
        .then(data => {
          this.loadScript(data.uid);
          this.$router.push({ path: "/", query: { scriptUid: data.uid } });
          this.dialog = false;
        })
        .catch(console.error); // TODO: error handling
    }
  },
  watch: {
    dialog(value) {
      if (value) {
        this.name = this.script.name + " - копия";
      }
    }
  }
});
</script>
