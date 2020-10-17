package main

import "testing"

func TestHello(t *testing.T) {
	assertCorrectMessage := func(t *testing.T, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Chris", "EN")
		want := "Hello, Chris"
		assertCorrectMessage(t, got, want)
	})

	t.Run("empty string defaults to 'World'", func(t *testing.T) {
		got := Hello("", "EN")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})

	t.Run("unrecognised language defaults to English", func(t *testing.T) {
		got := Hello("Bob", "GY")
		want := "Hello, Bob"
		assertCorrectMessage(t, got, want)
	})

	t.Run("can handle FR for French", func(t *testing.T) {
		got := Hello("Bob", "FR")
		want := "Bonjour, Bob"
		assertCorrectMessage(t, got, want)
	})

}
