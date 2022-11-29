import type { ActionType } from "./constants";

export type ScriptInfo = {
  uid: string;
  name: string;
  version: string;
};

export type Script = {
  uid: string;
  name: string;
  version: string;
  firstFrame: {
    uid: string;
  };
  frames: Frame[];
};

export type TraversableScript = Omit<Script, "frames"> & {
  frameByUid: Record<string, Frame>;
  path: PathItem[];
  branchNumByUid: Record<string, number>;
};

export type PathItem = {
  frameUid: string;
  branchNum: number;
};

export type Frame = {
  uid: string;
  pictureLink?: string;
  actions?: Array<Action>;
  taskText?: string;
  hintText?: string;
};

export type Action = {
  actionType?: ActionType;
  uid: string;
  nextFrame?: {
    uid: string;
  };
  key?: string;
  modKey?: string;
  xLeft?: number;
  xRight?: number;
  yLeft?: number;
  yRight?: number;
  startXLeft?: number;
  startYLeft?: number;
  startXRight?: number;
  startYRight?: number;
  finishXLeft?: number;
  finishYLeft?: number;
  finishXRight?: number;
  finishYRight?: number;
  switchPictures?: SwitchPicture[];
  duration?: number;
  ticksCount?: number;
};

type MouseAction = Action & {
  xLeft: number;
  xRight: number;
  yLeft: number;
  yRight: number;
};

export type LeftMouseClick = MouseAction & {
  actionType: ActionType.LeftMouseClick;
};
export type LeftMouseDown = MouseAction & {
  actionType: ActionType.LeftMouseDown;
};
export type LeftMouseUp = MouseAction & {
  actionType: ActionType.LeftMouseUp;
};
export type LeftMouseDoubleClick = MouseAction & {
  actionType: ActionType.LeftMouseDoubleClick;
};
export type RightMouseClick = MouseAction & {
  actionType: ActionType.RightMouseClick;
};
export type RightMouseDown = MouseAction & {
  actionType: ActionType.RightMouseDown;
};
export type RightMouseUp = MouseAction & {
  actionType: ActionType.RightMouseUp;
};
export type RightMouseDobleClick = MouseAction & {
  actionType: ActionType.RightMouseDobleClick;
};

type KeyboardAction = Action & {
  key: string;
};
export type KeyClick = KeyboardAction & {
  actionType: ActionType.KeyClick;
};
export type KeyDown = KeyboardAction & {
  actionType: ActionType.KeyDown;
};
export type KeyUp = KeyboardAction & {
  actionType: ActionType.KeyUp;
};
export type KeyWithMod = KeyboardAction & {
  actionType: ActionType.KeyWithMod;
  modKey: string;
};

export type Drag = Action & {
  actionType: ActionType.Drag;
  startXLeft: number;
  startYLeft: number;
  startXRight: number;
  startYRight: number;
  finishXLeft: number;
  finishYLeft: number;
  finishXRight: number;
  finishYRight: number;
  switchPictures: SwitchPicture;
};

export type SwitchPicture = {
  pictureLink?: string;
  x: number;
  y: number;
};

type WheelMotionAction = Action & {
  ticksCount: number;
};
export type WheelUp = WheelMotionAction & {
  actionType: ActionType.WheelUp;
};
export type WheelDown = WheelMotionAction & {
  actionType: ActionType.WheelDown;
};

export type WheelClick = Action & {
  actionType: ActionType.WheelClick;
};

export type Pause = Action & {
  actionType: ActionType.Pause;
  duration?: number;
};
