package aoc

import "testing"

func TestDay20(t *testing.T) {

	a, _ := Day20("../../data/20-test.txt")

	if a != 35 {
		t.Fatalf("expected 35, got %v\n", a)
	}
}
