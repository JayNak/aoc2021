package aoc

import (
	"strconv"
	"strings"

	"github.com/jaynak/aoc2021/pkg/util"
)

func Day6(path string) (int, int) {

	lines := util.ReadToStrings(path)

	l := ReadFromInput(lines[0])

	count80 := 0
	count256 := 0

	for i := 0; i < 256; i++ {
		l = ProcessGeneration(l)

		if i == 79 {
			for _, v := range l {
				count80 += v
			}
		}
	}

	for _, v := range l {
		count256 += v
	}

	return count80, count256
}

type lanternfish []int

func ReadFromInput(s string) []int {
	fish := strings.Split(s, ",")

	l := make([]int, 9)

	for _, f := range fish {
		n, err := strconv.Atoi(f)
		if err != nil {
			panic(err)
		}

		l[n]++
	}

	return l
}

func ProcessGeneration(l []int) []int {
	n := l[0]
	l = l[1:]
	l = append(l, n)
	l[6] += n

	return l
}
