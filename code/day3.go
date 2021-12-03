package main

import "strconv"

func Day3() (int, int) {

	lines := ReadToStrings("../data/day3.txt")

	numLines := len(lines)
	digits := len(lines[0])

	counters := make([]int, digits)

	for _, line := range lines {
		for i, r := range line {
			if r == '1' {
				counters[i]++
			}
		}
	}

	gamma := ""
	epsilon := ""

	for _, c := range counters {
		if c > numLines/2 {
			gamma = gamma + "1"
			epsilon = epsilon + "0"
		} else {
			gamma = gamma + "0"
			epsilon = epsilon + "1"
		}
	}

	return ToInt(gamma) * ToInt(epsilon), ToInt(FindOxygen(lines)) * ToInt(FindCO2(lines))
}

func ToInt(bits string) int {
	i, err := strconv.ParseInt(bits, 2, 64)
	if err != nil {
		panic(err)
	}
	return int(i)
}

func FindOxygen(lines []string) string {
	working := lines
	digit := 0

	for len(working) > 1 {
		working = Split(working, digit, true)
		digit++
	}

	return working[0]
}

func FindCO2(lines []string) string {
	working := lines
	digit := 0

	for len(working) > 1 {
		working = Split(working, digit, false)
		digit++
	}

	return working[0]
}

func Split(lines []string, index int, more bool) []string {

	var zeroes, ones []string

	for _, line := range lines {
		if line[index] == '1' {
			ones = append(ones, line)
		} else {
			zeroes = append(zeroes, line)
		}
	}

	if more {
		if len(ones) >= len(zeroes) {
			return ones
		} else {
			return zeroes
		}
	} else {
		if len(zeroes) <= len(ones) {
			return zeroes
		} else {
			return ones
		}
	}
}
