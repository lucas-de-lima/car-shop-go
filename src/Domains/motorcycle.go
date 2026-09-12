package domains

type Motorcycle struct {
	Vehicle
	Category       string
	EngineCapacity int
}

func NewMotorcycle(model string, year int, color string, status bool, buyValue float64, category string, engineCapacity int) Motorcycle {
	return Motorcycle{
		Vehicle:        NewVehicle(model, year, color, status, buyValue),
		Category:       category,
		EngineCapacity: engineCapacity,
	}
}

func (m *Motorcycle) GetCategory() string {
	return m.Category
}

func (m *Motorcycle) GetEngineCapacity() int {
	return m.EngineCapacity
}