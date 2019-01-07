package main

import "testing"

func TestTask11(t *testing.T) {
	v := Task11()
	if v != 1788696 {
		t.Error("Expected 1788696, got", v)
	}
}
