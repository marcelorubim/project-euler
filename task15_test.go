package main

import "testing"

func TestTask15(t *testing.T) {
	v := Task15(2)
	if v != 6 {
		t.Error("Expected 6, got", v)
	}
}
