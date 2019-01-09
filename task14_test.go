package main

import "testing"

func TestTask14(t *testing.T) {
	v := Task14(13)
	if len(v) != 10 {
		t.Error("Expected 10, got", len(v))
	}
}
