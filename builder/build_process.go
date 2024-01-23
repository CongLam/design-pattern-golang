package builder

type BuildProcess interface {
	// 3 method needed to build 1 vehicle
	SetWheels() BuildProcess
	SetSeats() BuildProcess
	SetStructure() BuildProcess

	// Get vehicle infomation
	GetVehicle() VehicleProduct
}
