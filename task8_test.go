package main

import "testing"

func TestTask8(t *testing.T) {
	v := Task8(13)
	if v != 5832 {
		t.Error("Expected 5832, got", v)
	}
}
