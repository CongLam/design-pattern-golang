package machinemanager

import "sync"

type manager struct {
	state string
}

var singleton *manager
var once sync.Once

func GetMachine() *manager {
	once.Do(func() {
		singleton = &manager{state: "off"}
	})
	return singleton
}

func (m *manager) GetState() string {
	return m.state
}

func (m *manager) SetState(s string) {
	m.state = s
}
