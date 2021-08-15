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
            :value="currentFrame.taskText"
            label="Текст задания"
            dense
            hide-details
            @change="handleTextChange($event, 'taskText')"
          />
        </v-col>
        <v-col class="d-flex" cols="12" md="6">
          <v-text-field
            :value="currentFrame.hintText"
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
            item-text="text"
            item-value="value"
            v-model="currentActionType"
            :items="actionItems"
            label="Действие"
            dense
            hide-details
          />
          <component
            v-if="actionDialogComponent"
            :is="actionDialogComponent"
            :frameUid="currentFrame.uid"
            :action="currentAction"
          >
            <template v-slot:activator="{ on, attrs }">
              <v-btn v-bind="attrs" v-on="on" elevation="1" icon>
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
import Vue from "vue";
import { mapGetters, mapActions } from "vuex";
import { ActionTypes } from "@/store/action-types";
import { initialActionItems } from "./constants";
import EditKeyboardActionDialog from "./EditKeyboardActionDialog.vue";
import EditPauseActionDialog from "./EditPauseActionDialog.vue";
import { ActionGroup } from "@/common/constants";

export default Vue.extend({
  name: "Toolbar",
  components: {
    EditKeyboardActionDialog,
    EditPauseActionDialog
  },
  data() {
    return {
      expanded: false,
      actionItems: initialActionItems
    };
  },
  computed: {
    ...mapGetters(["currentFrame", "currentAction", "currentActionGroup"]),
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
    currentActionType: {
      get(): number {
        return this.currentAction?.actionType;
      },
      set(newValue: number) {
        this.updateScript({
          frames: [
            {
              uid: this.currentFrame.uid,
              actions: [
                {
                  uid: this.currentAction.uid,
                  actionType: newValue
                }
              ]
            }
          ]
        });
      }
    }
  },
  methods: {
    ...mapActions({
      updateScript: ActionTypes.UPDATE_SCRIPT
    }),
    handleToggleExpanded() {
      this.expanded = !this.expanded;
    },
    handleTextChange(newValue: string, field: string) {
      this.updateScript({
        frames: [
          {
            uid: this.currentFrame.uid,
            [field]: newValue
          }
        ]
      });
    }
  }
});
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
</style>
