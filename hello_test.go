package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Patrick")
	want := "Hello, Patrick"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
