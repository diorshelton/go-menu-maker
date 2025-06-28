package main

import (
	"testing"
)

func TestCreateItem(t *testing.T) {
	t.Run("item name is an empty string", func(t *testing.T) {
		_, err := CreateMenuItem("", 12)

		assertInvalidStringError(t, err, ErrInvalidName)
	})

	t.Run("item price 0 or less", func(t *testing.T) {
		_, err := CreateMenuItem("Roast beef Sandwich", 0)

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

func TestSaveItem(t *testing.T) {
	t.Run("save item", func(t *testing.T) {

		tiramisu := MenuItem{"Tiramisu", 15.99}

		got := tiramisu.SaveItem()
		want := "Tiramisu saved"

		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	})
}

func TestPrintDetails(t *testing.T) {
	pizza, _ := CreateMenuItem("Cheese Pizza", 12.87)

	got := pizza.PrintDetails()
	want := "Cheese Pizza, $12.87"

	if got != want {
		t.Errorf("got %v want %v,", got, want)
	}
}
