package main

import "testing"

func TestWithoutGoMod(t *testing.T) {
	word := "Hello World"
	if got := Hello(); got != word {
		t.Errorf("Hello()= %q, want %q", got, word)
	}
}
