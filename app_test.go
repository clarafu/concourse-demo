package main

import (
	"testing"
)

func TestCalculate(t *testing.T) {
	message, err := ConvertToMessage("1")
	if err != nil {
		t.Error("Converting to message failed")
	}

	if message != "This app is using commit number 1.\n" {
		t.Error("Did not match expected message")
	}

	message, err = ConvertToMessage("a")
	if err == nil {
		t.Error("Should not be able to convert letter to a number")
	}
}
