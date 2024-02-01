package main

import (
	"fmt"
	"log"
)

func PrintRow(len int, first, second rune) {
	for i := 0; i < len; i++ {
		var cell rune = first
		// Use the second rune if it is an odd cell
		if i%2 == 1 {
			cell = second
		}
		fmt.Printf("%c ", cell)
	}
	fmt.Println()
}

func PrintBoard(size int, first, second rune) {
	for i := 0; i < size; i++ {
		// Change the order of the first and second runes depending on the evenness of the row
		if i%2 == 0 {
			PrintRow(size, first, second)
			continue
		}
		PrintRow(size, second, first)
	}
}

func main() {
	var size int

	// Input
	fmt.Print("Enter the size of the board: ")
	_, err := fmt.Scanf("%d\n", &size)

	// Checks
	if err != nil {
		log.Fatal("Wrong input: ", err)
	}
	if size < 0 || size > 20 {
		log.Fatalf("Size should be in range [0, 20], got %d", size)
	}

	// Output
	fmt.Println("Your board:")
	PrintBoard(size, '0', '1')
}
