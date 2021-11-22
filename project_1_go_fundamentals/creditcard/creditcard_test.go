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

	got := card.Number
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

	num := "1"
	c, _ := creditcard.New(num)

	err := c.SetNum("")

	if err != nil {
		t.Fatal(err)
	}

	err = c.SetNum("22")

	if err == nil {
		t.Fatal(err)
	}

	want := creditcard.New(num)
	got, _ := want.GetNum()

	if want != got {
		t.Errorf("want %v, got %v", want, got)
	}

}
