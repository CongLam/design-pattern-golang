package singleton

type singleton struct {
	counter int
}

var instance *singleton

func GetInstance() *singleton {
	if instance == nil {
		instance = new(singleton)
	}
	return instance
}

func (s *singleton) Increase() {
	s.counter++
}

func (s *singleton) GetCount() int {
	return s.counter
}
