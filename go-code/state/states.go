package state

type dayState struct{}

func (ds *dayState) doClock(context context, hour int) {
	if hour < 9 || 17 <= hour {
		context.changeState(&nightState{})
	}
}

func (ds *dayState) doUse(context context) {
	context.recordLog("エンカウント率: 低")
}

type nightState struct{}

func (ns *nightState) doClock(context context, hour int) {
	if !(hour < 9 || 17 <= hour) {
		context.changeState(&dayState{})
	}
}

func (ns *nightState) doUse(context context) {
	context.recordLog("エンカウント率: 高")
}
