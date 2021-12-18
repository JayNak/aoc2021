package aoc

import (
	"math"
	"regexp"
	"strconv"

	"github.com/jaynak/aoc2021/pkg/util"
)

func Day17(path string) (int, int) {

	lines := util.ReadToStrings(path)

	r := regexp.MustCompile("[-0-9]+")

	m := r.FindAllStringSubmatch(lines[0], -1)

	coords := []int{}
	for _, v := range m {
		n, err := strconv.Atoi(v[0])
		if err != nil {
			panic(err)
		}
		coords = append(coords, n)
	}

	x_vals := find_valid_x_values(coords[0], coords[1])
	initial_velo, max_height := findMaxY(coords[2], coords[3])

	count := find_valid_coords(coords[2], coords[3], coords[0], coords[1], initial_velo, x_vals)

	return max_height, count
}

func find_valid_x_values(target_min_x int, target_max_x int) map[int]bool {

	ret := make(map[int]bool)

	for i := target_max_x; i >= 0; i-- {

		sum, cnt := 0, 0
		wrk := i
		for sum <= target_max_x && wrk > 0 {
			sum += wrk
			cnt++
			wrk--
			if target_min_x <= sum && sum <= target_max_x {
				// This one works, add to the bin
				ret[i] = true
			}
		}
	}

	// fmt.Println(ret)

	return ret
}

func find_valid_coords(target_min_y int, target_max_y int, target_min_x int, target_max_x int, max_y int, valid_x map[int]bool) int {

	count := 0

	for x := range valid_x {
		for y := target_min_y; y <= max_y; y++ {

			curr_x, curr_y := 0, 0
			for i := 0; i < math.MaxInt32; i++ {

				if x-i >= 0 {
					curr_x = curr_x + x - i
				}

				curr_y = curr_y + y - i

				if target_min_y <= curr_y && target_max_y >= curr_y && target_min_x <= curr_x && target_max_x >= curr_x {
					count++
					break
				}

				if curr_x > target_max_x || curr_y < target_min_y {
					break
				}
			}

		}
	}

	return count
}

func findMaxY(minY int, maxY int) (int, int) {
	// Example -10, -5
	// for height greater than 0, gravity will force it back to y=0 again
	// so what's the biggest y that fits between the min and max?

	initial_velo := int(math.Abs(float64(minY))) - 1

	max_height := 0
	for i := initial_velo; i > 0; i-- {
		max_height += i
	}

	return initial_velo, max_height

}
