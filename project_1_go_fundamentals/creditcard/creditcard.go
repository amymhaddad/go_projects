package creditcard

import "fmt"

type card struct {
	Number string
}

func New(cardNumber string) (card, error) {
	if cardNumber == "" {
		return card{}, fmt.Errorf("%s is an invalid value", cardNumber)
	}

	return card{
		Number: cardNumber,
	}, nil
}

//take out validation in New()
