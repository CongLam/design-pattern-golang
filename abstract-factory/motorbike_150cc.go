package abstract_factory

type Motorbike150CC struct{}

func (m *Motorbike150CC) GetCC() string {
	return "150CC"
}

func (m *Motorbike150CC) NumWheels() int {
	return 2
}

func (m *Motorbike150CC) NumSeats() int {
	return 2
}
