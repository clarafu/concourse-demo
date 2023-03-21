package demo_test

import (
	"demo"
	"testing"
)

func TestCalculate(t *testing.T) {
	message, err := demo.ConvertToMessage("1")
	if err != nil {
		t.Error("Converting to message failed")
	}

	if message != "This app is using commit number 1.\n" {
		t.Error("Did not match expected message")
	}

	message, err = demo.ConvertToMessage("a")
	if err == nil {
		t.Error("Should not be able to convert letter to a number")
	}
}
