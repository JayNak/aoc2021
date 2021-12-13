package aoc

import "testing"

func TestDay13(t *testing.T) {
	a, _ := Day13("../../data/13-test.txt")

	if a != 17 {
		t.Fail()
	}
}

func TestDay13Full(t *testing.T) {
	a, _ := Day13("../../data/13.txt")

	if a != 17 {
		t.Fail()
	}
}
