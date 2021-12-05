package aoc

import "testing"

func TestDay5(t *testing.T) {
	g, _ := Day5("../../data/5-test.txt")

	if g != 5 {
		t.Fail()
	}
}
