package prototype

// ClonableObjectを名前付きで管理する構造体
// ClonableObjectの具象構造体が一切登場しないことに注目
type Manager struct {
	showcase map[string]ClonableObject
}

func NewManager() *Manager {
	m := new(Manager)
	m.showcase = make(map[string]ClonableObject)
	return m
}

func (m *Manager) Register(name string, proto ClonableObject) {
	m.showcase[name] = proto
}

func (m *Manager) Create(name string) ClonableObject {
	p := m.showcase[name]
	return p.CreateClone() // 登録されたものではなく、そのCloneを返す
}
