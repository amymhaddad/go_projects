package creditcard_test

import (
	"creditcard"
	"testing"
)

//want a creditcard type to always be valid
//This means we won't let people construct their own value of the type bc they could create an invalid type
func TestNewValid(t *testing.T) {

	want := "123456789"
	//New() is a function that validates, so it'll return an error
	card, err := creditcard.New(want)

	if err != nil {
		t.Fatal(err)
	}

	//This is a method call to get a number bc the nubmer field is NOT exported
	got := card.GetNum()
	if want != got {
		t.Errorf("want: %s, got: %s", want, got)
	}
}

func TestInValid(t *testing.T) {
	//If try to create a cc with empty value, I should get some error.
	_, err := creditcard.New("")

	//In this case, I don't want a nil value. So if I get a nil value, then raise an error
	if err == nil {
		t.Fatal(err)
	}
}

func TestSetNum(t *testing.T) {
	t.Parallel()

	want := "12"
	//Need to call New() to get a new creditcard number
	c, err := creditcard.New(want)

	//First check if I have an error value once I get the new cc number
	if err != nil {
		t.Error(err)
	}

	//Take the returned cc type and set a number to it.
	//Checking invalid case first
	err = c.SetNum("")
	if err == nil {
		t.Error("want error on setting invalid number, but got nil")
	}

	want = "aa"
	//Check valid case
	err = c.SetNum(want)

 	//If I err is NOT a nil value, then there's a problem
	if err != nil {
		t.Error(err)
	}

	got := c.GetNum()
	if want != got {
		t.Errorf("want %v, got %v", want, got)
	}

}
