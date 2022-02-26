package implementation

import (
	"archive/zip"
	"context"
	"encoding/json"
	"io"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"sync"

	"golang.org/x/sync/errgroup"

	"github.com/prometium/tutoreditor/editorsvc"
	"github.com/prometium/tutoreditor/editorsvc/utils"
)

type switchPicture struct {
	PictureLink string  `json:"pictureLink,omitempty"`
	X           float32 `json:"x,omitempty"`
	Y           float32 `json:"y,omitempty"`
}

type rawMouseAction struct {
	XLeft  float32 `json:"xLeft,omitempty"`
	XRight float32 `json:"xRight,omitempty"`
	YLeft  float32 `json:"yLeft,omitempty"`
	YRight float32 `json:"yRight,omitempty"`
}

type rawDragAction struct {
	StartXLeft   float32 `json:"startXLeft,omitempty"`
	StartYLeft   float32 `json:"startYLeft,omitempty"`
	StartXRight  float32 `json:"startXRight,omitempty"`
	StartYRight  float32 `json:"startYRight,omitempty"`
	FinishXLeft  float32 `json:"finishXLeft,omitempty"`
	FinishYLeft  float32 `json:"finishYLeft,omitempty"`
	FinishXRight float32 `json:"finishXRight,omitempty"`
	FinishYRight float32 `json:"finishYRight,omitempty"`
}

type rawWheelAction struct {
	TicksCount int `json:"ticksCount,omitempty"`
}

type rawKeyboardAction struct {
	Key    string `json:"key,omitempty"`
	ModKey string `json:"modKey,omitempty"`
}

type rawPauseAction struct {
	Duration float32 `json:"duration,omitempty"`
}

type rawAction struct {
	ActionType int `json:"actionId,omitempty"`
	rawMouseAction
	rawDragAction
	rawWheelAction
	rawKeyboardAction
	rawPauseAction
	SwitchPictures []switchPicture `json:"switchPictures,omitempty"`
}

type rawFrame struct {
	FrameNumber  int       `json:"frameNumber,omitempty"`
	PictureLink  string    `json:"pictureLink,omitempty"`
	ActionSwitch rawAction `json:"actionSwitch,omitempty"`
	Task         string    `json:"task,omitempty"`
	Hint         string    `json:"hint,omitempty"`
}

type rawScriptArchiveSaver struct {
	Images        []*zip.File
	PictureWidth  int        `json:"pictureWidth,omitempty"`
	PictureHeight int        `json:"pictureHeight,omitempty"`
	Frames        []rawFrame `json:"frames,omitempty"`
}

// https://docs.microsoft.com/ru-ru/dotnet/api/system.windows.forms.keys?view=windowsdesktop-6.0
// https://developer.mozilla.org/en-US/docs/Web/API/KeyboardEvent/code/code_values
var codeByWindowsKey = map[string]string{
	"A":               "KeyA",
	"B":               "KeyB",
	"C":               "KeyC",
	"D":               "KeyD",
	"E":               "KeyE",
	"F":               "KeyF",
	"G":               "KeyG",
	"H":               "KeyH",
	"I":               "KeyI",
	"J":               "KeyJ",
	"K":               "KeyK",
	"L":               "KeyL",
	"M":               "KeyM",
	"N":               "KeyN",
	"O":               "KeyO",
	"P":               "KeyP",
	"Q":               "KeyQ",
	"R":               "KeyR",
	"S":               "KeyS",
	"T":               "KeyT",
	"U":               "KeyU",
	"V":               "KeyV",
	"W":               "KeyW",
	"X":               "KeyX",
	"Y":               "KeyY",
	"Z":               "KeyZ",
	"D0":              "Digit0",
	"D1":              "Digit1",
	"D2":              "Digit2",
	"D3":              "Digit3",
	"D4":              "Digit4",
	"D5":              "Digit5",
	"D6":              "Digit6",
	"D7":              "Digit7",
	"D8":              "Digit8",
	"D9":              "Digit9",
	"NumPad0":         "Numpad0",
	"NumPad1":         "Numpad1",
	"NumPad2":         "Numpad2",
	"NumPad3":         "Numpad3",
	"NumPad4":         "Numpad4",
	"NumPad5":         "Numpad5",
	"NumPad6":         "Numpad6",
	"NumPad7":         "Numpad7",
	"NumPad8":         "Numpad8",
	"NumPad9":         "Numpad9",
	"OemMinus":        "Minus",
	"OemPlus":         "Equal",
	"Oem1":            "Semicolon",
	"Oem3":            "Backquote",
	"Oem5":            "Backslash",
	"Oem6":            "BracketRight",
	"OemOpenBrackets": "BracketLeft",
	"OemQuotes":       "Quote",
	"OemComma":        "Comma",
	"OemPeriod":       "Period",
	"OemQuestion":     "Slash",
	"Back":            "Backspace",
	"LeftCtrl":        "ControlLeft",
	"LeftAlt":         "AltLeft",
	"LeftShift":       "ShiftLeft",
	"LWin":            "MetaLeft",
	"RightCtrl":       "ControlRight",
	"RightAlt":        "AltRight",
	"RightShift":      "ShiftRight",
	"RWin":            "MetaRight",
	"Return":          "Enter",
	"Capital":         "CapsLock",
	"Next":            "PageDown",
	"Up":              "ArrowUp",
	"Down":            "ArrowDown",
	"Left":            "ArrowLeft",
	"Right":           "ArrowRight",
}

func (controller *rawScriptArchiveSaver) init(r io.Reader) error {
	zipReader, err := utils.CreateZipReader(r)
	if err != nil {
		return err
	}

	controller.Images = make([]*zip.File, 0, len(zipReader.File))
	var scriptFile *zip.File = nil
	for _, file := range zipReader.File {
		if filepath.Ext(strings.TrimSpace(file.Name)) == ".png" {
			controller.Images = append(controller.Images, file)
		} else if file.Name == "Script.json" {
			scriptFile = file
		}
	}

	scriptJSON, err := utils.ReadAllFromZip(scriptFile)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(scriptJSON, &controller); err != nil {
		return err
	}

	return nil
}

func (controller *rawScriptArchiveSaver) saveImages(ctx context.Context, imagesDir string) (map[string]string, error) {
	os.MkdirAll(imagesDir, os.ModePerm)

	lock := sync.RWMutex{}
	errs, _ := errgroup.WithContext(ctx)
	var linksMap map[string]string = make(map[string]string)
	for _, file := range controller.Images {
		currentFile := file
		errs.Go(func() error {
			hash, err := utils.HashZipFileMD5(currentFile)
			if err != nil {
				return err
			}

			path := filepath.Join(imagesDir, hash+".png")
			if _, err := os.Stat(path); os.IsNotExist(err) {
				err = utils.CopyZipFile(currentFile, path)
				if err != nil {
					return err
				}
			}

			lock.Lock()
			defer lock.Unlock()
			linksMap[currentFile.Name] = hash + ".png"

			return nil
		})
	}

	if err := errs.Wait(); err != nil {
		return nil, err
	}
	return linksMap, nil
}

func (controller *rawScriptArchiveSaver) createScript(name string, linksMap map[string]string) (*editorsvc.Script, error) {
	frames := make([]editorsvc.Frame, len(controller.Frames))

	for i, frame := range controller.Frames {
		action := &frame.ActionSwitch

		var nextFrame *editorsvc.NextFrame
		if i+1 < len(controller.Frames) {
			nextFrame = &editorsvc.NextFrame{
				UID: strconv.Itoa(controller.Frames[i+1].FrameNumber),
			}
		}

		var key = action.Key
		adaptedKey, ok := codeByWindowsKey[action.Key]
		if ok {
			key = adaptedKey
		}

		frames[i] = editorsvc.Frame{
			UID:         strconv.Itoa(frame.FrameNumber),
			PictureLink: linksMap[frame.PictureLink],
			TaskText:    frame.Task,
			HintText:    frame.Hint,
			Actions: []editorsvc.Action{
				{
					NextFrame:    nextFrame,
					ActionType:   action.ActionType,
					XLeft:        action.XLeft,
					XRight:       action.XRight,
					YLeft:        action.YLeft,
					YRight:       action.YRight,
					StartXLeft:   action.StartXLeft,
					StartYLeft:   action.StartYLeft,
					StartXRight:  action.StartXRight,
					StartYRight:  action.StartYRight,
					FinishXLeft:  action.FinishXLeft,
					FinishYLeft:  action.FinishYLeft,
					FinishXRight: action.FinishXRight,
					FinishYRight: action.FinishYRight,
					TicksCount:   action.TicksCount,
					Key:          key,
					ModKey:       action.ModKey,
					Duration:     action.Duration,
				},
			},
		}

		if action.SwitchPictures == nil {
			continue
		}

		frames[i].Actions[0].SwitchPictures = make([]editorsvc.SwitchPicture, len(action.SwitchPictures))
		for j, switchPicture := range action.SwitchPictures {
			frames[i].Actions[0].SwitchPictures[j] = editorsvc.SwitchPicture{
				PictureLink: linksMap[switchPicture.PictureLink],
				X:           switchPicture.X,
				Y:           switchPicture.Y,
			}
		}
	}

	script := editorsvc.Script{
		Name:          name,
		PictureWidth:  controller.PictureWidth,
		PictureHeight: controller.PictureHeight,
		Frames:        frames,
	}

	if len(frames) > 0 {
		script.FirstFrame = &editorsvc.NextFrame{
			UID: frames[0].UID,
		}
	}

	return &script, nil
}
