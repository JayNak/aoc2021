package aoc

import (
	"fmt"

	"github.com/jaynak/aoc2021/pkg/util"
)

func Day25(path string) (int, int) {

	lines := util.ReadToStrings(path)

	grid := read_sea_cucumbers(lines)

	// print_sea_cucumbers(grid)
	c := 0
	count := 0
	for {
		grid, c = simulate_generation(grid)

		count++
		if c == 0 {
			break
		}

	}

	return count, 0
}

func clone_sea_cucumbers(grid [][]rune) [][]rune {

	duplicate := make([][]rune, len(grid))
	for i := range grid {
		duplicate[i] = make([]rune, len(grid[i]))
		copy(duplicate[i], grid[i])
	}

	return duplicate
}

func read_sea_cucumbers(lines []string) [][]rune {

	rows := len(lines)
	cols := len(lines[0])

	grid := make([][]rune, rows)

	for i, line := range lines {
		grid[i] = make([]rune, cols)
		for j, r := range line {
			grid[i][j] = r
		}
	}

	return grid

}

func simulate_generation(grid [][]rune) ([][]rune, int) {

	changes := 0

	// East
	clone_east := clone_sea_cucumbers(grid)

	for i, row := range grid {
		for j, r := range row {

			// Check the left and current cell
			var prev int
			if j == 0 {
				prev = len(row) - 1
			} else {
				prev = j - 1
			}

			if grid[i][prev] == '>' && r == '.' {
				clone_east[i][prev] = '.'
				clone_east[i][j] = '>'
				changes++
			} else {
				// clone_east[i][prev] = grid[i][prev]
				if clone_east[i][j] != '.' {
					clone_east[i][j] = r
				}

			}
		}
	}

	// South
	clone_south := clone_sea_cucumbers(clone_east)

	for i, row := range clone_east {

		var prev int
		if i == 0 {
			prev = len(clone_east) - 1
		} else {
			prev = i - 1
		}
		for j, r := range row {
			// Check the up and current cell

			if clone_east[prev][j] == 'v' && r == '.' {
				clone_south[prev][j] = '.'
				clone_south[i][j] = 'v'
				changes++
			} else {
				if clone_south[i][j] != '.' {
					clone_south[i][j] = r
				}
			}
		}
	}

	return clone_south, changes
}

func print_sea_cucumbers(grid [][]rune) {
	for _, row := range grid {
		for _, r := range row {
			fmt.Printf("%v", string(r))
		}
		fmt.Print("\n")
	}
	fmt.Print("\n")
}
