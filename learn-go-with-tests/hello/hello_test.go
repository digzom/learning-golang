package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Dickson")
	want := "Hello, Dickson!"

	if got != want {
		// %q wraps the value in double quotes
		t.Errorf("got %q want %q", got, want)
	}
}
