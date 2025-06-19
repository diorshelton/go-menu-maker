package main

import (
	"testing"
)

func TestCreateItem(t *testing.T) {
	var pizza Item

	got := pizza.CreateItem("Cheese Pizza", 12.87)
	want := Item{"Cheese Pizza", 12.87}

	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}
func TestPrintDetails(t *testing.T) {
	var pizza Item
	createdItem := pizza.CreateItem("Cheese Pizza", 12.87)

	got := createdItem.PrintDetails()
	want := "Cheese Pizza, $12.87"

	if got != want {
		t.Errorf("got %v want %v,", got, want)
	}
}
