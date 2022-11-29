import type { State } from "./state";
import type { Action, Frame, PathItem } from "../common/types";
import { ActionGroup, ActionType } from "../common/constants";
import { configurePath } from "../helpers/configurePath";

type Getters = {
  path(this: ComputedGetters, state: State): PathItem[];
  currentFrame(this: ComputedGetters, state: State): Frame | null;
  currentAction(this: ComputedGetters, state: State): Action | null;
  currentActionGroup(this: ComputedGetters, state: State): number;
  currentPathItemIndex(this: ComputedGetters, state: State): number;
  currentPathItem(this: ComputedGetters, state: State): PathItem | null;
  prevPathItem(this: ComputedGetters, state: State): PathItem | null;
  prevFrame(this: ComputedGetters, state: State): Frame | null;
  prevAction(this: ComputedGetters, state: State): Action | null;
  nextFrame(this: ComputedGetters, state: State): Frame | null;
  nextAction(this: ComputedGetters, state: State): Action | null;
};

export type ComputedGetters = {
  [Getter in keyof Getters]: ReturnType<Getters[Getter]>;
};

export const getters: Getters = {
  path: (state) => {
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
  currentAction(state) {
    return (
      this.currentFrame?.actions?.[
        state.script.branchNumByUid[this.currentFrame.uid] || 0
      ] || null
    );
  },
  currentActionGroup() {
    switch (this.currentAction?.actionType) {
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
  currentPathItemIndex() {
    const pathItemIndex = this.path.findIndex(
      (pathItem: PathItem) => pathItem.frameUid === this.currentFrame?.uid
    );
    return pathItemIndex || 0;
  },
  currentPathItem() {
    const currentPathItem = this.path[this.currentPathItemIndex];
    return currentPathItem || null;
  },
  prevPathItem() {
    const prevPathItem = this.path[this.currentPathItemIndex - 1];
    return prevPathItem || null;
  },
  prevFrame(state) {
    if (!this.prevPathItem) return null;
    return state.script.frameByUid[this.prevPathItem.frameUid];
  },
  prevAction() {
    if (!this.prevPathItem) return null;
    return this.prevFrame?.actions?.[this.prevPathItem.branchNum] || null;
  },
  nextFrame(state) {
    if (!this.currentAction?.nextFrame?.uid) return null;
    return state.script.frameByUid[this.currentAction?.nextFrame.uid] || null;
  },
  nextAction() {
    const pathItemIndex = this.path.findIndex(
      (pathItem: PathItem) => pathItem.frameUid === this.currentFrame?.uid
    );
    const nextPathItem = this.path[pathItemIndex + 1];
    if (!nextPathItem) return null;
    return this.nextFrame?.actions?.[nextPathItem.branchNum] || null;
  },
};
