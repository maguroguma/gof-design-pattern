package observer

type Teacher struct {
	name    string
	doneNum int
}

func NewTeacher(name string) *Teacher {
	t := new(Teacher)
	t.name = name
	t.doneNum = 0
	return t
}

func (t *Teacher) update() {
	t.doneNum++
}
