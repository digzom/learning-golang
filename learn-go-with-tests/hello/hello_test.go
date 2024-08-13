package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("saying hello", func(t *testing.T) {
		got := Hello("Dickson", "")
		want := "Hello, Dickson!"

		assert(t, got, want)
	})

	t.Run("says 'Hello, World!' when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "English")
		want := "Hello, World!"

		assert(t, got, want)
	})

	t.Run("says hello in portuguese", func(t *testing.T) {
		got := Hello("Carlos", "Portuguese")
		want := "Olá, Carlos!"

		assert(t, got, want)
	})

	t.Run("says hello in spanish", func(t *testing.T) {
		got := Hello("Carlos", "Spanish")
		want := "Ola, Carlos!"

		assert(t, got, want)
	})

	t.Run("says hello in french", func(t *testing.T) {
		got := Hello("Antonie", "French")
		want := "Bonjour, Antonie!"

		assert(t, got, want)
	})
}

func assert(t testing.TB, got, want string) {
	// specifying this, in case of failing, the line number of the test will be
	// in the function call instead of here
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
