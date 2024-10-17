package main

import "fmt"

type Car struct {
	Brand        string
	Model        string
	Year         int
	FuelLevel    float64
	TankCapacity float64
	Mileage      float64
}

func (c *Car) GetInfo() string {
	return fmt.Sprintf("Brand: %v;\nModel: %v;\nYear: %v;\nFuelLevel: %v;\nTankCapacity: %v;\nMileage: %v\n",
		c.Brand, c.Model, c.Year, c.FuelLevel, c.TankCapacity, c.Mileage)
}

func (c *Car) Drive(distance float64) {
	ConsumptionPerDistance := distance / 15
	if c.FuelLevel-ConsumptionPerDistance < 0 {
		fmt.Printf("НЕОБХОДИМА ДОЗАПРПАВКА!\nТоплива недостаточно на дистанцию = %v км\nУровень топлива = %v л.\n\n", distance, c.FuelLevel)
	} else {
		c.FuelLevel -= ConsumptionPerDistance
		c.Mileage += distance
		if c.FuelLevel == 0 {
			fmt.Printf("Все хорошо! топлива достаточно для дистанции = %v км\n"+
				"Но НЕОБХОДИМА ДОЗАПРАВКА, бак пуст (уровень топлива = %v)\n\n", distance, c.FuelLevel) //???\v
		} else {
			fmt.Printf("Все хорошо! топлива достаточно для дистанции = %v км\n"+
				"После проезда данной дистанции топлива останется - %v л\n"+
				"Оставшегося топлива достаточно, чтобы проехать дистанцию = %v км\n\n", distance, c.FuelLevel, c.FuelLevel*15)
		}

	}
}

func (c *Car) Refuel(amount float64) {
	if c.FuelLevel+amount < c.TankCapacity {
		c.FuelLevel += amount
	} else {
		c.FuelLevel = c.TankCapacity
	}

	fmt.Println("Автомобиль заправлен!\n")
}

func (c *Car) NeedsService() {
	if c.Year > 5 || c.Mileage > 100000.0 {
		fmt.Println("Необходимо тех. обслуживание!\n")
	} else {
		fmt.Println("Автомобиль в тех. обслуживании не нуждается\n")
	}
}

func main() {
	OpelAstra := Car{
		"Opel",
		"Astra",
		4,
		30.0,
		40.0,
		100001.0,
	}
	OpelAstra.Drive(5000)
	fmt.Println(OpelAstra.GetInfo())
	OpelAstra.Refuel(20)
	fmt.Println(OpelAstra.GetInfo())
	OpelAstra.NeedsService()
}
