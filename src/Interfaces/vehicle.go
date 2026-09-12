package interfaces

type IVehicle interface {
	GetID() string
	GetModel() string
	GetYear() int
	GetColor() string
	GetStatus() bool
	GetBuyValue() float64
}