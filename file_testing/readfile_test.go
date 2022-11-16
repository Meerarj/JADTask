package main

import (
	"testing"
)

// test function
func TestReturn(t *testing.T) {
	actualString := fileRead()
	expectedString := "Five"
	if actualString != expectedString {
		t.Errorf("Expected String(%s) is not same as"+
			" actual string (%s)", expectedString, actualString)
	}
}
