package interfaces

type IMotorcycle interface {
	IVehicle
	GetCategory() string
	GetEngineCapacity() int
}