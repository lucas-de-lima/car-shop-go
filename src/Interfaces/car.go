package interfaces

type ICar interface {
	IVehicle
	GetDoorsQty() int
	GetSeatsQty() int
}