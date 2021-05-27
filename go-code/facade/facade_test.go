package facade

import "testing"

func TestFacadePattern(t *testing.T) {
	f := NewFacade()
	actual := f.Get()

	if actual != "(HTTP経由で得られた何か) と (DBから得られた何か) で作られた何か" {
		t.Errorf("Error")
	}
}
