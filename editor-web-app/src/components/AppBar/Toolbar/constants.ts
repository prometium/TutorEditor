import { ActionType } from "@/common/constants";

export const initialActionItems = [
    {
        value: ActionType.LeftMouseClick,
        text: "Левый щелчек мышью"
    },
    {
        value: ActionType.RightMouseClick,
        text: "Правый щелчек мышью"
    },
    {
        value: ActionType.LeftMouseUp,
        text: "Левое отжатие мышью"
    },
    {
        value: ActionType.LeftMouseDoubleClick,
        text: "Двойной левый щелчок мышью"
    },
    {
        value: ActionType.RightMouseClick,
        text: "Правый щелчок мышью"
    },
    {
        value: ActionType.RightMouseDown,
        text: "Правое нажатие мышью"
    },
    {
        value: ActionType.RightMouseDown,
        text: "Правое отжатие мышью"
    },
    {
        value: ActionType.RightMouseDobleClick,
        text: "Двойной правый щелчок мышью"
    },
    {
        value: ActionType.KeyClick,
        text: "Щелчок кнопки на клавиатуре"
    },
    {
        value: ActionType.KeyDown,
        text: "Нажатие кнопки на клавиатуре"
    },
    {
        value: ActionType.KeyUp,
        text: "Отжатие кнопки на клавиатуре"
    },
    {
        value: ActionType.KeyWithMod,
        text: "Кнопка на клавиатуре с модификатором"
    },
    {
        value: ActionType.Drag,
        text: "Перетаскивание"
    },
    {
        value: ActionType.WheelUp,
        text: "Прокрутка колесиком вверх"
    },
    {
        value: ActionType.WheelDown,
        text: "Прокрутка колесиком вниз"
    },
    {
        value: ActionType.WheelClick,
        text: "Нажатие на колесико"
    },
    {
        value: ActionType.Pause,
        text: "Пауза"
    }
];