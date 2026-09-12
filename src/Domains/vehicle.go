package domains

type Vehicle struct {
	ID       string
	Model    string
	Year     int
	Color    string
	Status   bool
	BuyValue float64
}

func NewVehicle(model string, year int, color string, status bool, buyValue float64) Vehicle {
	return Vehicle{
		Model:    model,
		Year:     year,
		Color:    color,
		Status:   status,
		BuyValue: buyValue,
	}
}

func (v Vehicle) GetID() string {
	return v.ID
}

func (v Vehicle) GetModel() string {
	return v.Model
}

func (v Vehicle) GetYear() int {
	return v.Year
}

func (v Vehicle) GetColor() string {
	return v.Color
}

func (v Vehicle) GetStatus() bool {
	return v.Status
}

func (v Vehicle) GetBuyValue() float64 {
	return v.BuyValue
}

func (v *Vehicle) SetID(id string) {
	v.ID = id
}