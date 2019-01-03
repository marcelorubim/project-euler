package main

import "testing"

func TestProblem1(t *testing.T) {
	v := Problem1()
	if v != 233168 {
		t.Error("Expected 233168, got", v)
	}

}
