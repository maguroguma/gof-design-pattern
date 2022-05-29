package frontend

import (
	"reflect"
	"testing"
)

func TestFrontendObserverPattern(t *testing.T) {
	ob := NewObservable()

	beforeResult := ob.notify("before")
	if !reflect.DeepEqual(beforeResult, []string{}) {
		t.Errorf("got %v, want %v", beforeResult, []string{})
	}

	ob.subscribe(logger)
	ob.subscribe(toastify)
	afterResult := ob.notify("after")
	expected := []string{"INFO: after", "notified by toast ui: after"}
	if !reflect.DeepEqual(afterResult, expected) {
		t.Errorf("got %v, want %v", afterResult, expected)
	}
}
