<template>
  <div class="toolbar">
    <v-btn @click="toggleExpansion" elevation="1" icon :disabled="!showAction">
      <v-icon>{{ expanded ? "mdi-chevron-up" : "mdi-chevron-down" }}</v-icon>
    </v-btn>
    <v-container fluid class="toolbar__container">
      <v-row align="center">
        <v-col class="d-flex" cols="12" md="6">
          <v-text-field v-model="frame.taskText" label="Текст задания" dense />
        </v-col>
        <v-col class="d-flex" cols="12" md="6">
          <v-text-field
            v-model="frame.hintText"
            label="Текст подсказки"
            dense
          />
        </v-col>
      </v-row>
      <v-row v-if="showAction" v-show="expanded">
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
</template>

<script lang="ts">
import Vue from "vue";
import { mapState, mapGetters, mapActions } from "vuex";
import { ActionType } from "@/common/constants";
import { ActionTypes } from "@/store/action-types";

export default Vue.extend({
  name: "Toolbar",
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
    ...mapActions({
      updateFrames: ActionTypes.UPDATE_FRAMES
    }),
    toggleExpansion() {
      this.expanded = !this.expanded;
    }
  },
  computed: {
    ...mapState(["frame", "script"]),
    ...mapGetters(["path"]),
    showAction(): boolean {
      return !!this.frame.actions?.length;
    },
    selectedAction: {
      get(): ActionType | null {
        const currentBranchNum =
          this.script.branchNumByUid[this.frame.uid] || 0;
        return this.frame.actions[currentBranchNum].actionType;
      },
      set(newValue) {
        const currentFrame = this.frame;
        const currentBranchNum =
          this.script.branchNumByUid[this.frame.uid] || 0;
        const currentAction = this.frame.actions[currentBranchNum];
        this.updateFrames([
          {
            uid: currentFrame.uid,
            actions: [
              {
                uid: currentAction.uid,
                actionType: newValue
              }
            ]
          }
        ]);
        this.frame.actions[
          this.script.branchNumByUid[this.frame.uid] || 0
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
