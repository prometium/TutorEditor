<template>
  <div class="toolbar">
    <v-btn
      @click="handleToggleExpanded"
      elevation="1"
      icon
      :disabled="!showAction"
    >
      <v-icon>{{ expanded ? "mdi-chevron-up" : "mdi-chevron-down" }}</v-icon>
    </v-btn>
    <v-container v-if="currentFrame" fluid class="toolbar__container">
      <v-row align="center" de>
        <v-col class="d-flex" cols="12" md="6">
          <v-text-field
            :model-value="currentFrame.taskText"
            label="Текст задания"
            dense
            hide-details
            @change="handleTextChange($event, 'taskText')"
          />
        </v-col>
        <v-col class="d-flex" cols="12" md="6">
          <v-text-field
            :model-value="currentFrame.hintText"
            label="Текст подсказки"
            dense
            hide-details
            @change="handleTextChange($event, 'hintText')"
          />
        </v-col>
      </v-row>
      <v-row v-if="showAction" v-show="expanded">
        <v-col class="d-flex">
          <v-select
            item-title="text"
            item-value="value"
            v-model="currentActionType"
            :items="actionItems"
            label="Действие"
            dense
            hide-details
          />
          <span v-if="isTicksCountShown" class="action-immutable-value">
            {{ currentAction?.ticksCount || 0 }} щелчков
          </span>
          <component
            v-if="actionDialogComponent"
            :is="actionDialogComponent"
            :frameUid="currentFrame.uid"
            :action="currentAction"
          >
            <template v-slot:activator="{ on, attrs }">
              <v-btn
                v-bind="attrs"
                v-on="on"
                elevation="1"
                icon
                class="action-settings-btn"
              >
                <v-icon> mdi-cog </v-icon>
              </v-btn>
            </template>
          </component>
        </v-col>
      </v-row>
    </v-container>
  </div>
</template>

<script lang="ts">
import { mapActions, mapState } from "pinia";
import { useStore } from "@/store";
import { initialActionItems } from "./constants";
import EditKeyboardActionDialog from "./EditKeyboardActionDialog.vue";
import EditPauseActionDialog from "./EditPauseActionDialog.vue";
import { ActionGroup } from "@/common/constants";

export default {
  name: "AppToolbar",
  components: {
    EditKeyboardActionDialog,
    EditPauseActionDialog,
  },
  data() {
    return {
      expanded: false,
      actionItems: initialActionItems,
    };
  },
  computed: {
    ...mapState(useStore, [
      "script",
      "scriptsInfo",
      "currentFrame",
      "currentAction",
      "currentActionGroup",
    ]),
    showAction(): boolean {
      return !!this.currentFrame?.actions?.length;
    },
    actionDialogComponent() {
      switch (this.currentActionGroup) {
        case ActionGroup.Keyboard:
          return "EditKeyboardActionDialog";
        case ActionGroup.Pause:
          return "EditPauseActionDialog";
        default:
          return null;
      }
    },
    isTicksCountShown(): boolean {
      return this.currentActionGroup === ActionGroup.Tick;
    },
    currentActionType: {
      get(): number {
        return this.currentAction?.actionType || 0; // TODO
      },
      set(newValue: number) {
        this.updateScript({
          frames: [
            {
              uid: this.currentFrame?.uid || "",
              actions: [
                {
                  uid: this.currentAction?.uid || "",
                  actionType: newValue,
                },
              ],
            },
          ],
        });
      },
    },
  },
  methods: {
    ...mapActions(useStore, ["updateScript"]),
    handleToggleExpanded() {
      this.expanded = !this.expanded;
    },
    handleTextChange(newValue: string, field: string) {
      this.updateScript({
        frames: [
          {
            uid: this.currentFrame?.uid || "",
            [field]: newValue,
          },
        ],
      });
    },
  },
};
</script>

<style scoped lang="scss">
.toolbar {
  display: flex;
  padding: 4px 12px 0;
}

.toolbar__container {
  padding-left: 16px;
  padding-top: 0;
  padding-bottom: 0;
}

.action-immutable-value {
  margin-left: 8px;
  display: flex;
  align-items: center;
  cursor: default;
  color: rgba(0, 0, 0, 0.6);
}

.action-settings-btn {
  margin-left: 8px;
}
</style>
