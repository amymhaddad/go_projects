package main

import (
	"creditcard"
	"fmt"
	"log"
)

func main() {
	num := "a"
	cc, err := creditcard.New(num)

	if err != nil {
		//log.Fatal() - only call w/in main() bc it'll terminate the program
		log.Fatal(err)
	}

	//Use dot notation bc cc is a struct type. Number field is exported, so I can access it.
	//This could be a problem bc someone could do this: card.Number = "" This means I don't have an always valid function
	fmt.Println(cc.Number)
}

/*
Even though card is unexported in the creditcard package, I can access it bc New() is an exported function that accesses this struct type.
This is why creditcard.New(num) works -- New() is exported and gives me access to card, even tho card is unexported
However, I couldn't use: creditcard.card bc the name "card" is unexported

If the card had a capital C:
type Card struct {
	Number string
}

Then I could use the NAME Card outside of the package


*/
