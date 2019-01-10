package main

import "testing"

func TestTask18(t *testing.T) {
	v := Task18("task18.data")
	if v != 23 {
		t.Error("Expected 23, got", v)
	}

}
