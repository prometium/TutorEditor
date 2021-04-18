import { GetterTree } from "vuex";
import { State } from "./state";
import { Frame, PathItem } from "@/common/types";
import { ActionGroup, ActionType } from "@/common/constants";

type Getters = {
  path(state: State): Array<PathItem>;
  frame(state: State): Frame;
};

export const getters: GetterTree<State, State> & Getters = {
  path(state) {
    if (state.script == null) return [];
    return state.script.path;
  },
  frame(state) {
    return state.script.frameByUid[state.frameUid] || {}
  },
  selectedAction(state, getters) {
    return getters.frame.actions && getters.frame.actions[
      state.script.branchNumByUid[getters.frame.uid] || 0
    ];
  },
  selectedActionGroup(_, getters) {
    switch (getters.selectedAction?.actionType) {
      case ActionType.LeftMouseClick:
      case ActionType.LeftMouseDown:
      case ActionType.LeftMouseUp:
      case ActionType.LeftMouseDoubleClick:
      case ActionType.RightMouseClick:
      case ActionType.RightMouseDown:
      case ActionType.RightMouseUp:
      case ActionType.RightMouseDobleClick:
        return ActionGroup.Mouse;
      case ActionType.KeyClick:
      case ActionType.KeyDown:
      case ActionType.KeyUp:
      case ActionType.KeyWithMod:
        return ActionGroup.Keyboard;
      default:
        return ActionGroup.Other;
    }
  }
};
