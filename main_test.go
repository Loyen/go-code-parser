package main

import "testing"

func TestPrintHelloWorld(t *testing.T) {
	expected := "Hello World!"

	source := "print \"Hello World!\""

	value, err := Run(source)

	if err != nil {
		t.Errorf("TestPrintHelloWorld returned error: %q", err)
	}

	if value != expected {
		t.Errorf("TestPrintHelloWorld returned %q, expected %q", value, expected)
	}
}

func TestPrintNoExpressionGiven(t *testing.T) {
	expectedError := "Unexpected end of file"

	source := "print"

	value, err := Run(source)

	if value != "" && err.Error() == expectedError {
		t.Errorf("TestHello returned %q, expected %q", err, expectedError)
	}
}
