package aoc

import "testing"

func TestDay9(t *testing.T) {
	a, _ := Day9("../../data/9-test.txt")

	if a != 15 {
		t.Fatalf("Expected 15, got %v\n", a)
	}
}
