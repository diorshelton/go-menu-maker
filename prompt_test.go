package main

import (
	"strings"
	"testing"
)

func TestItemPrompt(t *testing.T) {
	input := "boba tea\n7.99\n1\nyes\n"
	reader := strings.NewReader(input)

	got := ItemPrompt(reader)
	want := MenuItem{Name: "boba tea", Price: 7.99, Category: 1}

	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}
