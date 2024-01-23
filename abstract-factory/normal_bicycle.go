package abstract_factory

type NormalBicycle struct{}

func (n *NormalBicycle) GetType() string {
	return "Normal Bicycle"
}

func (n *NormalBicycle) NumSeats() int {
	return 1
}

func (*NormalBicycle) NumWheels() int {
	return 2
}
