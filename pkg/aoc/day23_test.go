package aoc

import (
	"fmt"
	"testing"
)

func TestAStar(t *testing.T) {
	start := rooms("21001332       ")
	target := rooms("00112233       ")

	cost, _ := a_star(start, target)

	if cost != 0 {
		t.Fatalf("Expected 0, got %v\n", cost)
	}
}

func TestAStarExample(t *testing.T) {
	start := rooms("10231230       ")
	target := rooms("00112233       ")

	cost, path := a_star(start, target)

	for _, rm := range path {
		rm.print()
	}

	if cost != 12521 {
		t.Fatalf("Expected 12521, got %v\n", cost)
	}
}

func TestAStarGmailpart1(t *testing.T) {
	start := rooms("03231012       ")
	target := rooms("00112233       ")

	cost, path := a_star(start, target)

	for _, rm := range path {
		rm.print()
	}

	if cost != 12521 {
		t.Fatalf("Expected 12521, got %v\n", cost)
	}
}

func TestTargetCost(t *testing.T) {
	test1 := rooms("00112233       ")

	cost1 := test1.weight_to_complete()

	if cost1 != 0 {
		t.Fatalf("Expected 0, got %v\n", cost1)
	}

}

func TestTargetCost2(t *testing.T) {
	test2 := rooms("21001332       ")

	cost2 := test2.weight_to_complete()

	if cost2 != 6199 {
		t.Fatalf("Expected 6199, got %v\n", cost2)
	}
}

func TestTargetCost3(t *testing.T) {
	test2 := rooms("1023 230      1")

	test2.print()

	cost2 := test2.weight_to_complete()

	if cost2 != 6199 {
		t.Fatalf("Expected 6199, got %v\n", cost2)
	}
}

func TestTargetCost4(t *testing.T) {
	test2 := rooms("1023 230     1 ")
	test2.print()
	cost2 := test2.weight_to_complete()

	if cost2 != 6199 {
		t.Fatalf("Expected 6199, got %v\n", cost2)
	}
}

func TestTargetCost5(t *testing.T) {
	test2 := rooms("1023 230 1     ")
	test2.print()
	cost2 := test2.weight_to_complete()

	if cost2 != 7489 {
		t.Fatalf("Expected 7489, got %v\n", cost2)
	}
}

func TestFindValidMoves(t *testing.T) {
	test := rooms("10231230       ")

	moves := test.find_valid_moves()
	fmt.Println(test)
	for room, move := range moves {
		fmt.Printf("%v: %v\n", room, move)
	}

	if len(moves) != 27 {
		t.Fatalf("Expected 28, got %v\n", len(moves))
	}
}

func TestFindValidMoves2(t *testing.T) {
	test := rooms("   1033221    0")

	moves := test.find_valid_moves()
	fmt.Println(test)
	for room, move := range moves {
		fmt.Printf("%v: %v\n", room, move)
	}

	if len(moves) != 28 {
		t.Fatalf("Expected 28, got %v\n", len(moves))
	}
}

func TestFindValidMoves3(t *testing.T) {
	test := rooms("   322131100    0033221")

	moves := test.find_valid_moves()
	fmt.Println(test)
	for room, move := range moves {
		fmt.Printf("%v: %v\n", room, move)
	}

	if len(moves) != 28 {
		t.Fatalf("Expected 28, got %v\n", len(moves))
	}
}

func TestFindValidMoves4(t *testing.T) {
	test := rooms("0323 012 1     ")

	moves := test.find_valid_moves()
	fmt.Println(test)
	for room, move := range moves {
		fmt.Printf("%v: %v\n", room, move)
	}

	if len(moves) != 20 {
		t.Fatalf("Expected 20, got %v\n", len(moves))
	}
}
