package main

import "testing"

func TestGreet(t *testing.T) {
	want := "Hello world"

	got, err := greet(language("en"))
	if err != nil {
		t.Errorf("Failed to greet with err: %q", err)
	}

	if want != got {
		t.Errorf("expected %q, got: %q", want, got)
	}
}
