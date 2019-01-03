package main

import "testing"

func TestProblem3(t *testing.T) {
	v := Problem3(13195)
	if v[len(v)-1] != 29 {
		t.Error("Expected 29, got", v)
	}
}
