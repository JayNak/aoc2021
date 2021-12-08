package aoc

import "testing"

func TestDay8(t *testing.T) {
	a, b := Day8("../../data/8-test.txt")

	if a != 26 {
		t.Fatalf("Expected 26, got %v\n", a)
	}

	if b != 61229 {
		t.Fatalf("Expected 61229, got %v\n", b)
	}
}

func TestDigitReader(t *testing.T) {
	d := readDigit("fdgacbe")

	if d != 127 {
		t.Fatalf("Expected 127, got %v\n", d)
	}
}

func TestDecodeLine(t *testing.T) {
	d := decodeLine("acedgfb cdfbe gcdfa fbcad dab cefabd cdfgeb eafb cagedb ab | cdfeb fcadb cdfeb cdbaf")

	if d != 5353 {
		t.Fail()
	}
}
