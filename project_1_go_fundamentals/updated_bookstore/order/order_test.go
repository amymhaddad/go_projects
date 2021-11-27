package order_test

import (
	"order"
	"testing"
)

func TestNew(t *testing.T) {
	t.Parallel()

	want := "4"
	val, err := order.New(want)
	if err != nil {
		t.Fatal(err)
	}

	got := val.GetId()
	if want != got {
		t.Errorf("want: %s, got: %s", want, got)
	}
}

func TestInvalidNewId(t *testing.T) {
	t.Parallel()

	id := ""
	_, err := order.New(id)

	if err == nil {
		t.Fatal(err)
	}

}

func TestSetId(t *testing.T) {
	t.Parallel()

	want := "4"
	o, err := order.New(want)

	if err != nil {
		t.Error(err)
	}

	err = o.SetId("")
	if err == nil {
		t.Fatal(err)
	}

	err = o.SetId(want)
	if err != nil {
		t.Fatal(err)
	}

	got := o.GetId()
	if want != got {
		t.Errorf("want: %s, got: %s", want, got)
	}

}
