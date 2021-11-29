package creditcard

import (
	"errors"
	"fmt"
)

type card struct {
	number string
}

//New returns a pointer to type card
func New(cardNumber string) (*card, error) {
	if cardNumber == "" {
		return &card{}, fmt.Errorf("%s is an invalid value", cardNumber)
	}

	//This function returns a vaue of some type that we don't need to know the name of -- bc we don't use its name.
	//Instead we call GetNumber() to get the value -- bc this method is exported. 
	return &card{
		number: cardNumber,
	}, nil
}

//SetNum sets a card number
func (c *card) SetNum(num string) error {
	if num == "" {
		return errors.New("invalid number (must not be empty)")
	}
	c.number = num
	return nil
}

//GetNum gets a number that's set on a card
func (c card) GetNum() string {
	return c.number
}