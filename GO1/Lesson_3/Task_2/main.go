package main

import (
	"fmt"
	"strings"
)

func IsLengthCorrect(password string) bool {
	return len(password) >= 8 && len(password) <= 15
}

func HasLowerCase(password string) bool {
	return strings.Compare(strings.ToUpper(password), password) != 0
}

func HasUpperCase(password string) bool {
	return strings.Compare(strings.ToLower(password), password) != 0
}
func HasNumber(password string) bool {
	numbers := "0123456789"
	for _, ch := range password {
		if strings.Contains(numbers, string(ch)) {
			return true
		}
	}
	return false
}
func HasSpecial(password string) bool {
	special := "\\_!@#$%^&"
	for _, ch := range password {
		if strings.Contains(special, string(ch)) {
			return true
		}
	}
	return false
}
func CheckPassword(password string) (ok bool, description string) {
	if !IsLengthCorrect(password) {
		description = "Password length mismatch"
	} else if !HasLowerCase(password) {
		description = "Password should contain lowercase"
	} else if !HasUpperCase(password) {
		description = "Password should contain uppercase"
	} else if !HasNumber(password) {
		description = "Password should contain numbers"
	} else if !HasSpecial(password) {
		description = "Password should contain special symbol"
	} else {
		ok = true
	}
	return
}
func InputPassword() (password string) {
	fmt.Printf("Enter your password: ")
	fmt.Scanf("%s\n", &password)
	return
}

func main() {
	for i := 1; i < 6; i++ {
		fmt.Printf("%d) ", i)
		password := InputPassword()
		ok, description := CheckPassword(password)
		if ok {
			fmt.Println("Thank you, password accepted")
			return
		} else {
			fmt.Println(description)
		}
	}
	fmt.Println("The end")
}
