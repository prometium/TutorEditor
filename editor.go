package editorsvc

import "context"

type mouseAction struct {
	XLeft  float32 `json:"xLeft,omitempty"`
	XRight float32 `json:"xRight,omitempty"`
	YLeft  float32 `json:"yLeft,omitempty"`
	YRight float32 `json:"yRight,omitempty"`
}

type dragAction struct {
	StartXLeft   float32 `json:"startXLeft,omitempty"`
	StartYLeft   float32 `json:"startYLeft,omitempty"`
	StartXRight  float32 `json:"startXRight,omitempty"`
	StartYRight  float32 `json:"startYRight,omitempty"`
	FinishXLeft  float32 `json:"finishXLeft,omitempty"`
	FinishYLeft  float32 `json:"finishYLeft,omitempty"`
	FinishXRight float32 `json:"finishXRight,omitempty"`
	FinishYRight float32 `json:"finishYRight,omitempty"`
}

type wheelAction struct {
	TicksCount int `json:"ticksCount,omitempty"`
}

type nextFrame struct {
	UID string `json:"uid,omitempty"`
}

type keyboardAction struct {
	Key    string `json:"key,omitempty"`
	ModKey string `json:"modKey,omitempty"`
}

// Action represents a script
type Action struct {
	UID       string    `json:"uid,omitempty"`
	NextFrame nextFrame `json:"next_frame,omitempty"`
	mouseAction
	dragAction
	wheelAction
	keyboardAction
	DType []string `json:"dgraph.type,omitempty"`
}

// Frame represents a script
type Frame struct {
	UID         string   `json:"uid,omitempty"`
	PictureLink string   `json:"pictureLink,omitempty"`
	Actions     []Action `json:"actions,omitempty"`
	Task        string   `json:"task,omitempty"`
	Hint        string   `json:"hint,omitempty"`
	DType       []string `json:"dgraph.type,omitempty"`
}

// Script represents a script
type Script struct {
	Frames []Frame `json:"frames"`
}

// Repository describes the persistence on editor model
type Repository interface {
	AddScript(ctx context.Context, frames []Frame) error
}
