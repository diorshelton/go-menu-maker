package main

import (
	"testing"
)

func TestCreateItem(t *testing.T) {
	t.Run("item name is an empty string", func(t *testing.T) {
		salad := Item{}
		_, err := salad.CreateItem("", 12)

		assertInvalidStringError(t, err, ErrInvalidName)
	})

	t.Run("item price 0 or less", func(t *testing.T) {
		roastBeefSandwich := Item{}
		_, err := roastBeefSandwich.CreateItem("Roast beef Sandwich", 0)

		assertInvalidPriceError(t, err, ErrInvalidPrice)
	})

}

func assertInvalidStringError(t testing.TB, got, want error) {
	t.Helper()
	if got == nil {
		t.Fatal("wanted an error but didn't get one")
	}
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func assertInvalidPriceError(t testing.TB, got, want error) {
	t.Helper()
	if got == nil {
		t.Fatal("wanted an error but didn't get one")
	}
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func TestPrintDetails(t *testing.T) {
	pizza := Item{}
	createdItem, _ := pizza.CreateItem("Cheese Pizza", 12.87)

	got := createdItem.PrintDetails()
	want := "Cheese Pizza, $12.87"

	if got != want {
		t.Errorf("got %v want %v,", got, want)
	}
}
