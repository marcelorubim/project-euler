package main

import "testing"

func TestTask21(t *testing.T) {
	v := SumProperDivisors(284)
	if v != 220 {
		t.Error("Expected 284, got", v)
	}

}
