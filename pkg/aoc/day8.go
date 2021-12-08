package aoc

import (
	"strconv"
	"strings"

	"github.com/jaynak/aoc2021/pkg/util"
)

func Day8(path string) (int, int) {

	lines := util.ReadToStrings(path)

	count := 0

	for _, line := range lines {
		count += decodeLine(line)
	}

	return countUnique(lines), count
}

func decodeLine(line string) int {
	fields := strings.Fields(line)

	digits := make(map[uint8]int)
	revmap := make(map[int]uint8)
	fives, sixes := []uint8{}, []uint8{}
	digits_to_decode := []uint8{}
	in_payload := false

	for _, field := range fields {

		if field == "|" {
			in_payload = true
			continue
		}

		d := readDigit(field)

		if in_payload {
			digits_to_decode = append(digits_to_decode, d)
		}

		switch len(field) {
		case 2:
			// 1
			digits[d] = 1
			revmap[1] = d
		case 3:
			//7
			digits[d] = 7
			revmap[7] = d
		case 4:
			//4
			digits[d] = 4
			revmap[4] = d
		case 7:
			//8
			digits[d] = 8
			revmap[8] = d
		case 5:
			// 2, 3, 5
			fives = append(fives, d)
		case 6:
			// 0, 6, 9
			sixes = append(sixes, d)
		}
	}

	// Sort out and sixes
	n := 0
	for _, d := range sixes {
		if d&revmap[4] == revmap[4] {
			n = 9
		} else if d&revmap[1] == revmap[1] {
			n = 0
		} else {
			n = 6
		}

		digits[d] = n
		revmap[n] = d
	}

	// Sort out the fives
	for _, d := range fives {
		if d&revmap[7] == revmap[7] {
			n = 3
		} else if d&revmap[6] == d {
			n = 5
		} else {
			n = 2
		}

		digits[d] = n
		revmap[n] = d
	}

	digit_string := ""

	for _, digit := range digits_to_decode {
		digit_string = digit_string + strconv.Itoa(digits[digit])
	}

	ret, err := strconv.Atoi(digit_string)
	if err != nil {
		panic(err)
	}

	return ret
}

func readDigit(s string) uint8 {

	bitmap := map[rune]uint8{
		'a': 1,
		'b': 2,
		'c': 4,
		'd': 8,
		'e': 16,
		'f': 32,
		'g': 64,
	}

	var dig uint8
	for _, r := range s {
		dig = dig | bitmap[r]
	}

	return dig
}

func countUnique(lines []string) int {
	count := 0
	for _, line := range lines {
		f := strings.Fields(line)

		after := false
		for _, digit := range f {
			if digit == "|" {
				after = true
			}

			if after {
				// Count this one
				switch len(digit) {
				case 2, 3, 4, 7:
					count++
				}
			}
		}
	}

	return count
}
