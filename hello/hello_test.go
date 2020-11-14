package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Mayor")
	want := "Hello Mayor"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}