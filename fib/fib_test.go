package fib

import (
	"testing"
)

func TestFib(t *testing.T) {
	//her går testen
	input := 1
	got := Fib(input)
	want := 1
	if got != want {
		t.Errorf("Fib(%d): want %d, got %d", input, want, got)
	}
	type test struct {
		input int
		want  int
	}

	var tests = []test{
		{0, 0},
		{1, 1},
		{2, 1},
		{3, 2},
		// her kan vi bare legge på nye test cases
	}

	// tc - test case
	for _, tc := range tests {
		got := Fib(tc.input)
		if got != tc.want {
			t.Errorf("Fib(%d): want %d, got %d", tc.input, tc.want, got)
		}
	}
}
