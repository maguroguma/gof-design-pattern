package proxy

type ProxyServer struct {
	real *RealServer
}

func NewProxyServer() *ProxyServer {
	ps := new(ProxyServer)
	ps.real = nil
	return ps
}

func (ps *ProxyServer) RightRequest() string {
	return "やるだけ"
}

func (ps *ProxyServer) HeavyRequest() string {
	ps.realize()
	return ps.real.HeavyRequest()
}

func (ps *ProxyServer) realize() {
	if ps.real == nil {
		ps.real = NewRealServer()
	}
}
