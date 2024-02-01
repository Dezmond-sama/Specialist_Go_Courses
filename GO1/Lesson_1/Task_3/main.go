package main

import (
	"fmt"
)

func main() {
	//Let's name the variables in an extravagant manner, why not.

	var primo int
	var secondo int
	var terzo int

	// Read three numbers (stil no cycles)

	fmt.Printf("Enter the first number: ")
	fmt.Scanf("%d\n", &primo)

	fmt.Printf("Enter the second number: ")
	fmt.Scanf("%d\n", &secondo)

	fmt.Printf("Enter the third number: ")
	fmt.Scanf("%d\n", &terzo)

	// Calculations (sorting for three numbers)

	if primo > secondo {
		primo, secondo = secondo, primo
	}
	if secondo > terzo {
		secondo, terzo = terzo, secondo
	}
	if primo > secondo {
		primo, secondo = secondo, primo
	}

	// Result

	fmt.Printf("Numbers in ascending order: %d, %d, %d", primo, secondo, terzo)
}
