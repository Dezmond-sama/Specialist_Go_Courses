package store

import (
	"errors"
	"fmt"
	"strings"
)

const (
	indexLength = 6
)

type Address struct {
	postcode   string
	city       string
	street     string
	house      string
	appartment string
}

func (address *Address) checkPostcode(value string) error {
	if len(value) != indexLength {
		return errors.New("wrong length of postcode")
	}
	for _, ch := range value {
		if !strings.Contains("1234567890", string(ch)) {
			return errors.New("postcode is corrupted")
		}
	}
	return nil
}

func (address *Address) SetPostcode(value string) error {
	value = trim(value)
	err := address.checkPostcode(value)
	if err != nil {
		return err
	}
	address.postcode = value
	return nil
}
func (address *Address) Postcode() string {
	return address.postcode
}

func (address *Address) SetCity(value string) error {
	value = trim(value)
	//We can check if the city contains letters, but we don't have it in the specification

	if len(value) == 0 {
		return errors.New("city name cannot be empty")
	}
	address.city = value
	return nil
}
func (address *Address) City() string {
	return address.city
}

func (address *Address) SetStreet(value string) error {
	value = trim(value)
	if len(value) == 0 {
		return errors.New("street name cannot be empty")
	}
	address.street = value
	return nil
}
func (address *Address) Street() string {
	return address.street
}

func (address *Address) SetHouse(value string) error {
	// I decided to use string instead of int here because house number can include letters
	value = trim(value)
	if len(value) == 0 {
		return errors.New("house cannot be empty")
	}
	address.house = value
	return nil
}
func (address *Address) House() string {
	return address.house
}

func (address *Address) SetAppartment(value string) error {
	// I decided to use string instead of int here because house number can include letters
	value = trim(value)
	if len(value) == 0 {
		return errors.New("appartment cannot be empty, if you live in private house enter 1")
	}
	address.appartment = value
	return nil
}
func (address *Address) Appartment() string {
	return address.appartment
}

func NewAddress(rawData string, sep string) (*Address, error) {
	const count = 5 //count of the address fields

	fmt.Println(rawData)
	data := strings.Split(rawData, sep)

	if len(data) != count {
		return nil, errors.New("field count mismatch")
	}

	address := new(Address)

	//Slice to collect all the error messages
	var messages []string = make([]string, 0, count)
	//Setters to iterate through them
	var setters = [count]func(*Address, string) error{
		func(a *Address, value string) error { return a.SetPostcode(value) },
		func(a *Address, value string) error { return a.SetCity(value) },
		func(a *Address, value string) error { return a.SetStreet(value) },
		func(a *Address, value string) error { return a.SetHouse(value) },
		func(a *Address, value string) error { return a.SetAppartment(value) },
	}

	//Iterate and collect error messages
	for i := 0; i < count; i++ {
		err := setters[i](address, data[i])
		if err != nil {
			messages = append(messages, err.Error())
		}
	}

	//No messages => all correct
	if len(messages) == 0 {
		return address, nil
	}

	//otherwise return error
	return nil, errors.New(strings.Join(messages, "\n"))
}

func (address *Address) ToString() string {
	if address == nil {
		// We set address by NewAddress method, it can be whether correct or nil
		return "<NOT SPECIFIED>"
	}
	return fmt.Sprintf("%s, %s, %s, appt %s [%s]", address.city, address.street, address.house, address.appartment, address.postcode)
}
