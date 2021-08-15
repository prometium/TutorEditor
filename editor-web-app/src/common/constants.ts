export enum ActionType {
    LeftMouseClick = 1,
    LeftMouseDown,
    LeftMouseUp,
    LeftMouseDoubleClick,
    RightMouseClick,
    RightMouseDown,
    RightMouseUp,
    RightMouseDobleClick,
    KeyClick,
    KeyDown,
    KeyUp,
    KeyWithMod,
    Drag,
    WheelUp,
    WheelDown,
    WheelClick,
    Pause
}

export enum ActionGroup {
    Mouse,
    Keyboard,
    Pause,
    Other
}