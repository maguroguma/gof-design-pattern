package proxy

import "time"

type RealServer struct{}

// RealServerは重たい（と想定して）必要になったタイミングで初めて生成したい
func NewRealServer() *RealServer {
	time.Sleep(1 * time.Second)
	return new(RealServer)
}

func (rs *RealServer) RightRequest() string {
	return "やるだけ"
}

func (rs *RealServer) HeavyRequest() string {
	return "構文解析"
}
