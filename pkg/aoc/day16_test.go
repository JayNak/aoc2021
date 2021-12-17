package aoc

import "testing"

func TestDay16(t *testing.T) {
	c, _ := Day16("../../data/16-test.txt")

	if c != 16 {
		t.Fatalf("expected 16, got %v\n", c)
	}
}

func TestDay16alt(t *testing.T) {
	_, b := Day16("../../data/16-kevin.txt")

	if b != 10626195124371 {
		t.Fatalf("expected 10626195124371 got %v\n", b)
	}
}

//10626195124371
//10631807338486
//10631806242976
//10631806244017

func TestDay16Full(t *testing.T) {
	c, _ := Day16("../../data/16.txt")

	if c != 16 {
		t.Fail()
	}
}

func TestPackets(t *testing.T) {
	dec := PacketToBinaryString("38006F45291200")

	if dec != "00111000000000000110111101000101001010010001001000000000" {
		t.Fail()
	}

	dec2 := PacketToBinaryString("EE00D40C823060")

	if dec2 != "11101110000000001101010000001100100000100011000001100000" {
		t.Fail()
	}
}

func TestPacketRead(t *testing.T) {
	p := ReadPackets("110100101111111000101000", 1)

	if p[0].typ != 4 || p[0].val != 2021 {
		t.Fail()
	}
}

func TestPacketRead2(t *testing.T) {
	p := ReadPackets("00111000000000000110111101000101001010010001001000000000", 1)

	if len(p) != 3 {
		t.Fail()
	}
}
