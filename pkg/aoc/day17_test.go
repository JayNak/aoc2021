package aoc

import "testing"

func TestDay17(t *testing.T) {
	a, b := Day17("../../data/17-test.txt")

	if a != 45 {
		t.Fail()
	}

	if b != 112 {
		t.Fatalf("expected 112 got %v\n", b)
	}
}

func TestDay17Full(t *testing.T) {
	a, b := Day17("../../data/17.txt")

	if a != 8256 {
		t.Fail()
	}

	if b != 2326 {
		t.Fail()
	}
}
