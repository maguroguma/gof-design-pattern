package observer

type Observer interface {
	update()
}

type Subject interface {
	addObserver(Observer) // 本来ならばスーパークラスで実装する
	notifyObservers()     // 本来ならばスーパークラスで実装する
	done()
}
