package main

import (
	"fmt"
	"log"
)

func main() {
	// Read input data

	fmt.Printf("Enter 3-digit number: ")
	var number int
	fmt.Scanf("%d\n", &number)

	// Check the conditions

	if number < 100 || number > 999 {
		log.Fatalf("The number should have 3 digits, got %d", number)
		return
	}

	// Calculations (without cycles)

	fmt.Print("Reverted: ")

	//==== This part can be cycled ====
	fmt.Printf("%d", number%10)
	number /= 10
	fmt.Printf("%d", number%10)
	number /= 10
	fmt.Printf("%d", number%10)
	//=================================

	fmt.Println("")
}
