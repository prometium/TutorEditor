package editorsvc

import (
	"context"
)

// NextFrame represents a next frame
type NextFrame struct {
	UID string `json:"uid,omitempty"`
}

// SwitchPicture represents a switch picture
type SwitchPicture struct {
	UID           string   `json:"uid,omitempty"`
	PictureNumber int      `json:"pictureNumber,omitempty"`
	PictureLink   string   `json:"pictureLink,omitempty"`
	X             float32  `json:"x,omitempty"`
	Y             float32  `json:"y,omitempty"`
	DType         []string `json:"dgraph.type,omitempty"`
}

// Action represents an action
type Action struct {
	UID        string     `json:"uid,omitempty"`
	ActionType int        `json:"actionType,omitempty"`
	NextFrame  *NextFrame `json:"nextFrame,omitempty"`

	// mouse
	XLeft  *float32 `json:"xLeft"`
	XRight *float32 `json:"xRight"`
	YLeft  *float32 `json:"yLeft"`
	YRight *float32 `json:"yRight"`

	// drag
	StartXLeft   *float32 `json:"startXLeft"`
	StartYLeft   *float32 `json:"startYLeft"`
	StartXRight  *float32 `json:"startXRight"`
	StartYRight  *float32 `json:"startYRight"`
	FinishXLeft  *float32 `json:"finishXLeft"`
	FinishYLeft  *float32 `json:"finishYLeft"`
	FinishXRight *float32 `json:"finishXRight"`
	FinishYRight *float32 `json:"finishYRight"`

	// wheel
	TicksCount int `json:"ticksCount,omitempty"`

	// keyboard
	Key    string `json:"key,omitempty"`
	ModKey string `json:"modKey,omitempty"`

	// pause
	Duration float32 `json:"duration,omitempty"`

	SwitchPictures     []SwitchPicture `json:"switchPictures,omitempty"`
	SwitchPicturesJSON string          `json:"switchPicturesJSON,omitempty"`

	DType []string `json:"dgraph.type,omitempty"`
}

// Frame represents a frame
type Frame struct {
	UID         string   `json:"uid,omitempty"`
	PictureLink string   `json:"pictureLink,omitempty"`
	Actions     []Action `json:"actions,omitempty"`
	TaskText    string   `json:"taskText,omitempty"`
	HintText    string   `json:"hintText,omitempty"`
	DType       []string `json:"dgraph.type,omitempty"`
}

// Script represents a script
type Script struct {
	UID              string     `json:"uid,omitempty"`
	Name             string     `json:"name,omitempty"`
	Version          string     `json:"version,omitempty"`
	ModificationDate string     `json:"modificationDate,omitempty"`
	PictureWidth     int        `json:"pictureWidth,omitempty"`
	PictureHeight    int        `json:"pictureHeight,omitempty"`
	FirstFrame       *NextFrame `json:"firstFrame,omitempty"`
	Frames           []Frame    `json:"frames,omitempty"`
	DType            []string   `json:"dgraph.type,omitempty"`
}

// Repository describes the persistence on editor model
type Repository interface {
	Setup(ctx context.Context) error
	AddScript(ctx context.Context, script *Script) (string, error)
	GetScriptsList(ctx context.Context) ([]Script, error)
	GetScript(ctx context.Context, id string) (*Script, error)
	DeleteScript(ctx context.Context, id string) error
	UpdateScript(ctx context.Context, script *Script, frameIdsToDel []string, actionIdsToDel []string) (map[string]string, error)
}
