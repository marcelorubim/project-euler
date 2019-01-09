package main

import "testing"

func TestTask16(t *testing.T) {
	v := Task16(1000)
	if v != 26 {
		t.Error("Expected 26, got", v)
	}
}
