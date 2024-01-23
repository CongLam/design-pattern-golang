package abstract_factory

import (
	"errors"
	"fmt"
)

const (
	Motorbike125CCType = 1
	Motorbike150CCType = 2
)

type MotorbikeFactory struct{}

func (m *MotorbikeFactory) NewVehicle(v int) (Vehicle, error) {
	switch v {
	case Motorbike125CCType:
		return new(NormalBicycle), nil
	case Motorbike150CCType:
		return new(SportsBicycle), nil
	default:
		err := fmt.Sprintf("Vehicle of type %d not recognized \n", v)
		return nil, errors.New(err)
	}
}
