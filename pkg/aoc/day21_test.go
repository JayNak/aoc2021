package aoc

import "testing"

func TestSimulateGame(t *testing.T) {

	v := simulate_game(4, 8, 100, 1000)

	if v != 739785 {
		t.Fatalf("expected 739785, got %v\n", v)
	}
}

func TestDay21(t *testing.T) {
	a, b := Day21("../../data/21-test.txt")

	if a != 739785 {
		t.Fail()
	}
	if b != 341960390180808 {
		t.Fatalf("Expected 341960390180808, got %v\n", b)
	}
}
