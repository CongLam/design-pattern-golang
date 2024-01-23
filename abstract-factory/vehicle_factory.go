package abstract_factory

import (
	"errors"
	"fmt"
)

const (
	BicycleFactoryType   = 1
	MotorbikeFactoryType = 2
)

type VehicleFactory interface {
	NewVehicle(v int) (Vehicle, error)
}

func BuildFactory(f int) (VehicleFactory, error) {
	switch f {
	case BicycleFactoryType:
		return new(BicycleFactory), nil
	case MotorbikeFactoryType:
		return new(MotorbikeFactory), nil
	default:
		err := fmt.Sprintf("Factory with id %d not recognized\n", f)
		return nil, errors.New(err)
	}
}
