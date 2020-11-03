package implementation

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

type keyboardAction struct {
	Key    string `json:"key,omitempty"`
	ModKey string `json:"modKey,omitempty"`
}

type action struct {
	ActionID int `json:"ActionID"`
	mouseAction
	dragAction
	wheelAction
	keyboardAction
}

type frame struct {
	FrameNumber  int    `json:"frameNumber"`
	PictureLink  string `json:"pictureLink"`
	ActionSwitch action `json:"action"`
	Task         string `json:"task,omitempty"`
	Hint         string `json:"hint,omitempty"`
}

type script struct {
	Frames []frame `json:"frames"`
}
