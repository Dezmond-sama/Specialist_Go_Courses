package main

import (
	"errors"
	"fmt"
	"log"
)

func DecryptSymbol(ch string) (result rune, err error) {
	// Use map here. We can also use cycle and cast int to rune to make it shorter
	decryptionMap := map[string]rune{
		"00": 'a', "01": 'b', "02": 'c', "03": 'd', "04": 'e', "05": 'f',
		"06": 'g', "07": 'h', "08": 'i', "09": 'j', "10": 'k', "11": 'l',
		"12": 'm', "13": 'n', "14": 'o', "15": 'p', "16": 'q', "17": 'r',
		"18": 's', "19": 't', "20": 'u', "21": 'v', "22": 'w', "23": 'x',
		"24": 'y', "25": 'z', "26": ' ',
	}
	result = decryptionMap[ch]
	if result == 0 {
		err = errors.New("rune not found")
	}
	return
}

func DecryptString(code string) (result string, err error) {
	// We can add a separate function to check the code string, but we already have an error check in the DecryptSymbol function. So just check the string length.
	if len(code)%2 == 0 {
		for i := 0; i < len(code); i += 2 {
			ch, e := DecryptSymbol(code[i : i+2])
			if e == nil {
				result += string(ch)
			} else {
				err = e
				break
			}
		}
	} else {
		err = errors.New("code string is corrupted")
	}
	return
}

func InputEncryptedString() (code string) {
	fmt.Printf("Enter code string: ")
	fmt.Scanf("%s\n", &code)
	return
}

func main() {
	//Input

	code := InputEncryptedString()

	//Decrypting

	decrypted, err := DecryptString(code)
	if err != nil {
		log.Fatal("Decryption failed: ", err)
	}

	//Output

	fmt.Printf("Decrypted string: %s", decrypted)
}
