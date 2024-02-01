package main

import (
	"fmt"
	"log"
)

func CheckData(count int) bool {
	return count >= 0 && count <= 200
}

func GetWordEnding(count int) string {
	switch {
	case count%100 > 10 && count%100 < 20: // [ x11, x12, x13 ... x19 ]
		return "ок"
	case count%10 == 1: // ends to 1, exclude previous case
		return "ка"
	case count%10 > 1 && count%10 < 5: //ends to 2-4, exclude the first case
		return "ки"
	default:
		return "ок"
	}
}

func main() {
	var bottles int

	// Input data
	fmt.Print("Enter the number of bottles: ")
	_, err := fmt.Scanf("%d\n", &bottles)
	if err != nil {
		log.Fatal("Wrong input data: ", err)
	}

	//Check the range

	if !CheckData(bottles) {
		log.Fatal("The bottles number is out of range [0-200]: ", bottles)
	}

	//Output

	fmt.Printf("%d бутыл%s", bottles, GetWordEnding(bottles))
}
