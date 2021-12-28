package aoc

import "testing"

func TestDay25(t *testing.T) {
	a, _ := Day25("../../data/25-test.txt")

	if a != 58 {
		t.Fatalf("expected 58, got %v\n", a)
	}
}

func TestDay25Full(t *testing.T) {
	a, _ := Day25("../../data/25.txt")

	if a != 58 {
		t.Fatalf("expected 58, got %v\n", a)
	}
}
