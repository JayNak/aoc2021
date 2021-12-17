package aoc

import "testing"

func TestDay15(t *testing.T) {
	a, b := Day15("../../data/15-test.txt")

	if a != 40 {
		t.Fatalf("expected 40, got %v\n", a)
	}

	if b != 315 {
		t.Fatalf("expected 315, got %v\n", b)
	}
}
