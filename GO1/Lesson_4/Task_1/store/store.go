package store

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	MainMenuCount = 4
)

func ReadInt() (int, error) {
	text, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	return strconv.Atoi(strings.Trim(text, " \n\r"))
}
func ReadString() string {
	text, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	return strings.Trim(text, " \n\r")
}
func MainMenu() int {
	fmt.Println("------- MAIN MENU -------")
	fmt.Println("| 1 | User info         |")
	fmt.Println("| 2 | Add product       |")
	fmt.Println("| 3 | Print the invoice |")
	fmt.Println("| 4 | Exit              |")
	fmt.Println("-------------------------")
	fmt.Printf("Select an option [1-%d]: ", MainMenuCount)

	opt, err := ReadInt()

	if err != nil || opt < 1 || opt > MainMenuCount {
		fmt.Println("Wrong input.", err)
	}

	fmt.Println()
	return opt
}

func NewUserMenu(cart *Cart) error {
	fmt.Println("User not created.\nWrite the name of the new user: ")
	name := ReadString()

	user, err := NewUser(name)
	if err != nil {
		return err
	}
	cart.SetCustomer(user)
	return nil
}

func UserMenu(user *User) int {
	fmt.Println("----- CUSTOMER INFO -----")
	fmt.Println(user.ToString())
	fmt.Println("-------------------------")
	fmt.Println("| 1 | Set name          |")
	fmt.Println("| 2 | Set phone         |")
	fmt.Println("| 3 | Set address       |")
	fmt.Println("| 4 | To main menu      |")
	fmt.Println("-------------------------")
	fmt.Print("Select an option [1-4]: ")

	opt, err := ReadInt()

	if err != nil || opt < 1 || opt > 4 {
		fmt.Print("Wrong option.\n\n")
		return 1
	}
	fmt.Println()
	return MainMenuCount + opt
}

func SetUserName(user *User) int {
	fmt.Print("Write new name: ")
	name := ReadString()
	err := user.SetName(name)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println()
	return 1
}

func SetUserPhone(user *User) int {
	fmt.Print("Write new phone: ")
	phone := ReadString()
	err := user.SetPhone(phone)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println()
	return 1
}

func SetUserAddress(user *User) int {
	fmt.Print("Write new address in format:\npostcode, city, street, house, appartment\nwith comma (,) separation:\n")

	address := ReadString()
	err := user.SetAddress(address)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println()
	return 1
}

func AddProduct(cart *Cart) {
	fmt.Print("Write the product title: ")
	title := ReadString()

	fmt.Print("Write the product count: ")
	count, err := ReadInt()
	if err != nil {
		fmt.Println(err)
		fmt.Println()
		return
	}

	mess, err := cart.AddProduct(title, count)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(mess)
	}
	fmt.Println()
}

func PrintCart(cart *Cart) {
	fmt.Println(cart.ToString())
	fmt.Println()
}

func Run() {
	option := 0
	var cart = NewCart()
cycle:
	for {
		switch option {
		case 0:
			option = MainMenu()
		case 1:
			if cart.Customer() == nil {
				err := NewUserMenu(cart)
				if err != nil {
					fmt.Println(err)
					option = 0
				}
				fmt.Println()
			} else {
				option = UserMenu(cart.Customer())
			}
		case 2:
			AddProduct(cart)
			option = 0
		case 3:
			PrintCart(cart)
			option = 0
		case 5:
			option = SetUserName(cart.Customer())
		case 6:
			option = SetUserPhone(cart.Customer())
		case 7:
			option = SetUserAddress(cart.Customer())
		case 8:
			option = 0
		default:
			break cycle
		}
	}
}
