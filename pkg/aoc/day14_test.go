package aoc

import "testing"

func TestDay14(t *testing.T) {
	a, b := Day14("../../data/14-test.txt")

	if a != 1588 {
		t.Fatalf("expected 1588, got %v\n", a)
	}

	if b != 2188189693529 {
		t.Fatalf("expected 2188189693529 got %v\n", b)
	}
}
