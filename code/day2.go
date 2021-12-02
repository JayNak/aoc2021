package main

import (
	"regexp"
	"strconv"
)

func Day2() (int, int) {
	instructions := ReadToStrings("../data/day2.txt")

	r := regexp.MustCompile(`([a-z]+) ([0-9]+)`)
	x := 0
	y := 0

	y2 := 0

	for _, s := range instructions {
		match := r.FindAllStringSubmatch(s, -1)

		dist, err := strconv.Atoi(match[0][2])
		if err != nil {
			panic(err)
		}

		switch match[0][1] {
		case "forward":
			x = x + dist
			y2 = y2 + dist*y
		case "down":
			y = y + dist
		case "up":
			y = y - dist
		}

	}

	return x * y, x * y2

}
