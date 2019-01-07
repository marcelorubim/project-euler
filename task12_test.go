package main

import "testing"

func TestTask12(t *testing.T) {
	v := Task12(500)
	if v != 28 {
		t.Error("Expected 28, got", v)
	}
}
