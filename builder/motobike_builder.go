package builder

type MotorbikeBuilder struct {
	v VehicleProduct
}

func (m *MotorbikeBuilder) SetWheels() BuildProcess {
	m.v.Wheels = 2
	return m
}

func (m *MotorbikeBuilder) SetSeats() BuildProcess {
	m.v.Seats = 2
	return m
}

func (m *MotorbikeBuilder) SetStructure() BuildProcess {
	m.v.Structure = "Motorbike"
	return m
}

func (m *MotorbikeBuilder) GetVehicle() VehicleProduct {
	return m.v
}
