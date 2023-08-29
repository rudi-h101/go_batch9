package main

import "fmt"

type Car struct {
	Type   string
	FuelIn float64 
}

func (c Car) CalculateEstimatedDistance() float64 {
	fuelEfficiency := 1.5 
	estimatedDistance := c.FuelIn * fuelEfficiency
	return estimatedDistance
}

func main() {
	var TypeCar string
	var FuelInCar float64
	
	fmt.Printf("Tipe mobil: ")
	fmt.Scan(&TypeCar)

	fmt.Printf("Kapasitas bahan bakar mobil: ")
	fmt.Scan(&FuelInCar)
	
	car := Car{
		Type:   TypeCar,
		FuelIn: FuelInCar, 
	}

	estimatedDistance := car.CalculateEstimatedDistance()
	fmt.Printf("Mobil tipe %s dengan kapasitas bahan bakar %.2f liter dapat menempuh perkiraan jarak %.2f mill.\n", car.Type, car.FuelIn, estimatedDistance)
}
