package main

import "testing"

func TestTask19(t *testing.T) {
	v := Task19()
	if v != 23 {
		t.Error("Expected 23, got", v)
	}

}
