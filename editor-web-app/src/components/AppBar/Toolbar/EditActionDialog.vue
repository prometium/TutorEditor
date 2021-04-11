<template>
  <v-dialog v-model="dialog" width="600">
    <template v-slot:activator="activator">
      <slot name="activator" v-bind="activator" />
    </template>
    <v-card>
      <v-card-title class="headline lighten-2">
        Редактирование действия
      </v-card-title>
      <v-card-text style="max-height: 300px">
        <p class="body-1" v-if="editMode">
          <v-text-field
            @keyup="handleKeyUp"
            @keydown.prevent=""
            solo
            label="Новая клавиша"
          />
          <v-btn @click="editMode = ''" text small> Отмена </v-btn>
        </p>
        <template v-else>
          <p class="body-1">
            Клавиша: <b>{{ action.key || "[нет]" }}</b>
            <v-btn @click="handleClickChangeKey" icon>
              <v-icon> mdi-pencil </v-icon>
            </v-btn>
            <v-btn icon>
              <v-icon> mdi-delete </v-icon>
            </v-btn>
          </p>
          <p class="body-1">
            Модификатор: <b>{{ action.modKey || "[нет]" }}</b>
            <v-btn @click="handleClickChangeModKey" icon>
              <v-icon> mdi-pencil </v-icon>
            </v-btn>
            <v-btn icon>
              <v-icon> mdi-delete </v-icon>
            </v-btn>
          </p>
        </template>
      </v-card-text>
      <v-divider />
      <v-card-actions>
        <v-spacer></v-spacer>
        <v-btn @click="dialog = false" text> Отменить </v-btn>
        <v-btn @click="dialog = false" text color="primary"> Сохранить </v-btn>
      </v-card-actions>
    </v-card>
  </v-dialog>
</template>

<script lang="ts">
import Vue from "vue";
import { mapActions, mapState } from "vuex";
import { ActionTypes } from "@/store/action-types";

enum EditMode {
  Key = "key",
  ModKey = "modKey",
  None = ""
}

export default Vue.extend({
  name: "ConfigureActionDialog",
  props: {
    frameUid: { type: String, required: true },
    action: { type: Object, required: true }
  },
  data() {
    return {
      dialog: false,
      editMode: EditMode.None
    };
  },
  methods: {
    ...mapActions({
      updateFrames: ActionTypes.UPDATE_FRAMES
    }),
    handleClickChangeKey() {
      this.editMode = EditMode.Key;
    },
    handleClickChangeModKey() {
      this.editMode = EditMode.ModKey;
    },
    handleKeyUp(event: KeyboardEvent) {
      this.updateFrames([
        {
          uid: this.frameUid,
          actions: [
            {
              uid: this.action.uid,
              [this.editMode]: event.code
            }
          ]
        }
      ]);
      this.editMode = EditMode.None;
    }
  },
  computed: {
    ...mapState(["scriptsInfo"])
  }
});
</script>
