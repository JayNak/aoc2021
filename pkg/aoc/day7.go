package aoc

import (
	"math"

	"github.com/jaynak/aoc2021/pkg/util"
)

func Day7(path string) (int, int) {
	lines := util.ReadToStrings(path)

	vals := util.StringToIntSlice(lines[0], ",")

	return CalculateFuel(vals, false), CalculateFuel(vals, true)
}

func CalculateFuel(vals []int, exp bool) int {
	min, max := math.MaxInt, math.MinInt
	for _, n := range vals {
		if n < min {
			min = n
		}
		if n > max {
			max = n
		}
	}

	minfuel := math.MaxInt
	mem := make(map[int]int)

	for i := min; i <= max; i++ {
		thisfuel := 0
		for _, v := range vals {
			if exp {
				thisfuel += FuelCost(util.IntAbs(v-i), mem)
			} else {
				thisfuel += util.IntAbs(v - i)
			}

		}

		if thisfuel < minfuel {
			minfuel = thisfuel
		}
	}

	return minfuel
}

func FuelCost(dist int, mem map[int]int) int {
	if dist == 0 {
		return 0
	}

	if _, ok := mem[dist]; ok {
		return mem[dist]
	}

	if dist == 1 {
		mem[dist] = 1
	} else {
		mem[dist] = dist + FuelCost(dist-1, mem)
	}

	return mem[dist]
}
