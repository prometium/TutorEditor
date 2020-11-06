package dgraphdb

var schema = `
type Action {
	actionType
	nextFrame
	xLeft
	xRight
	yLeft
	yRight
	startXLeft
	startYLeft
	startXRight
	startYRight
	finishXLeft
	finishYLeft
	finishXRight
	finishYRight
	ticksCount
	key
	modKey
}

type Task {
	text
}

type Hint {
	text
}

type Frame {
	pictureLink
	actions
	task
	hint
}

type Script {
	name
	frames
}

actionType: int .
nextFrame: uid .
xLeft: float .
xRight: float .
yLeft: float .
yRight: float .
startXLeft: float .
startYLeft: float .
startXRight: float .
startYRight: float .
finishXLeft: float .
finishYLeft: float .
finishXRight: float .
finishYRight: float .
ticksCount: int .
key: string .
modKey: string .
text: string .
pictureLink: string .
actions: [uid] .
task: uid .
hint: uid .
name: string @index(exact) .
frames: [uid] .
`
