package main

import "testing"

func TestCalculateScore(t *testing.T) {
	v := CalculateScore(938, "COLIN")
	if v != 49714 {
		t.Error("Expected 49714, got", v)
	}

}
