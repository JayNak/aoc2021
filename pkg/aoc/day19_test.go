package aoc

import "testing"

func TestDay19(t *testing.T) {
	a, _ := Day19("../../data/19-test.txt")

	if a != 79 {
		t.Fatalf("expected 79, got %v\n", a)
	}
}

func TestDay19Full(t *testing.T) {
	a, _ := Day19("../../data/19.txt")

	if a != 79 {
		t.Fatalf("expected 79, got %v\n", a)
	}
}
