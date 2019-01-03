package main

import "testing"

func TestTask10(t *testing.T) {
	v := Task10(10)
	if v != 17 {
		t.Error("Expected 17, got", v)
	}
}
