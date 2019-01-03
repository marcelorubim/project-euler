package main

import "testing"

func TestProblem4(t *testing.T) {
	v := Problem4()
	if v != 906609 {
		t.Error("Expected 906609, got", v)
	}
}
