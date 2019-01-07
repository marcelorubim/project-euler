package main

import (
	"fmt"
	"testing"
)

func TestProblem3(t *testing.T) {
	v := Problem3(600851475143)
	if v[len(v)-1] != 29 {
		t.Error("Expected 29, got", v)
	}
	fmt.Println(v)
}
