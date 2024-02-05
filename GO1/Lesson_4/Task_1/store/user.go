package store

import (
	"errors"
	"fmt"
	"regexp"
)

type User struct {
	name    string
	phone   string
	address *Address
}

func (user *User) SetName(value string) error {
	value = trim(value)

	ok, _ := regexp.MatchString(`[a-zA-Zа-яА-яёЁ\s]+`, value)
	if !ok {
		return errors.New("wrong input data")
	}

	user.name = value
	return nil
}
func (user *User) Name() string {
	return user.name
}

func (user *User) SetPhone(value string) error {
	value = trim(value)

	if len(value) != 10 {
		return errors.New("phone number should be of length 10")
	}

	ok, _ := regexp.MatchString(`\d`, value)
	if !ok {
		return errors.New("you should use only digits")
	}

	user.phone = value
	return nil
}
func (user *User) Phone() string {
	return user.phone
}

func (user *User) SetAddress(value string) error {
	address, err := NewAddress(value, ",")
	if err != nil {
		return err
	}

	user.address = address
	return nil
}

func (user *User) Address() *Address {
	return user.address
}

func (user *User) ToString() string {
	if user == nil {
		return "<NOT SPECIFIED>"
	}
	name := user.name
	phone := user.phone
	if name == "" {
		name = "<NOT SPECIFIED>"
	}
	if phone == "" {
		phone = "<NOT SPECIFIED>"
	}
	return fmt.Sprintf("Name: %s\nPhone: %s\nAddress: \n%s", name, phone, user.address.ToString())
}
func NewUser(name string) (*User, error) {
	user := new(User)
	err := user.SetName(name)
	if err != nil {
		return nil, err
	}
	return user, nil
}
