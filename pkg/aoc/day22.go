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

	p2cubes := reduce_cuboids(instructions)

	p2count := 0
	for _, cube := range p2cubes {
		if cube.onoff {
			p2count += cube.size()
		} else {
			p2count -= cube.size()
		}
	}

	return count, p2count
}

func reduce_cuboids(cubes []cuboid) []cuboid {

	new_cubes := []cuboid{}

	for _, cube := range cubes {

		temp_cubes := []cuboid{}
		// Look for overlaps
		for _, new := range new_cubes {
			b, overlap := cube.overlaps(new)

			if !b {
				continue
			}

			if cube.onoff && overlap.onoff {
				overlap.onoff = false
			} else if !cube.onoff && !overlap.onoff {
				overlap.onoff = true
			} else {
				overlap.onoff = cube.onoff
			}

			temp_cubes = append(temp_cubes, overlap)
		}

		if cube.onoff {
			new_cubes = append(new_cubes, cube)
		}

		new_cubes = append(new_cubes, temp_cubes...)
	}

	return new_cubes
}

type cuboid struct {
	onoff bool
	x     []int
	y     []int
	z     []int
}

func (c cuboid) size() int {
	return (c.x[1] - c.x[0] + 1) * (c.y[1] - c.y[0] + 1) * (c.z[1] - c.z[0] + 1)
}

func (c cuboid) overlaps(other cuboid) (bool, cuboid) {

	x1 := util.MaxInt(c.x[0], other.x[0])
	x2 := util.MinInt(c.x[1], other.x[1])

	if x1 > x2 {
		return false, cuboid{}
	}

	y1 := util.MaxInt(c.y[0], other.y[0])
	y2 := util.MinInt(c.y[1], other.y[1])

	if y1 > y2 {
		return false, cuboid{}
	}

	z1 := util.MaxInt(c.z[0], other.z[0])
	z2 := util.MinInt(c.z[1], other.z[1])

	if z1 > z2 {
		return false, cuboid{}
	}

	// Overlap in all 3 dimensions, this one overlaps
	overlap := cuboid{
		onoff: other.onoff,
		x:     []int{x1, x2},
		y:     []int{y1, y2},
		z:     []int{z1, z2},
	}

	return true, overlap
}

func apply_instructions(boxes [][][]bool, instructions []cuboid) [][][]bool {

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
				for z := in.z[0]; z <= in.z[1]; z++ {
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

func parse_instructions(lines []string) []cuboid {

	instructions := []cuboid{}

	r := regexp.MustCompile("([onf]+) x=([0-9-]+)..([0-9-]+),y=([0-9-]+)..([0-9-]+),z=([0-9-]+)..([0-9-]+)")

	for _, line := range lines {
		m := r.FindAllStringSubmatch(line, -1)

		i := cuboid{}

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
