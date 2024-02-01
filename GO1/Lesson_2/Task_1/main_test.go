package main

import "testing"

func TestCheckData(t *testing.T) {
	singleCheck := func(value int, correctResult bool) {
		result := CheckData(value)
		if result != correctResult {
			t.Errorf("CheckData(%d) must be %t, got %t", value, correctResult, result)
		}
	}
	singleCheck(-1, false)
	singleCheck(0, true)
	singleCheck(42, true)
	singleCheck(200, true)
	singleCheck(201, false)
	singleCheck(12412, false)
	singleCheck(-12412, false)
}

func TestGetWordEnding(t *testing.T) {
	singleCheck := func(value int, correctResult string) {
		ending := GetWordEnding(value)
		if ending != correctResult {
			t.Errorf("GetWordEnding(%d) must have ending '%s', got %s", value, correctResult, ending)
		}
	}
	singleCheck(0, "ок")
	singleCheck(1, "ка")
	singleCheck(2, "ки")
	singleCheck(5, "ок")
	singleCheck(10, "ок")
	singleCheck(14, "ок")
	singleCheck(21, "ка")
	singleCheck(93, "ки")
	singleCheck(100, "ок")
	singleCheck(121, "ка")
	singleCheck(102, "ки")
	singleCheck(115, "ок")
	singleCheck(110, "ок")
	singleCheck(114, "ок")
	singleCheck(121, "ка")
	singleCheck(142, "ки")
	singleCheck(200, "ок")
}
