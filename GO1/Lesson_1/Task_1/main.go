package main

import (
	"fmt"
	"log"
)

const (
	fuelPrice = 48
)

func main() {
	var distance float64
	var fuelPer100Km float64

	// Read input data

	fmt.Print("Enter distance (50 - 10000): ")
	fmt.Scanf("%f\n", &distance)
	fmt.Print("Enter fuel consumption (5-25): ")
	fmt.Scanf("%f\n", &fuelPer100Km)

	// Check the conditions

	if distance < 50 || distance > 10000 {
		log.Fatalf("The distance should be in between 50 and 10000, got %.2f", distance)
		return
	}
	if fuelPer100Km < 5 || fuelPer100Km > 25 {
		log.Fatalf("The fuel consumption should be in between 5 and 25, got %.2f", fuelPer100Km)
		return
	}

	// Result

	totalPrice := distance * fuelPer100Km * fuelPrice / 100

	fmt.Printf("Total price: %.2f\n", totalPrice)
}
