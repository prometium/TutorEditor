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

type Frame {
	pictureLink
	actions
	taskText
	hintText
}

type Script {
	name
	firstFrame
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
pictureLink: string .
actions: [uid] .
hintText: string .
taskText: string .
name: string @index(exact) .
firstFrame: uid .
frames: [uid] .
`
