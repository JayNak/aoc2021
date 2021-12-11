package aoc

import (
	"testing"

	"github.com/jaynak/aoc2021/pkg/util"
)

func TestDay11(t *testing.T) {
	a, b := Day11("../../data/11-test.txt")

	if a != 1656 {
		t.Fatalf("Expected 1656, got %v\n", a)
	}

	if b != 195 {
		t.Fatalf("Expected 195, got %v\n", b)
	}
}

func TestGridBuild(t *testing.T) {
	grid := util.ReadToIntArray("../../data/11-test2.txt")

	og := NewOctoGrid(grid)

	if len(og.grid) != 20 {
		t.Fail()
	}
}
