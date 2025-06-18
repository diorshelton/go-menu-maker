package main

import (
	"testing"
)

func TestCreateItem(t *testing.T) {
	var pizza Item

	got, _ := pizza.CreateItem("Cheese Pizza", 12.87)
	want := "Cheese Pizza was created"

	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}
func TestPrintDetails(t *testing.T) {
	var pizza Item
	_, createdItem := pizza.CreateItem("Cheese Pizza", 12.87)

	got := createdItem.PrintDetails()
	want := "Cheese Pizza, $12.87"

	if got != want {
		t.Errorf("got %v want %v,", got, want)
	}
}
