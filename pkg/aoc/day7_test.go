package aoc

import "testing"

func TestDay7(t *testing.T) {
	a, b := Day7("../../data/7-test.txt")

	if a != 37 {
		t.Fatalf("Expected 37, got %v\n", a)
	}

	if b != 168 {
		t.Fatalf("Expected 168, got %v\n", b)
	}
}
