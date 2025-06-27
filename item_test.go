package main

import (
	"testing"
)

func TestCreateItem(t *testing.T) {
	assertValidPrice := func(t testing.TB, got error, want error) {
		t.Helper()

		if got == nil {
			t.Fatal("wanted and error but didn't get one")
		}

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	}

	t.Run("create item struct", func(t *testing.T) {
		pizza := Item{}

		got, _ := pizza.CreateItem("Cheese Pizza", 12.87)
		want := Item{"Cheese Pizza", 12.87}

		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	})

	t.Run("invalid price", func(t *testing.T) {
		roastBeefSandwich := Item{}
		_, err := roastBeefSandwich.CreateItem("Roast beef Sandwich", 0)

		assertValidPrice(t, err, ErrInvalidPrice)
	})
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
