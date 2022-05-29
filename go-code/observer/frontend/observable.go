package frontend

// observerを購読するobservable構造体

type EventHandler func(event string) string

type Observable struct {
	observers []EventHandler
}

func NewObservable() *Observable {
	res := new(Observable)
	res.observers = []EventHandler{}
	return res
}

// unsubscribeは実装が面倒なので省略

func (ob *Observable) subscribe(handler EventHandler) {
	ob.observers = append(ob.observers, handler)
}

func (ob *Observable) notify(e string) []string {
	res := []string{}

	for _, handler := range ob.observers {
		res = append(res, handler(e))
	}

	return res
}
