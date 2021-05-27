package proxy

import "testing"

func TestProxyPattern(t *testing.T) {
	var r Respondent
	r = NewProxyServer()

	// HeavyRequestを2回読んでもスリープ時間は約1回分で済む
	println(r.RightRequest())
	println(r.RightRequest())
	println(r.HeavyRequest())
	println(r.HeavyRequest())
}
