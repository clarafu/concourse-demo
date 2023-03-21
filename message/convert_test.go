package message_test

import (
	"demo/message"
	"testing"
)

func TestCalculate(t *testing.T) {
	convertedMes, err := message.Convert("1")
	if err != nil {
		t.Error("Converting to message failed")
	}

	if convertedMes != "This app is using commit number 1.\n" {
		t.Error("Did not match expected message")
	}

	convertedMes, err = message.Convert("a")
	if err == nil {
		t.Error("Should not be able to convert letter to a number")
	}
}
