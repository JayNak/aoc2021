package aoc

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/jaynak/aoc2021/pkg/util"
)

func Day13(path string) (int, int) {
	lines := util.ReadToStrings(path)

	plot, linebreak := PlotDots(lines)
	count := 0

	// Read folding instructions
	for i := linebreak + 1; i < len(lines); i++ {
		f := strings.Fields(lines[i])
		coords := strings.FieldsFunc(f[2], func(r rune) bool { return r == '=' })
		val, _ := strconv.Atoi(coords[1])

		plot = FoldPaper(plot, rune(coords[0][0]), val)

		if i == linebreak+1 {
			for _, row := range plot {
				for _, v := range row {
					if v > 0 {
						count++
					}
				}
			}
		}
	}

	fmt.Printf("Day 13: %v\n", path)

	for _, row := range plot {
		for _, v := range row {
			if v > 0 {
				fmt.Printf("X ")
			} else {
				fmt.Printf(". ")
			}
		}
		fmt.Println()
	}
	fmt.Println()

	return count, 0
}

func FoldPaper(plot [][]int, dir rune, val int) [][]int {

	var new_paper [][]int

	if dir == 'x' {
		new_paper = make([][]int, len(plot))
		for i := range new_paper {
			new_paper[i] = make([]int, val)

			for j := 0; j < val; j++ {
				new_paper[i][j] = plot[i][j]

				if 2*val-j < len(plot[i]) {
					new_paper[i][j] += plot[i][2*val-j]
				}

			}
		}
	} else {
		new_paper = make([][]int, val)
		for i := range new_paper {
			new_paper[i] = make([]int, len(plot[0]))

			for j := range new_paper[i] {
				if 2*val-i < len(plot) {
					new_paper[i][j] = plot[i][j] + plot[2*val-i][j]
				}
			}
		}
	}

	return new_paper
}

func PlotDots(lines []string) ([][]int, int) {

	point := [][]int{}

	max_x := 0
	max_y := 0
	last := 0

	for i, line := range lines {
		if line == "" {
			// This is the break between the plotted points and folding instructions
			last = i
			break
		}

		split := strings.FieldsFunc(line, func(r rune) bool { return r == ',' })

		x, _ := strconv.Atoi(split[0])
		y, _ := strconv.Atoi(split[1])

		if x > max_x {
			max_x = x
		}

		if y > max_y {
			max_y = y
		}

		point = append(point, []int{x, y})
	}

	plot := make([][]int, max_y+1)
	for i := 0; i < max_y+1; i++ {
		plot[i] = make([]int, max_x+1)
	}

	for _, v := range point {
		plot[v[1]][v[0]] = 1
	}

	return plot, last
}
