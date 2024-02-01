package main

import (
	"fmt"
	"log"
	"math"
)

func CheckNotNegative(value int) {
	if value < 0 {
		log.Fatalf("All values must be not negative, got %d", value)
	}
}

func ReadPoint(description string) int {
	var res int
	fmt.Print(description)
	fmt.Scanf("%d\f", &res)
	return res
}

func ReadPoints() (int, int, int, int, int, int) {
	x1 := ReadPoint("Input X coordinate of point 1: ")
	CheckNotNegative(x1)
	y1 := ReadPoint("Input Y coordinate of point 1: ")
	CheckNotNegative(y1)
	x2 := ReadPoint("Input X coordinate of point 2: ")
	CheckNotNegative(x2)
	y2 := ReadPoint("Input Y coordinate of point 2: ")
	CheckNotNegative(y2)
	x3 := ReadPoint("Input X coordinate of point 3: ")
	CheckNotNegative(x3)
	y3 := ReadPoint("Input Y coordinate of point 3: ")
	CheckNotNegative(y3)

	return x1, y1, x2, y2, x3, y3
}

func GetDistance(x1, y1, x2, y2 int) float64 {
	return math.Sqrt(float64((x1-x2)*(x1-x2) + (y1-y2)*(y1-y2)))
}

func GetSides(x1, y1, x2, y2, x3, y3 int) (float64, float64, float64) {
	ab := GetDistance(x1, y1, x2, y2)
	bc := GetDistance(x2, y2, x3, y3)
	ac := GetDistance(x3, y3, x1, y1)
	return ab, bc, ac
}

func IsTriangleExists(x1, y1, x2, y2, x3, y3 int) bool {
	ab, bc, ac := GetSides(x1, y1, x2, y2, x3, y3)
	return ab < ac+bc && ac < ab+bc && bc < ac+ab
}

func GetTriangleSquare(x1, y1, x2, y2, x3, y3 int) float64 {
	// Use Heron's formula
	ab, bc, ac := GetSides(x1, y1, x2, y2, x3, y3)
	p := float64(ab+bc+ac) / 2
	square := math.Sqrt(p * (p - ab) * (p - bc) * (p - ac))
	return square
}

func IsTriangleRightAngled(x1, y1, x2, y2, x3, y3 int) bool {
	ab, bc, ac := GetSides(x1, y1, x2, y2, x3, y3)

	// We can convert squares back to integers, otherwise we cannot compare it properly, only with some threshold.
	abSquared := int(ab * ab)
	bcSquared := int(bc * bc)
	acSquared := int(ac * ac)

	return abSquared == bcSquared+acSquared || bcSquared == abSquared+acSquared || acSquared == bcSquared+abSquared
}

func main() {
	x1, y1, x2, y2, x3, y3 := ReadPoints()
	if IsTriangleExists(x1, y1, x2, y2, x3, y3) {
		fmt.Println("Triangle exists")
		fmt.Printf("Square: %.2f\n", GetTriangleSquare(x1, y1, x2, y2, x3, y3))
		if IsTriangleRightAngled(x1, y1, x2, y2, x3, y3) {
			fmt.Println("The triangle is right angled")

		} else {
			fmt.Println("The triangle isn't right angled")
		}
	} else {
		fmt.Println("Triangle doesn't exist")
	}
}
