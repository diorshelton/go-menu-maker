package main

import (
	"strings"
	"testing"
)

func TestItemPrompt(t *testing.T) {
	input := "boba tea\n7.99\nyes\n"
	reader := strings.NewReader(input)

	got := ItemPrompt(reader)
	want := Item{Name: "boba tea", Price: 7.99}

	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}
