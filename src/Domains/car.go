package domains

type Car struct {
	Vehicle
	DoorsQty int
	SeatsQty int
}

func NewCar(model string, year int, color string, status bool, buyValue float64, doorsQty int, seatsQty int) Car {
	return Car{
		Vehicle:  NewVehicle(model, year, color, status, buyValue),
		DoorsQty: doorsQty,
		SeatsQty: seatsQty,
	}
}

func (c *Car) GetDoorsQty() int {
	return c.DoorsQty
}

func (c *Car) GetSeatsQty() int {
	return c.SeatsQty
}