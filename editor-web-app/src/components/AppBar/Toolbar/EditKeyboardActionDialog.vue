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
            @keydown.prevent
            solo
            label="Новая клавиша (с модификатором или без)"
          />
          <v-btn @click="editMode = ''" text small> Отмена </v-btn>
        </p>
        <template v-else>
          <p class="body-1">
            Клавиша: <b>{{ key || "[нет]" }}</b>
            <v-btn @click="handleChange" icon>
              <v-icon> mdi-pencil </v-icon>
            </v-btn>
            <v-btn @click="handleDelete" icon>
              <v-icon> mdi-delete </v-icon>
            </v-btn>
          </p>
          <p class="body-1">
            Модификатор: <b>{{ modKey || "[нет]" }}</b>
          </p>
        </template>
      </v-card-text>
      <v-divider />
      <v-card-actions>
        <v-spacer />
        <v-btn @click="handleDiscard" text> Отменить </v-btn>
        <v-btn @click="handleSubmit" text color="primary"> Сохранить </v-btn>
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
  None = ""
}

// TODO: other mod keys
const modKeyArgs = ["Alt", "Control", "Shift"];

export default Vue.extend({
  name: "EditKeyboardActionDialog",
  props: {
    frameUid: { type: String, required: true },
    action: { type: Object, required: true }
  },
  data() {
    return {
      dialog: false,
      editMode: EditMode.None,
      key: "",
      modKey: ""
    };
  },
  methods: {
    ...mapActions({
      updateScript: ActionTypes.UPDATE_SCRIPT
    }),
    handleChange() {
      this.editMode = EditMode.Key;
    },
    handleKeyUp(event: KeyboardEvent) {
      this.key = event.key;
      this.modKey = "";
      for (let modKeyArg of modKeyArgs) {
        if (event.getModifierState(modKeyArg)) {
          this.modKey = modKeyArg;
          break;
        }
      }
      this.editMode = EditMode.None;
    },
    handleDelete() {
      this.key = "";
      this.modKey = "";
    },
    handleDiscard() {
      this.key = this.action.key;
      this.modKey = this.action.modKey;
      this.dialog = false;
    },
    async handleSubmit() {
      await this.updateScript({
        actionIdsToDel: [this.action.uid],
        frames: [
          {
            uid: this.frameUid,
            actions: [
              {
                uid: this.action.uid,
                actionType: this.action.actionType,
                nextFrame: this.action.nextFrame,
                key: this.key,
                modKey: this.modKey
              }
            ]
          }
        ]
      });
      this.dialog = false;
    }
  },
  computed: {
    ...mapState(["scriptsInfo"])
  },
  watch: {
    action: {
      immediate: true,
      handler(value) {
        this.key = value.key;
        this.modKey = value.modKey;
      }
    }
  }
});
</script>
