package main

import "testing"

func TestProblem2(t *testing.T) {
	v := Problem2(4000000)
	if v != 4613732 {
		t.Error("Expected 233168, got", v)
	}

}
