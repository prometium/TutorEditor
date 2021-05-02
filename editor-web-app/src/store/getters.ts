import { State } from "./state";
import { Action, Frame, PathItem } from "@/common/types";
import { ActionGroup, ActionType } from "@/common/constants";

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
    let frameUid = state.script.firstFrame.uid;
    const path: PathItem[] = [];
    while (path.length <= Object.keys(state.script.frameByUid).length) {
      const pathItem: PathItem = {
        frameUid,
        branchNum: state.script.branchNumByUid[frameUid] || 0
      };
      path.push(pathItem);

      const actions = state.script.frameByUid[frameUid].actions;
      if (actions == null || !actions.length) break;

      const nextFrame = actions[pathItem.branchNum].nextFrame;
      if (!nextFrame) break;

      frameUid = nextFrame.uid;
    }

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
      default:
        return ActionGroup.Other;
    }
  },
  prevPathItem(_, getters) {
    const pathItemIndex = getters.path.findIndex(
      (pathItem: PathItem) => pathItem.frameUid === getters.currentFrame?.uid
    );
    const prevPathItem = getters.path[pathItemIndex - 1];
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
    if (!getters.currentAction?.nextFrame.uid) return null;
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
