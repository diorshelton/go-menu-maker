package main

import (
	"testing"
)

func TestPrompt(t *testing.T) {
	questions := []string{"What is the item name"}
	got := Prompt(questions)
	want := "What is the item name"

	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestItemPrompt(t *testing.T) {
	got := ItemPrompt()
	want := Item{Name: "Boba Tea", Price: 7.99}

	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}
