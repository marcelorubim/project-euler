package main

import "testing"

func TestTask20(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Task20(tt.args.n); got != tt.want {
				t.Errorf("Task20() = %v, want %v", got, tt.want)
			}
		})
	}
}
