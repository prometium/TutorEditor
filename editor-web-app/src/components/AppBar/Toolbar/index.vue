<template>
  <div class="toolbar">
    <v-btn @click="toggleExpansion" elevation="1" icon :disabled="!showAction">
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
          <v-btn elevation="1" icon>
            <v-icon> mdi-cog </v-icon>
          </v-btn>
        </v-col>
      </v-row>
    </v-container>
  </div>
</template>

<script>
import Vue from "vue";
import { mapState, mapGetters, mapActions } from "vuex";
import { ActionTypes } from "@/store/action-types";
import { initialActionItems } from "./constants";

export default Vue.extend({
  name: "Toolbar",
  data() {
    return {
      expanded: false,
      actionItems: initialActionItems
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
    showAction() {
      return !!this.frame.actions?.length;
    },
    currentBranchNum() {
      return this.script.branchNumByUid[this.frame.uid] || 0;
    },
    selectedAction() {
      return this.frame.actions && this.frame.actions[this.currentBranchNum];
    },
    selectedActionType: {
      get() {
        return this.selectedAction?.actionType;
      },
      set(newValue) {
        this.updateFrames([
          {
            uid: this.frame.uid,
            actions: [
              {
                uid: this.selectedAction.uid,
                actionType: newValue
              }
            ]
          }
        ]);
      }
    }
  },
  watch: {
    selectedActionType(newValue) {
      this.updateFrames([
        {
          uid: this.frame.uid,
          actions: [
            {
              uid: this.selectedAction.uid,
              actionType: newValue
            }
          ]
        }
      ]);
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
