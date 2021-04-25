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
    <v-container fluid class="toolbar__container">
      <v-row align="center" de>
        <v-col class="d-flex" cols="12" md="6">
          <v-text-field
            v-model="frame.taskText"
            label="Текст задания"
            dense
            hide-details
          />
        </v-col>
        <v-col class="d-flex" cols="12" md="6">
          <v-text-field
            v-model="frame.hintText"
            label="Текст подсказки"
            dense
            hide-details
          />
        </v-col>
      </v-row>
      <v-row v-if="showAction" v-show="expanded">
        <v-col class="d-flex">
          <v-select
            item-text="text"
            item-value="value"
            v-model="selectedActionType"
            :items="actionItems"
            label="Действие"
            dense
            hide-details
          />
          <component
            v-if="actionDialogComponent"
            :is="actionDialogComponent"
            :frameUid="frame.uid"
            :action="selectedAction"
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
import { ActionGroup } from "@/common/constants";

export default Vue.extend({
  name: "Toolbar",
  components: {
    EditKeyboardActionDialog
  },
  data() {
    return {
      expanded: false,
      actionItems: initialActionItems
    };
  },
  computed: {
    ...mapGetters(["frame", "selectedAction", "selectedActionGroup"]),
    showAction(): boolean {
      return !!this.frame.actions?.length;
    },
    actionDialogComponent() {
      switch (this.selectedActionGroup) {
        case ActionGroup.Keyboard:
          return "EditKeyboardActionDialog";
        default:
          return null;
      }
    },
    selectedActionType: {
      get(): number {
        return this.selectedAction?.actionType;
      },
      set(newValue: number) {
        this.updateFrames({
          frames: [
            {
              uid: this.frame.uid,
              actions: [
                {
                  uid: this.selectedAction.uid,
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
      updateFrames: ActionTypes.UPDATE_FRAMES
    }),
    handleToggleExpanded() {
      this.expanded = !this.expanded;
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
