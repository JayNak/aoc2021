package aoc

import "testing"

func TestDay22(t *testing.T) {
	a, b := Day22("../../data/22-test.txt")

	if a != 590784 {
		t.Fatalf("expected 590784, got %v\n", a)
	}

	if b != 0 {
		t.Fail()
	}
}
