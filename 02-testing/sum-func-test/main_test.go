package main

import "testing"

// constants for test case
const (
	a        = 5
	b        = 10
	expected = 15
)

func TestSum(t *testing.T) {

	sum := Sum(a, b)

	if sum != expected {
		t.Errorf("Sum was incorrect, expected: %d, got: %d.", expected, sum)
	}
}
