package aoc

import (
	"math"
	"strings"

	"github.com/jaynak/aoc2021/pkg/util"
)

func Day14(path string) (int, int) {

	lines := util.ReadToStrings(path)

	polymer := lines[0]

	pairs := PolymerPairs(polymer)

	subs := make(map[string]string)

	for _, line := range lines {
		if line == polymer || line == "" {
			continue
		}

		f := strings.Fields(line)
		subs[f[0]] = f[2]
	}

	min10, max10 := 0, 0

	for i := 0; i < 40; i++ {
		pairs = RunIteration(pairs, subs)

		if i == 9 {
			min10, max10 = CountInstancesFromPairs(pairs, rune(polymer[0]))
		}
	}

	min, max := CountInstancesFromPairs(pairs, rune(polymer[0]))

	return max10 - min10, max - min
}

func RunIteration(pairs map[string]int, subs map[string]string) map[string]int {
	ret := make(map[string]int)

	for p, count := range pairs {
		if _, ok := subs[p]; !ok {
			panic("No Rule!")
		}

		left := string(p[0]) + subs[p]
		right := subs[p] + string(p[1])

		ret[left] += count
		ret[right] += count
	}

	return ret
}

func CountInstancesFromPairs(pairs map[string]int, firstchar rune) (int, int) {

	ret := make(map[rune]int)

	ret[firstchar] = 1

	for p, count := range pairs {
		ret[rune(p[1])] += count
	}

	// Need to use int64s here
	min := math.MaxInt64
	max := math.MinInt64
	for _, v := range ret {
		if v > max {
			max = v
		}
		if v < min {
			min = v
		}
	}

	// fmt.Println(ret)
	// fmt.Printf("%v: %v\n", max, min)
	return min, max
}

func PolymerPairs(source string) map[string]int {

	ret := make(map[string]int)

	for i := range source {
		if i == len(source)-1 {
			break
		}

		pair := source[i : i+2]
		ret[pair] += 1

	}

	return ret
}
