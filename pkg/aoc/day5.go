package aoc

import (
	"regexp"
	"strconv"

	"github.com/jaynak/aoc2021/pkg/util"
)

func Day5(path string) (int, int) {

	lines := util.ReadToStrings(path)
	vectors, grid_size := InterpretLines(lines)

	g := make(grid, grid_size)
	for i := range g {
		g[i] = make([]int, grid_size)
	}

	orth, diag := SplitLines(vectors)

	overlaps := 0

	for _, vec := range orth {
		overlaps += g.PlotLine(vec)
	}

	// Add the diagonals
	diagoverlaps := overlaps
	for _, vec := range diag {
		diagoverlaps += g.PlotLine(vec)
	}

	// fmt.Println(g)

	return overlaps, diagoverlaps
}

type grid [][]int

func (g grid) PlotLine(v *vector) int {

	overlaps := 0

	// Horizontal or Vertical
	if v.x1 == v.x2 || v.y1 == v.y2 {

		for i := v.x1; i <= v.x2; i++ {
			for j := v.y1; j <= v.y2; j++ {
				g[i][j]++
				if g[i][j] == 2 {
					overlaps++
				}
			}
		}

		return overlaps
	}

	// Diagonals - quick validation HACK
	if v.x2-v.x1 != v.y2-v.y1 && v.x2-v.x1 != v.y1-v.y2 {
		panic("Not a 45 degree diagonal!")
	}

	// We know that x is increasing
	j := v.y1
	for i := v.x1; i <= v.x2; i++ {
		g[i][j]++
		if g[i][j] == 2 {
			overlaps++
		}

		if v.y1 < v.y2 {
			j++
		} else {
			j--
		}

	}

	return overlaps

}

type vector struct {
	x1, x2 int
	y1, y2 int
}

func InterpretLines(lines []string) ([]*vector, int) {
	ret := []*vector{}
	max := 0

	r := regexp.MustCompile("([0-9]+),([0-9]+) -> ([0-9]+),([0-9]+)")

	for _, line := range lines {

		// x1,y1 -> x2,y2
		m := r.FindAllStringSubmatch(line, -1)

		v := []int{}
		for i := 1; i < len(m[0]); i++ {
			n, err := strconv.Atoi(m[0][i])
			if err != nil {
				panic(err)
			}
			v = append(v, n)

			if n > max {
				max = n
			}
		}

		// Orient the line here
		vec := &vector{x1: v[0], y1: v[1], x2: v[2], y2: v[3]}

		// For horizontal lines, align it for plotting later
		if vec.x1 == vec.x2 || vec.y1 == vec.y2 {
			if vec.x1 > vec.x2 {
				vec.x1, vec.x2 = vec.x2, vec.x1
			}

			if vec.y1 > vec.y2 {
				vec.y1, vec.y2 = vec.y2, vec.y1
			}
		} else {
			// Make it so x increases
			if vec.x1 > vec.x2 {
				vec.x1, vec.x2 = vec.x2, vec.x1
				vec.y1, vec.y2 = vec.y2, vec.y1
			}
		}

		ret = append(ret, vec)
	}

	return ret, max + 1
}

// Orthoganal, Diag
func SplitLines(vectors []*vector) ([]*vector, []*vector) {
	orth := []*vector{}
	diag := []*vector{}

	for _, v := range vectors {
		if v.x1 == v.x2 || v.y1 == v.y2 {
			orth = append(orth, v)
		} else {
			diag = append(diag, v)
		}
	}

	return orth, diag
}
