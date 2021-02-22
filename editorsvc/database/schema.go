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
	switchPicturesJSON
}

type Frame {
	pictureLink
	actions
	taskText
	hintText
}

type Script {
	name
	version
	firstFrame
	frames
}

actionType: int .
nextFrame: uid .
x: float .
y: float .
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
switchPicturesJSON: string .
pictureLink: string .
dragXCoords: [float] .
dragYCoords: [float] .
actions: [uid] @reverse .
hintText: string .
taskText: string .
name: string @index(exact) .
version: string @index(exact) .
firstFrame: uid .
frames: [uid] .
`