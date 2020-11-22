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
	XLeft  float32 `json:"xLeft,omitempty"`
	XRight float32 `json:"xRight,omitempty"`
	YLeft  float32 `json:"yLeft,omitempty"`
	YRight float32 `json:"yRight,omitempty"`

	// drag
	StartXLeft   float32 `json:"startXLeft,omitempty"`
	StartYLeft   float32 `json:"startYLeft,omitempty"`
	StartXRight  float32 `json:"startXRight,omitempty"`
	StartYRight  float32 `json:"startYRight,omitempty"`
	FinishXLeft  float32 `json:"finishXLeft,omitempty"`
	FinishYLeft  float32 `json:"finishYLeft,omitempty"`
	FinishXRight float32 `json:"finishXRight,omitempty"`
	FinishYRight float32 `json:"finishYRight,omitempty"`

	// wheel
	TicksCount int `json:"ticksCount,omitempty"`

	// keyboard
	Key    string `json:"key,omitempty"`
	ModKey string `json:"modKey,omitempty"`

	SwitchPictures []SwitchPicture `json:"switchPictures,omitempty"`

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
	UID        string     `json:"uid,omitempty"`
	Name       string     `json:"name,omitempty"`
	Version    string     `json:"version,omitempty"`
	FirstFrame *NextFrame `json:"firstFrame,omitempty"`
	Frames     []Frame    `json:"frames,omitempty"`
	DType      []string   `json:"dgraph.type,omitempty"`
}

// Branch represents a branch
type Branch struct {
	FirstMainFrameID string  `json:"firstMainFrameId,omitempty"`
	LastMainFrameID  string  `json:"lastMainFrameId,omitempty"`
	ConnectedFrames  []Frame `json:"connectedFrames,omitempty"`
}

// BranchToDelete represents a branch to be deleted
type BranchToDelete struct {
	RootFrameID   string `json:"rootFrameId,omitempty"`
	FirstActionID string `json:"firstActionId,omitempty"`
	LastActionID  string `json:"lastActionId,omitempty"`
}

// Repository describes the persistence on editor model
type Repository interface {
	Setup(ctx context.Context) error
	AddScript(ctx context.Context, script *Script) (string, error)
	GetScriptsList(ctx context.Context) ([]Script, error)
	GetScript(ctx context.Context, id string) ([]Script, error)
	DeleteScript(ctx context.Context, id string) error
	UpdateScript(ctx context.Context, script *Script) (map[string]string, error)
	AddBranch(ctx context.Context, script *Script, branch *Branch) (map[string]string, error)
	DeleteBranch(ctx context.Context, script *Script, branchToDelete *BranchToDelete) error
	AddFrame(ctx context.Context, script *Script, framesPair []Frame) (map[string]string, error)
	DeleteFrame(ctx context.Context, script *Script, id string) error
	GetScriptVersion(ctx context.Context, id string) (string, error)
}
