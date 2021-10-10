package decorator

import (
	"reflect"
	"testing"
)

var isLocked = false

func Lock() {
	isLocked = true
}
func UnLock() {
	isLocked = false
}

func Do() string {
	return "Done"
}

func CheckLockDecorator(f func() string) func() string {
	return func() string {
		if isLocked {
			return "Failed by lock"
		}

		return f()
	}
}

func TestDecoratedFunction(t *testing.T) {
	doIfPossible := CheckLockDecorator(Do)

	actuals := []string{}
	actuals = append(actuals, doIfPossible())
	Lock()
	actuals = append(actuals, doIfPossible())
	UnLock()
	actuals = append(actuals, doIfPossible())

	expected := []string{"Done", "Failed by lock", "Done"}

	if !reflect.DeepEqual(actuals, expected) {
		t.Errorf("got %v, want %v", actuals, expected)
	}
}
