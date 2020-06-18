package main

import "testing"

func TestHello(t *testing.T) {
	expected := "Hello World!"

	source := "print \"Hello World!\""

	value, err := Run(source)

	if err != nil {
		t.Errorf("TestHello returned error: %q", err)
	}

	if value != expected {
		t.Errorf("TestHello returned %q, expected %q", value, expected)
	}
}
