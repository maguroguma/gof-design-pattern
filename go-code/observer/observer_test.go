package observer

import "testing"

func TestObserverPattern(t *testing.T) {
	t1, t2 := NewTeacher("Tanaka"), NewTeacher("Yamada")
	s1, s2, s3, s4 := NewStudent("A"), NewStudent("B"), NewStudent("C"), NewStudent("D")

	s2.addObserver(t1)
	s3.addObserver(t2)
	s4.addObserver(t1)
	s4.addObserver(t2)

	s1.done()
	s2.done()
	s4.done()

	if !(t1.doneNum == 2 && t2.doneNum == 1) {
		t.Errorf("expect (t1, t2) = (2, 1), got (%d, %d)", t1.doneNum, t2.doneNum)
	}
}
