package main

import "testing"

func TestTask9(t *testing.T) {
	v := Task9()
	if v != 31875000 {
		t.Error("Expected 31875000, got", v)
	}
}
