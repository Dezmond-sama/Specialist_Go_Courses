package main

import (
	"fmt"
	"log"
)

func main() {
	// Read input data

	var number int
	fmt.Printf("Enter 4-digit number: ")
	fmt.Scanf("%d\n", &number)

	// Check the conditions

	if number < 1000 || number > 9999 {
		log.Fatalf("The number should have 4 digits, got %d", number)
		return
	}

	// Reversing

	var revertedNumber int

	revertedNumber += number % 10 // 4th digit
	revertedNumber *= 10
	revertedNumber += (number / 10) % 10 // 3th digit
	revertedNumber *= 10
	revertedNumber += (number / 100) % 10 // 2th digit
	revertedNumber *= 10
	revertedNumber += number / 1000 // 1th digit

	// Output

	if revertedNumber == number {
		fmt.Println("palindrome")
	} else {
		fmt.Println("not palindrome")
	}
}
