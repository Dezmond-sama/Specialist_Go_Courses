package main

import "testing"

func TestCheckData(t *testing.T) {
	singleCheck := func(value string, correctResult rune, isErr bool) {
		result, err := DecryptSymbol(value)
		if (err == nil) == isErr {
			if err != nil {
				t.Errorf("DecryptSymbol(%s) return unexpected error", value)
			} else {
				t.Errorf("DecryptSymbol(%s) must return error", value)
			}
		} else if err == nil && result != correctResult {
			t.Errorf("DecryptSymbol(%s) must be %c, got %c", value, correctResult, result)
		}
	}
	singleCheck("00", 'a', false)
	singleCheck("05", 'f', false)
	singleCheck("25", 'z', false)
	singleCheck("26", ' ', false)
	singleCheck("27", 0, true)
	singleCheck("42", 0, true)
	singleCheck("asd", 0, true)
	singleCheck("0", 0, true)
	singleCheck("6", 0, true)
}

func TestDecryptString(t *testing.T) {
	singleCheck := func(value string, correctResult string, isErr bool) {
		result, err := DecryptString(value)
		if (err == nil) == isErr {
			if err != nil {
				t.Errorf("DecryptSymbol(%s) return unexpected error", value)
			} else {
				t.Errorf("DecryptSymbol(%s) must return error", value)
			}
		} else if err == nil && result != correctResult {
			t.Errorf("DecryptSymbol(%s) must be \"%s\", got \"%s\"", value, correctResult, result)
		}
	}
	singleCheck("00", "a", false)
	singleCheck("0704111114", "hello", false)
	singleCheck("", "", false)
	singleCheck("070411111427", "", true)
	singleCheck("07041111142", "", true)
}
