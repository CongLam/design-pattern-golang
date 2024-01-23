package builder

type BicycleBuilder struct {
	v VehicleProduct
}

func (b *BicycleBuilder) SetWheels() BuildProcess {
	b.v.Wheels = 2
	return b
}

func (b *BicycleBuilder) SetSeats() BuildProcess {
	b.v.Seats = 1
	return b
}

func (b *BicycleBuilder) SetStructure() BuildProcess {
	b.v.Structure = "Bicycle"
	return b
}

func (b *BicycleBuilder) GetVehicle() VehicleProduct {
	return b.v
}
