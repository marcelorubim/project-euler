package main

import "testing"

func TestProblem5(t *testing.T) {
	v := Problem5()
	if v != 232792560 {
		t.Error("Expected 232792560, got", v)
	}
}
