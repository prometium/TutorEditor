import { State } from "./state";
import { Action, Frame, PathItem } from "@/common/types";
import { ActionGroup, ActionType } from "@/common/constants";
import { configurePath } from "@/helpers/configurePath";

type Getters = {
  path(state: State): PathItem[];
  currentFrame(state: State): Frame | null;
  currentAction(
    state: State,
    getters: { [T in keyof Getters]: ReturnType<Getters[T]> }
  ): Action | null;
  currentActionGroup(
    state: State,
    getters: { [T in keyof Getters]: ReturnType<Getters[T]> }
  ): number;
  currentPathItemIndex(
    state: State,
    getters: { [T in keyof Getters]: ReturnType<Getters[T]> }
  ): number;
  currentPathItem(
    state: State,
    getters: { [T in keyof Getters]: ReturnType<Getters[T]> }
  ): PathItem | null;
  prevPathItem(
    state: State,
    getters: { [T in keyof Getters]: ReturnType<Getters[T]> }
  ): PathItem | null;
  prevFrame(
    state: State,
    getters: { [T in keyof Getters]: ReturnType<Getters[T]> }
  ): Frame | null;
  prevAction(
    state: State,
    getters: { [T in keyof Getters]: ReturnType<Getters[T]> }
  ): Action | null;
  nextFrame(
    state: State,
    getters: { [T in keyof Getters]: ReturnType<Getters[T]> }
  ): Frame | null;
  nextAction(
    state: State,
    getters: { [T in keyof Getters]: ReturnType<Getters[T]> }
  ): Action | null;
};

export const getters: Getters = {
  path(state) {
    const path = configurePath(
      state.script.firstFrame,
      state.script.frameByUid,
      state.script.branchNumByUid
    );
    return path;
  },
  currentFrame(state) {
    return state.frameUid ? state.script.frameByUid[state.frameUid] : null;
  },
  currentAction(state, getters) {
    return (
      (getters.currentFrame?.actions &&
        getters.currentFrame.actions[
          state.script.branchNumByUid[getters.currentFrame.uid] || 0
        ]) ||
      null
    );
  },
  currentActionGroup(_, getters) {
    switch (getters.currentAction?.actionType) {
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
      case ActionType.WheelDown:
      case ActionType.WheelUp:
        return ActionGroup.Tick;
      case ActionType.Pause:
        return ActionGroup.Pause;
      default:
        return ActionGroup.Other;
    }
  },
  currentPathItemIndex(_, getters) {
    const pathItemIndex = getters.path.findIndex(
      (pathItem: PathItem) => pathItem.frameUid === getters.currentFrame?.uid
    );
    return pathItemIndex || 0;
  },
  currentPathItem(_, getters) {
    const currentPathItem = getters.path[getters.currentPathItemIndex];
    return currentPathItem || null;
  },
  prevPathItem(_, getters) {
    const prevPathItem = getters.path[getters.currentPathItemIndex - 1];
    return prevPathItem || null;
  },
  prevFrame(state, getters) {
    if (!getters.prevPathItem) return null;
    return state.script.frameByUid[getters.prevPathItem.frameUid];
  },
  prevAction(_, getters) {
    if (!getters.prevPathItem) return null;
    return getters.prevFrame?.actions?.[getters.prevPathItem.branchNum] || null;
  },
  nextFrame(state, getters) {
    if (!getters.currentAction?.nextFrame?.uid) return null;
    return (
      state.script.frameByUid[getters.currentAction?.nextFrame.uid] || null
    );
  },
  nextAction(_, getters) {
    const pathItemIndex = getters.path.findIndex(
      (pathItem: PathItem) => pathItem.frameUid === getters.currentFrame?.uid
    );
    const nextPathItem = getters.path[pathItemIndex + 1];
    if (!nextPathItem) return null;
    return getters.nextFrame?.actions?.[nextPathItem.branchNum] || null;
  }
};
