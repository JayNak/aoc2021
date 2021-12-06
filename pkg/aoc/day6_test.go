package aoc

import "testing"

func TestDay6(t *testing.T) {
	a, b := Day6("../../data/6-test.txt")

	if a != 5934 {
		t.Fatalf("Expected 5934, got %v\n", a)
	}

	if b != 26984457539 {
		t.Fatalf("Expected 26984457539, got %v\n", a)
	}
}
