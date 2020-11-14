package main

import "testing"

func TestHello(t *testing.T) {

	assertCorrectMessage := func(t *testing.T, got, want string) {
		t.Helper() // By doing this when it fails the line number reported will be in our function call rather than inside our test helper. 
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Mayor", "")
		want := "Hello Mayor"
		assertCorrectMessage(t, got, want)
	})

	t.Run("say 'hello world' when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello world"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in Spanish", func(t *testing.T) {
        got := Hello("Elodie", "Spanish")
        want := "Hola Elodie"
        assertCorrectMessage(t, got, want)
	})
	
	t.Run("in french", func (t *testing.T) {
		got := Hello("Paris", "French")
		want := "Bonjour Paris"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in yoruba", func (t *testing.T) {
		got := Hello("Abike", "Yoruba")
		want := "Bawoni Abike"
		assertCorrectMessage(t, got, want)
	})
}