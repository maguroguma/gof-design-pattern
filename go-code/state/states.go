package state

type dayState struct{}

func (ds *dayState) doClock(context *MainProgram, hour int) {
	if hour < 9 || 17 <= hour {
		context.changeState(&nightState{})
	}
}

func (ds *dayState) doUse(context *MainProgram) {
	context.recordLog("エンカウント率: 低")
}

type nightState struct{}

func (ns *nightState) doClock(context *MainProgram, hour int) {
	if !(hour < 9 || 17 <= hour) {
		context.changeState(&dayState{})
	}
}

func (ns *nightState) doUse(context *MainProgram) {
	context.recordLog("エンカウント率: 高")
}
