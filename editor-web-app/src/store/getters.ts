import { State } from "./state";
import { Action, Frame, PathItem } from "@/common/types";
import { ActionGroup, ActionType } from "@/common/constants";

type Getters = {
  path(state: State): Array<PathItem>;
  frame(state: State): Frame;
  selectedAction(state: State, getters: { [T in keyof Getters]: ReturnType<Getters[T]> }): Action | null;
  selectedActionGroup(state: State, getters: { [T in keyof Getters]: ReturnType<Getters[T]> }): number;
};

export const getters: Getters = {
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
    ] || null;
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
