<template>
  <v-sheet tag="header" class="header" elevation="4">
    <div class="menubar-container">
      <div class="script-title text-h6">{{ script.name || "..." }}</div>
      <div class="menubar">
        <v-menu>
          <template v-slot:activator="{ on, attrs }">
            <v-btn v-bind="attrs" v-on="on" small text elevation="0">
              Файл
            </v-btn>
          </template>
          <v-list>
            <CreateScriptDialogButton />
            <OpenScriptDialogButton />
          </v-list>
        </v-menu>
        <v-btn small text elevation="0"> Редактирование </v-btn>
      </div>
      <v-btn icon large class="user-button">
        <v-icon>mdi-account-circle</v-icon>
      </v-btn>
    </div>
    <div class="toolbar">
      <v-btn @click="toggleExpansion" elevation="1" icon>
        <v-icon>{{ expanded ? "mdi-chevron-up" : "mdi-chevron-down" }}</v-icon>
      </v-btn>
      <v-container fluid class="toolbar__container">
        <v-row align="center">
          <v-col class="d-flex" cols="12" md="6">
            <v-text-field
              v-model="frame.taskText"
              label="Текст задания"
              dense
            />
          </v-col>
          <v-col class="d-flex" cols="12" md="6">
            <v-text-field
              v-model="frame.hintText"
              label="Текст подсказки"
              dense
            />
          </v-col>
        </v-row>
        <v-row v-show="expanded">
          <v-col>
            <v-select
              class="d-flex"
              item-text="text"
              item-value="value"
              v-model="selectedAction"
              :items="actionItems"
              label="Действие"
              dense
            />
          </v-col>
        </v-row>
      </v-container>
    </div>
  </v-sheet>
</template>

<script lang="ts">
import Vue from "vue";
import { mapState, mapGetters } from "vuex";
import OpenScriptDialogButton from "./OpenScriptDialogButton.vue";
import CreateScriptDialogButton from "./CreateScriptDialogButton.vue";
import { ActionType } from "@/common/constants";

export default Vue.extend({
  name: "AppBar",
  components: {
    OpenScriptDialogButton,
    CreateScriptDialogButton
  },
  data() {
    return {
      expanded: false,
      actionItems: [
        {
          value: ActionType.LeftMouseClick,
          text: "Левый щелчек мышью"
        },
        {
          value: ActionType.RightMouseClick,
          text: "Правый щелчек мышью"
        },
        {
          value: ActionType.LeftMouseUp,
          text: "Левое отжатие мышью"
        },
        {
          value: ActionType.LeftMouseDoubleClick,
          text: "Двойной левый щелчок мышью"
        },
        {
          value: ActionType.RightMouseClick,
          text: "Правый щелчок мышью"
        },
        {
          value: ActionType.RightMouseDown,
          text: "Правое нажатие мышью"
        },
        {
          value: ActionType.RightMouseDown,
          text: "Правое отжатие мышью"
        },
        {
          value: ActionType.RightMouseDobleClick,
          text: "Двойной правый щелчок мышью"
        },
        {
          value: ActionType.KeyClick,
          text: "Щелчок кнопки на клавиатуре"
        },
        {
          value: ActionType.KeyDown,
          text: "Нажатие кнопки на клавиатуре"
        },
        {
          value: ActionType.KeyUp,
          text: "Отжатие кнопки на клавиатуре"
        },
        {
          value: ActionType.KeyWithMod,
          text: "Кнопка на клавиатуре с модификатором"
        },
        {
          value: ActionType.Drag,
          text: "Перетаскивание"
        },
        {
          value: ActionType.WheelUp,
          text: "Прокрутка колесиком вверх"
        },
        {
          value: ActionType.WheelDown,
          text: "Прокрутка колесиком вниз"
        },
        {
          value: ActionType.WheelClick,
          text: "Нажатие на колесико"
        },
        {
          value: ActionType.Pause,
          text: "Пауза"
        }
      ]
    };
  },
  methods: {
    toggleExpansion: function () {
      this.expanded = !this.expanded;
    }
  },
  computed: {
    ...mapState(["frame", "script"]),
    ...mapGetters(["path"]),
    selectedAction: {
      get() {
        return (
          this?.frame?.actions &&
          this.frame.actions[this.script?.branchNumByUid[this.frame?.uid] || 0]
            .actionType
        );
      },
      set(newValue) {
        this.frame.actions[
          this.script.branchNumByUid[this.frame?.uid] || 0
        ].actionType = newValue;
      }
    }
  }
});
</script>

<style scoped lang="scss">
.header {
  position: relative;
  padding: 4px 16px 0;
}

.menubar-container {
  display: grid;
  grid-template-areas:
    "script-title user-button"
    "menubar user-button";
}

.script-title {
  grid-area: script-title;
  padding-left: 12px;
}

.menubar {
  grid-area: menubar;
}

.user-button {
  grid-area: user-button;
  align-self: center;
  justify-self: end;
}

.toolbar {
  display: flex;
  padding: 4px 12px 0;
}

.toolbar__container {
  padding-left: 16px;
  padding-top: 0;
  padding-bottom: 0;
}
</style>
