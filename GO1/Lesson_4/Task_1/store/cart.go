package store

import (
	"errors"
	"fmt"
	"strings"
)

// # Задача

// ## Сформировать данные для отправки заказа из магазина по накладной и вывести на экран:

// 1. Наименование товара (минимум 1, максимум 100)
// 2. Количество (только числа)
// 3. ФИО покупателя (только буквы)
// 4. Контактный телефон (10 цифр)
// 5. Адрес(индекс(ровно 6 цифр), город, улица, дом, квартира)

// Эти данные не могут быть пустыми.
// Проверить правильность заполнения полей.

// реализовать несколько методов у типа "Накладная"

// createReader == NewReader

type Cart struct {
	customer *User
	products map[string]int
}

func (cart *Cart) SetCustomer(customer *User) {
	cart.customer = customer
}

func (cart *Cart) Customer() *User {
	return cart.customer
}

func (cart *Cart) AddProduct(title string, count int) (string, error) {
	//Check product title
	if len(title) < 1 || len(title) > 100 {
		return "error", errors.New("product title length should be from 1 to 100")
	}

	if count == 0 {
		return "Nothing to do", nil
	}

	//Get new count (if we already have this product)
	prev, ok := cart.products[title]
	count = prev + count
	if count > 0 {
		cart.products[title] = count
		if ok {
			return "Product count updated", nil
		} else {
			return "Product added", nil
		}
	} else if ok {
		//Remove record if new value is zero (or less) and we have this key
		delete(cart.products, title)
		return "Product deleted", nil
	}
	return "Product doesn't exist", nil
}

func (cart *Cart) ToString() string {
	var products []string = make([]string, 0, len(cart.products))
	number := 1
	for k, v := range cart.products {
		products = append(products, fmt.Sprintf("| %02d | %20s | %5d |", number, k, v))
		number++
	}
	return fmt.Sprintf(
		"-------------- CUSTOMER -------------\n"+
			"%s\n"+
			"-------------------------------------\n"+
			"%s\n"+
			"-------------------------------------\n",
		cart.Customer().ToString(),
		strings.Join(products, "\n"),
	)
}

func NewCart() *Cart {
	cart := new(Cart)
	cart.products = make(map[string]int)
	return cart
}
