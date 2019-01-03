package main

import "testing"

func TestTask7(t *testing.T) {
	v := Task7(10001)
	if v != 13 {
		t.Error("Expected 13, got", v)
	}
}
