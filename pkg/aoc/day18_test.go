package aoc

import (
	"fmt"
	"testing"
)

func TestDay18(t *testing.T) {
	n, max := Day18("../../data/18-test.txt")

	if n != 4140 {
		t.Fatalf("expected 4140, got %v\n", n)
	}

	if max != 3993 {
		t.Fatalf("expected 3993, got %v\n", max)
	}
}

func TestDay18_2(t *testing.T) {
	n, _ := Day18("../../data/18-test2.txt")

	if n != 3488 {
		t.Fatalf("expected 3488, got %v\n", n)
	}
}

func TestDay18_3(t *testing.T) {
	n, _ := Day18("../../data/18-test3.txt")

	if n != 791 {
		t.Fatalf("expected 791, got %v\n", n)
	}
}

func TestDay18_4(t *testing.T) {
	n, _ := Day18("../../data/18-test4.txt")

	if n != 1384 {
		t.Fatalf("expected 1384, got %v\n", n)
	}
}

func TestDay18_5(t *testing.T) {
	n, _ := Day18("../../data/18-test5.txt")

	if n != 1384 {
		t.Fatalf("expected 1384, got %v\n", n)
	}
}

func TestDay18_6(t *testing.T) {
	n, _ := Day18("../../data/18-test-6.txt")

	if n != 2736 {
		t.Fail()
	}
}

func TestDay18_add_1(t *testing.T) {
	n, _ := Day18("../../data/18-test-add1.txt")

	if n != 1384 {
		t.Fatalf("expected 1384, got %v\n", n)
	}
}

func TestDay18Example(t *testing.T) {
	n, _ := Day18("../../data/18-test-2736.txt")

	if n != 2736 {
		t.Fatalf("expected 2736, got %v\n", n)
	}
}

func TestReadSnailfishNumbers(t *testing.T) {
	p := ReadSnailfishNumber("[[[0,[4,5]],[0,0]],[[[4,5],[2,6]],[9,5]]]")

	// [[0,[4,5]],[0,0]]
	// [[[4,5],[2,6]],[9,5]]

	// This is a bogus test for debugging purposes
	if p == nil {
		t.Fail()
	}
}

func TestMagnitude(t *testing.T) {
	p := ReadSnailfishNumber("[[1,2],[[3,4],5]]")
	p1 := ReadSnailfishNumber("[[[[0,7],4],[[7,8],[6,0]]],[8,1]]")
	p2 := ReadSnailfishNumber("[[[[6,6],[7,6]],[[7,7],[7,0]]],[[[7,7],[7,7]],[[7,8],[9,9]]]]")
	p3 := ReadSnailfishNumber("[[[[4,0],[5,4]],[[7,7],[6,0]]],[[8,[7,7]],[[7,9],[5,0]]]]")
	p4 := ReadSnailfishNumber("[[[[4,0],[5,4]],[[7,7],[6,0]]],[[8,[7,7]],[[7,9],[5,0]]]]")

	fmt.Println(p4.magnitude())

	m := p.magnitude()
	m1 := p1.magnitude()
	m2 := p2.magnitude()
	m3 := p3.magnitude()

	if m3 != 4140 {
		t.Fatalf("Expected 4140, got %v\n", m3)
	}

	if m != 143 {
		t.Fatalf("Expected 143, got %v\n", m)
	}

	if m1 != 1384 {
		t.Fatalf("Expected 1384, got %v\n", m1)
	}
	if m2 != 4140 {
		t.Fatalf("Expected 4140, got %v\n", m2)
	}

}

func TestReduceDebug(t *testing.T) {
	p := ReadSnailfishNumber("[[[[7,15],[9,13]],[[6,13],[7,[7,9]]]],[15,[[[6,2],[5,6]],[[7,6],[4,7]]]]]")

	p.Reduce()

	if p.magnitude() != 0 {
		t.Fail()
	}
}
