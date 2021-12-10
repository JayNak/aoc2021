package aoc

import (
	"fmt"
	"testing"
)

func TestDay10(t *testing.T) {
	a, b := Day10("../../data/10-test.txt")

	if a != 26397 {
		t.Fatalf("Expected 26397, got %v\n", a)
	}

	if b != 288957 {
		t.Fatalf("Expected 288957, got %v\n", b)
	}
}

func TestRunes(t *testing.T) {
	r := []rune{'[', ']', '{', '}', '(', ')', '<', '>'}

	for _, a := range r {
		fmt.Println(int(a))
	}
	t.Fail()
}
