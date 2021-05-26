package strategy

import "testing"

func TestStrategyPatternByDependencyInjection(t *testing.T) {
	carefulP := NewPlayer(NewCarefulStrategy())
	if carefulP.Do() != "宿題は計画的に -> 初日に終わった" {
		t.Errorf("Error")
	}
	shallowP := NewPlayer(NewShallowStrategy())
	if shallowP.Do() != "宿題は最後の一週間に -> 助けて" {
		t.Errorf("Error")
	}
}

func TestStrategyPatternByFirstClassFunction(t *testing.T) {
	p := NewPlayer(nil)
	actual := p.AdLib(func() string { return "ケセラセラ" })
	if actual != "戦略: ケセラセラ" {
		t.Errorf("Error")
	}
}
