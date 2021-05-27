package state

import (
	"reflect"
	"testing"
)

func TestStatePattern(t *testing.T) {
	mainProgram := NewMainProgram(&dayState{}) // 最初はnightStateを渡しても問題ない

	hours := []int{8, 9, 16, 17}

	for _, hour := range hours {
		mainProgram.SetClock(hour)
		mainProgram.Use()
	}

	actual := mainProgram.GetLogs()
	expect := []string{
		"エンカウント率: 高",
		"エンカウント率: 低",
		"エンカウント率: 低",
		"エンカウント率: 高",
	}

	if !reflect.DeepEqual(actual, expect) {
		t.Errorf("Error")
	}
}
