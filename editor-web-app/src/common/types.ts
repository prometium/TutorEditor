export interface ScriptInfo {
  uid: string;
  name: string;
  version: string;
}

export interface Script {
  uid: string;
  name: string;
  version: string;
  firstFrame: {
    uid: string;
  };
  frames: Array<Frame>;
}

export interface TraversableScript extends Omit<Script, "frames"> {
  frameByUid: Record<string, Frame>;
  path: Array<PathItem>;
  branchNumByUid: Record<string, number>;
}

export interface PathItem {
  frameUid: string;
  branchNumber: number;
}

export interface Frame {
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
  taskText: string;
  hintText: string;
}

export interface Action {
  uid: string;
  nextFrame: {
    uid: string;
  };
}

interface MouseAction extends Action {
  xLeft: number;
  xRight: number;
  yLeft: number;
  yRight: number;
}

interface LeftMouseClick extends MouseAction {
  actionType: 1;
}
interface LeftMouseDown extends MouseAction {
  actionType: 2;
}
interface LeftMouseUp extends MouseAction {
  actionType: 3;
}
interface LeftMouseDoubleClick extends MouseAction {
  actionType: 4;
}
interface RightMouseClick extends MouseAction {
  actionType: 5;
}
interface RightMouseDown extends MouseAction {
  actionType: 6;
}
interface RightMouseUp extends MouseAction {
  actionType: 7;
}
interface RightMouseDobleClick extends MouseAction {
  actionType: 8;
}

interface KeyboardAction extends Action {
  key: string;
}
interface KeyClick extends KeyboardAction {
  actionType: 9;
}
interface KeyDown extends KeyboardAction {
  actionType: 10;
}
interface KeyUp extends KeyboardAction {
  actionType: 11;
}
interface KeyWithMod extends KeyboardAction {
  actionType: 12;
  modKey: string;
}

interface Drag extends Action {
  actionType: 13;
  startXLeft: number;
  startYLeft: number;
  startXRight: number;
  startYRight: number;
  finishXLeft: number;
  finishYLeft: number;
  finishXRight: number;
  finishYRight: number;
}

interface WheelMotionAction extends Action {
  ticksCount: number;
}
interface WheelUp extends WheelMotionAction {
  actionType: 14;
}
interface WheelDown extends WheelMotionAction {
  actionType: 15;
}

interface WheelClick extends Action {
  actionType: 16;
}

interface Pause extends Action {
  actionType: 17;
  duration: number;
}
