package main

import (
	"math"
	"testing"
)

const (
	threshold = 0.01
)

func TestGetDistance(t *testing.T) {
	singleCheck := func(x1, y1, x2, y2 int, correctResult float64) {
		result := GetDistance(x1, y1, x2, y2)
		if math.Abs(result-correctResult) > threshold {
			t.Errorf("GetDistance(%d, %d, %d, %d) must return %f with threshold %f, got %f", x1, y1, x2, y2, correctResult, threshold, result)
		}
	}
	singleCheck(1, 1, 1, 1, 0)
	singleCheck(10, 15, 7, 5, 10.44)
	singleCheck(4, 2, 0, 12, 10.77)
	singleCheck(1, 1, 3, 3, 2.83)
	singleCheck(1, 1, 1, 5, 4)
	singleCheck(4, 3, 3, 4, 1.41)
	singleCheck(9, 7, 4, 4, 5.83)
	singleCheck(10, 0, 1, 0, 9)
}
func TestGetSides(t *testing.T) {

	singleCheck := func(x1, y1, x2, y2, x3, y3 int, correctAB, correctBC, correctAC float64) {
		ab, bc, ac := GetSides(x1, y1, x2, y2, x3, y3)
		if math.Abs(ab-correctAB) > threshold || math.Abs(ac-correctAC) > threshold || math.Abs(bc-correctBC) > threshold {
			t.Errorf("GetSides(%d, %d, %d, %d, %d, %d) must return (%f, %f, %f) with threshold %f, got (%f, %f, %f)", x1, y1, x2, y2, x3, y3, correctAB, correctBC, correctAC, threshold, ab, bc, ac)
		}
	}
	singleCheck(1, 2, 3, 4, 5, 6, 2.83, 2.83, 5.66)
	singleCheck(10, 0, 0, 10, 10, 10, 14.14, 10, 10)
	singleCheck(5, 5, 6, 6, 3, 9, 1.41, 4.24, 4.47)
	singleCheck(10, 6, 11, 4, 12, 2, 2.24, 2.24, 4.47)
	singleCheck(250, 0, 1000, 100, 0, 0, 756.64, 1004.99, 250)
	singleCheck(5, 5, 0, 0, 10, 0, 7.07, 10, 7.07)
	singleCheck(0, 0, 0, 0, 0, 0, 0, 0, 0)
}
func TestIsTriangleExists(t *testing.T) {
	singleCheck := func(x1, y1, x2, y2, x3, y3 int, correctResult bool) {
		result := IsTriangleExists(x1, y1, x2, y2, x3, y3)
		if result != correctResult {
			t.Errorf("IsTriangleExists(%d, %d, %d, %d, %d, %d) must return %t, got %t", x1, y1, x2, y2, x3, y3, correctResult, result)
		}
	}
	singleCheck(1, 2, 3, 4, 6, 5, true)
	singleCheck(1, 2, 3, 4, 5, 6, false)
	singleCheck(10, 20, 10, 30, 0, 0, true)
	singleCheck(1, 1, 2, 2, 3, 3, false)
	singleCheck(10, 5, 15, 5, 20, 5, false)
	singleCheck(10, 15, 10, 10, 10, 5, false)
	singleCheck(10, 5, 10, 5, 20, 5, false)
	singleCheck(0, 0, 0, 0, 0, 0, false)
}

func TestGetTriangleSquare(t *testing.T) {
	singleCheck := func(x1, y1, x2, y2, x3, y3 int, correctResult float64) {
		result := GetTriangleSquare(x1, y1, x2, y2, x3, y3)
		if math.Abs(result-correctResult) > threshold {
			t.Errorf("GetTriangleSquare(%d, %d, %d, %d, %d, %d) must return %f with threshold %f, got %f", x1, y1, x2, y2, x3, y3, correctResult, threshold, result)
		}
	}
	singleCheck(1, 2, 3, 4, 6, 5, 2)
	singleCheck(10, 20, 10, 30, 0, 0, 50)
	singleCheck(0, 0, 10, 0, 0, 10, 50)
	singleCheck(10, 10, 5, 15, 0, 10, 25)
}

func TestIsTriangleRightAngled(t *testing.T) {
	singleCheck := func(x1, y1, x2, y2, x3, y3 int, correctResult bool) {
		result := IsTriangleRightAngled(x1, y1, x2, y2, x3, y3)
		if result != correctResult {
			t.Errorf("GetTriangleSquare(%d, %d, %d, %d, %d, %d) must return %t, got %t", x1, y1, x2, y2, x3, y3, correctResult, result)
		}
	}
	singleCheck(1, 2, 3, 4, 6, 5, false)
	singleCheck(10, 20, 10, 30, 0, 0, false)
	singleCheck(0, 0, 10, 0, 0, 10, true)
	singleCheck(10, 10, 5, 15, 0, 10, true)
}
