package abstract_factory

type Motorbike125CC struct{}

func (m *Motorbike125CC) GetCC() string {
	return "125CC"
}

func (m *Motorbike125CC) NumWheels() int {
	return 2
}

func (m *Motorbike125CC) NumSeats() int {
	return 2
}
