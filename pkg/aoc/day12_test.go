package aoc

import "testing"

func TestDay12(t *testing.T) {
	a, b := Day12("../../data/12-test.txt")

	c, d := Day12("../../data/12-test2.txt")

	if a != 10 {
		t.Fatalf("Expected 10, got %v\n", a)
	}

	if b != 36 {
		t.Fatalf("Expected 36, got %v\n", a)
	}

	if c != 19 {
		t.Fatalf("Expected 19, got %v\n", c)
	}

	if d != 103 {
		t.Fatalf("Expected 103, got %v\n", c)
	}
}
