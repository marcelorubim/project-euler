package main

import "testing"

func TestTask6(t *testing.T) {
	v := Task6(100)
	if v != 25164150 {
		t.Error("Expected 25164150, got", v)
	}
}
