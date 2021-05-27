package state

type MainProgram struct {
	State state
	logs  []string
}

func NewMainProgram(state state) *MainProgram {
	mp := new(MainProgram)
	mp.State = state
	mp.logs = []string{}
	return mp
}

func (mp *MainProgram) SetClock(hour int) {
	mp.State.doClock(mp, hour)
}

func (mp *MainProgram) changeState(state state) {
	mp.State = state
}

func (mp *MainProgram) recordLog(log string) {
	mp.logs = append(mp.logs, log)
}

func (mp *MainProgram) Use() {
	mp.State.doUse(mp)
}

func (mp *MainProgram) GetLogs() []string {
	return mp.logs
}
