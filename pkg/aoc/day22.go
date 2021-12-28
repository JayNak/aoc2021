package aoc

import (
	"regexp"
	"strconv"

	"github.com/jaynak/aoc2021/pkg/util"
)

func Day22(path string) (int, int) {
	lines := util.ReadToStrings(path)

	boxes := init_grid(50)

	instructions := parse_instructions(lines)

	boxes = apply_instructions(boxes, instructions)

	count := 0
	for _, x := range boxes {
		for _, y := range x {
			for _, val := range y {
				if val {
					count++
				}
			}
		}
	}

	return count, 0
}

type inst struct {
	onoff bool
	x     []int
	y     []int
	z     []int
}

func apply_instructions(boxes [][][]bool, instructions []inst) [][][]bool {

	max := len(boxes) / 2

	for _, in := range instructions {

		// TODO: Catch non-overlapping cubes

		// Brute force
		for x := in.x[0]; x <= in.x[1]; x++ {
			if x < -max || x > max {
				continue
			}
			for y := in.y[0]; y <= in.y[1]; y++ {
				if y < -max || y > max {
					continue
				}
				for z := in.z[0]; y <= in.z[1]; y++ {
					if z < -max || z > max {
						continue
					}

					boxes[x+max][y+max][z+max] = in.onoff
				}
			}
		}

	}

	return boxes

}

func parse_instructions(lines []string) []inst {

	instructions := []inst{}

	r := regexp.MustCompile("([onf]+) x=([0-9-]+)..([0-9-]+),y=([0-9-]+)..([0-9-]+),z=([0-9-]+)..([0-9-]+)")

	for _, line := range lines {
		m := r.FindAllStringSubmatch(line, -1)

		i := inst{}

		i.onoff = m[0][1] == "on"

		for n := 2; n < len(m[0]); n++ {
			num, _ := strconv.Atoi(m[0][n])

			if n < 4 {
				i.x = append(i.x, num)
			} else if n < 6 {
				i.y = append(i.y, num)
			} else {
				i.z = append(i.z, num)
			}
		}

		instructions = append(instructions, i)
	}

	return instructions
}

func init_grid(size int) [][][]bool {
	actual_size := 2*size + 1
	g := make([][][]bool, actual_size)

	for x := range g {
		g[x] = make([][]bool, actual_size)
		for y := range g[x] {
			g[x][y] = make([]bool, actual_size)
		}
	}

	return g
}

func count_boxes(boxes [][][]bool) int {
	count := 0

	for _, tab := range boxes {
		for _, row := range tab {
			for _, val := range row {
				if val {
					count++
				}
			}
		}
	}

	return count
}
