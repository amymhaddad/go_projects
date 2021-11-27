package creditcard_test

import (
	"creditcard"
	"testing"
)

func TestNewValid(t *testing.T) {

	want := "123456789"
	card, err := creditcard.New(want)

	if err != nil {
		t.Fatal(err)
	}

	got := card.GetNum()
	if want != got {
		t.Errorf("want: %s, got: %s", want, got)
	}
}

func TestInValid(t *testing.T) {
	_, err := creditcard.New("")

	if err == nil {
		t.Fatal(err)
	}
}

func TestSetNum(t *testing.T) {
	t.Parallel()

	want := "12"
	c, err := creditcard.New(want)

	//Check if I have a valid value after created a new card type 
	if err != nil {
		t.Error(err)
	} 

	//Checking invalid case 
	err = c.SetNum("")
	//If I don't get an error w/this input, raise an error 
	if err == nil {
		t.Error("want error on setting invalid number, but got nil")
	}

	want = "aa"
	//Check valid case
	err = c.SetNum(want)

 	//If I get an error with a vaid case, raise error 
	if err != nil {
		t.Error(err)
	}

	got := c.GetNum()
	if want != got {
		t.Errorf("want %v, got %v", want, got)
	}

}
