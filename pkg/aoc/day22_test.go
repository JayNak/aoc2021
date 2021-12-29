package aoc

import "testing"

func TestDay22(t *testing.T) {
	_, b := Day22("../../data/22-test.txt")

	// if a != 590784 {
	// 	t.Fatalf("expected 590784, got %v\n", a)
	// }
	if b != 2758514936282235 {
		t.Fatalf("expected 2758514936282235, got %v\n", b)
	}
}

func TestDay22Full(t *testing.T) {
	a, b := Day22("../../data/22.txt")

	if a != 561032 {
		t.Fatalf("expected 561032, got %v\n", a)
	}

	if b != 0 {
		t.Fail()
	}
}
