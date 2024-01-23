package abstract_factory

type SportsBicycle struct{}

func (n *SportsBicycle) GetType() string {
	return "Sport Bicycle"
}

func (n *SportsBicycle) NumSeats() int {
	return 1
}

func (*SportsBicycle) NumWheels() int {
	return 1
}
