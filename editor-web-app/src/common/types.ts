import { ActionType } from "./constants"

export type ScriptInfo = {
  uid: string;
  name: string;
  version: string;
}

export type Script = {
  uid: string;
  name: string;
  version: string;
  firstFrame: {
    uid: string;
  };
  frames: Frame[];
}

export type TraversableScript = Omit<Script, "frames"> & {
  frameByUid: Record<string, Frame>;
  path: PathItem[];
  branchNumByUid: Record<string, number>;
}

export type PathItem = {
  frameUid: string;
  branchNum: number;
}

export type Frame = {
  uid: string;
  pictureLink: string;
  actions?: Array<
    | LeftMouseClick
    | LeftMouseDown
    | LeftMouseUp
    | LeftMouseDoubleClick
    | RightMouseClick
    | RightMouseDown
    | RightMouseUp
    | RightMouseDobleClick
    | KeyClick
    | KeyDown
    | KeyUp
    | KeyWithMod
    | Drag
    | WheelUp
    | WheelDown
    | WheelClick
    | Pause
  >;
  taskText?: string;
  hintText?: string;
}

export type Action = {
  actionType: ActionType,
  uid: string;
  nextFrame: {
    uid: string;
  };
}

type MouseAction = Action & {
  xLeft: number;
  xRight: number;
  yLeft: number;
  yRight: number;
}

type LeftMouseClick = MouseAction & {
  actionType: ActionType.LeftMouseClick;
}
type LeftMouseDown = MouseAction & {
  actionType: ActionType.LeftMouseDown;
}
type LeftMouseUp = MouseAction & {
  actionType: ActionType.LeftMouseUp;
}
type LeftMouseDoubleClick = MouseAction & {
  actionType: ActionType.LeftMouseDoubleClick;
}
type RightMouseClick = MouseAction & {
  actionType: ActionType.RightMouseClick;
}
type RightMouseDown = MouseAction & {
  actionType: ActionType.RightMouseDown;
}
type RightMouseUp = MouseAction & {
  actionType: ActionType.RightMouseUp;
}
type RightMouseDobleClick = MouseAction & {
  actionType: ActionType.RightMouseDobleClick;
}

type KeyboardAction = Action & {
  key: string;
}
type KeyClick = KeyboardAction & {
  actionType: ActionType.KeyClick;
}
type KeyDown = KeyboardAction & {
  actionType: ActionType.KeyDown;
}
type KeyUp = KeyboardAction & {
  actionType: ActionType.KeyUp;
}
type KeyWithMod = KeyboardAction & {
  actionType: ActionType.KeyWithMod;
  modKey: string;
}

type Drag = Action & {
  actionType: ActionType.Drag;
  startXLeft: number;
  startYLeft: number;
  startXRight: number;
  startYRight: number;
  finishXLeft: number;
  finishYLeft: number;
  finishXRight: number;
  finishYRight: number;
}

type WheelMotionAction = Action & {
  ticksCount: number;
}
type WheelUp = WheelMotionAction & {
  actionType: ActionType.WheelUp;
}
type WheelDown = WheelMotionAction & {
  actionType: ActionType.WheelDown;
}

type WheelClick = Action & {
  actionType: ActionType.WheelClick;
}

type Pause = Action & {
  actionType: ActionType.Pause;
  duration: number;
}
