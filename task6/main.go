package main

import (
	"fmt"
	"math/rand"
)

type Producter interface {
	GetPrice() float64
}

// PhysicalProduct ///////////////////////////////////////////////////
type PhysicalProduct struct {
	Name         string
	Price        float64
	Weight       float64
	ShippingCost float64
}

func (pp *PhysicalProduct) ApplyDiscount(discount float64) {
	pp.Price = pp.Price * (1 - discount/100)
}

func (pp *PhysicalProduct) CalculateShipping() {
	if pp.Price > 100 {
		pp.ShippingCost = 0
		return
	}

	pp.ShippingCost = 10 * pp.Weight

}

func (pp *PhysicalProduct) GetPrice() float64 {
	pp.ApplyDiscount(rand.Float64() * 10)
	return pp.Price + pp.ShippingCost
}

// DigitalProduct /////////////////////////////////////////////////////////////////
type DigitalProduct struct {
	Name       string
	Price      float64
	FileSize   float64
	LicenseKey string
}

func (dp *DigitalProduct) ApplyDiscount(discount float64) {
	dp.Price = dp.Price * (1 - discount/100)
}

func (dp *DigitalProduct) GenerateLicense() {
	key := make([]rune, 10, 10)

	for i := 0; i < 10; i++ {
		key[i] = rune(rand.Int31n(100))
	}

	dp.LicenseKey = string(key)
}

func (dp *DigitalProduct) GetPrice() float64 {
	dp.ApplyDiscount(rand.Float64() * 10)
	return dp.Price
}

///////////////////////////////////////////////

func CalculateTotalCost(products []Producter) float64 {
	var total float64
	for _, product := range products {
		total += product.GetPrice()
	}

	return total
}

func main() {
	dp := DigitalProduct{
		Name:       "AAAAAAAAA",
		Price:      1000,
		FileSize:   0,
		LicenseKey: "",
	}

	dp.GenerateLicense()

	fmt.Println(dp)
}
