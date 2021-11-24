package creditcard

import (
	"errors"
	"fmt"
)

//This package shows how the entire type is exported and how to get and set values on an unexported type 
type card struct {
	number string
}

//Return a pointer to card. Bc people will be calling methods on the result of calling New()
func New(cardNumber string) (*card, error) {
	if cardNumber == "" {
		return &card{}, fmt.Errorf("%s is an invalid value", cardNumber)
	}

	return &card{
		number: cardNumber,
	}, nil
}

//take out validation in New()

func (c *card) SetNum(num string) error {

	if num == "" {
		//Use fmt.Errorf when there could be multiple errors to return. But here we are only testing 1 case and know the only error is an empty string
		return errors.New("invalid number (must not be empty)")
	}
	c.number = num
	return nil
}

func (c card) GetNum() string {
	return c.number
}