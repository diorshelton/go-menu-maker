package main

import (
	"testing"
	// "strings"
)

func TestPrompt(t *testing.T) {

}

func TestItemPrompt(t *testing.T) {
	got := ItemPrompt()
	want := Item{Name: "Boba Tea", Price: 7.99}

	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}
