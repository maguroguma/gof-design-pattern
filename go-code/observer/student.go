package observer

type Student struct {
	name      string
	isDone    bool
	observers []Observer
}

func NewStudent(name string) *Student {
	s := new(Student)
	s.name = name
	s.isDone = false
	s.observers = []Observer{}
	return s
}

func (s *Student) addObserver(newOb Observer) {
	s.observers = append(s.observers, newOb)
}

func (s *Student) notifyObservers() {
	for _, ob := range s.observers {
		ob.update()
	}
}

func (s *Student) done() {
	s.isDone = true
	s.notifyObservers()
}
