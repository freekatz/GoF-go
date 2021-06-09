package prototype

import "fmt"

type Manager struct {
	showcase map[string]Prototype
}

func NewManager() *Manager {
	return &Manager{
		showcase: make(map[string]Prototype),
	}
}

func (m *Manager) Create(name string) (Prototype, error) {
	proto, ok := m.showcase[name]
	if !ok {
		return nil, fmt.Errorf("prototype does not exist with name: %s, please register firstly\n", name)
	}
	return proto.CreateClone(), nil
}

func (m *Manager) Register(name string, proto Prototype) error {
	if _, ok := m.showcase[name]; ok {
		return fmt.Errorf("prototype has exist with name: %s\n", name)
	}
	m.showcase[name] = proto
	return nil
}
