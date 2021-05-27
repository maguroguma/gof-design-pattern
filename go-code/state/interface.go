package state

// 「状態」を抽象化したもの
// 振る舞いを定義する
type state interface {
	doClock(context context, hour int)
	doUse(context context)
}

// 抽象化した「状態」を扱ってプログラムを構築する
// 実装クラスで「状態」をhasすることになる
type context interface {
	SetClock(hour int)
	changeState(state state)
	recordLog(log string)
}
