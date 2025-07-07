package main

import (
	"testing"
	"github.com/google/uuid"
)

func TestCreateMenuItem(t *testing.T) {
	t.Run("item name is an empty string", func(t *testing.T) {
		_, err := CreateMenuItem("", 12, Entree)

		assertInvalidStringError(t, err, ErrInvalidName)
	})

	t.Run("item price 0 or less", func(t *testing.T) {
		_, err := CreateMenuItem("roast beef sandwich", 0, Entree)

		assertInvalidPriceError(t, err, ErrInvalidPrice)
	})
	t.Run("item has a category", func(t *testing.T) {
		rb, err := CreateMenuItem("roast beef sandwich", 12.50, Entree)
		if err != nil {
			t.Fatalf("err %s", err)
		}
		got := rb.Category
		want := Entree

		if got != want {
			t.Errorf("want %v but got %v", want, got)
		}
	})
	t.Run("item has an id", func(t *testing.T) {
		spaghetti, err := CreateMenuItem("spaghetti", 24.95, Entree)

		if err != nil {
			t.Errorf("Item creation err:%v", err)
		}

		got := spaghetti.ID
		if got == uuid.Nil {
			t.Errorf("got %v a nil value for id",got)
		}
	})
}

func TestEditMenuItem(t *testing.T) {
	t.Run("Can edit menu item", func(t *testing.T) {

		capreseSalad, _ := CreateMenuItem("Caprese Salad", 15.94, Entree)
		EditMenuItem(&capreseSalad, "Chicken Salad", 12.99)

		got := capreseSalad
		want := MenuItem{"Chicken Salad", 12.99, Entree, capreseSalad.ID}

		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
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

/*
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
*/

func TestPrintDetails(t *testing.T) {
	pizza, _ := CreateMenuItem("Cheese Pizza", 12.87, Entree)

	got := pizza.PrintDetails()
	want := "Cheese Pizza, $12.87"

	if got != want {
		t.Errorf("got %v want %v,", got, want)
	}
}
