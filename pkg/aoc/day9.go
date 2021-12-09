package aoc

import (
	"sort"

	"github.com/jaynak/aoc2021/pkg/util"
)

func Day9(path string) (int, int) {
	lines := util.ReadToStrings(path)

	grid := util.StringSliceToIntGrid(lines)

	seen := make(map[gridpoint]bool)
	basins := []int{}
	sum := 0
	for i, row := range grid {
		for j, v := range row {
			gp := gridpoint{x: i, y: j}

			if IsLowPoint(grid, gp) {
				sum += v + 1
			}

			// Look for basins
			if _, ok := seen[gp]; !ok && v != 9 {
				basin := MapBasin(grid, gp, seen)
				basins = append(basins, basin)
			}
		}
	}

	sort.Ints(basins)
	size := basins[len(basins)-3]
	for _, v := range basins[len(basins)-2:] {
		size *= v
	}

	return sum, size
}

type gridpoint struct {
	x int
	y int
}

func MapBasin(grid [][]int, gp gridpoint, seen map[gridpoint]bool) int {

	basinsize := 0
	search := []gridpoint{gp}

	for len(search) > 0 {
		newsearch := make(map[gridpoint]bool)
		for _, gp := range search {
			// fmt.Printf("Searching point: %v\n", gp)
			seen[gp] = true
			if grid[gp.x][gp.y] != 9 {
				basinsize++

				n := getGridNeighbours(grid, gp)
				for _, g := range n {
					if _, ok := seen[g]; !ok {
						newsearch[g] = true
					}
				}
			}
		}

		search = []gridpoint{}
		for k := range newsearch {
			search = append(search, k)
		}
	}

	return basinsize

}

func getGridNeighbours(grid [][]int, gp gridpoint) []gridpoint {
	neighbours := []gridpoint{}

	if gp.y != 0 {
		neighbours = append(neighbours, gridpoint{x: gp.x, y: gp.y - 1})
	}
	// Up
	if gp.x != 0 {
		neighbours = append(neighbours, gridpoint{x: gp.x - 1, y: gp.y})
	}
	// Right
	if gp.y < len(grid[gp.x])-1 {
		neighbours = append(neighbours, gridpoint{x: gp.x, y: gp.y + 1})
	}
	// Down
	if gp.x < len(grid)-1 {
		neighbours = append(neighbours, gridpoint{x: gp.x + 1, y: gp.y})
	}

	return neighbours
}

func IsLowPoint(grid [][]int, gp gridpoint) bool {

	neighbours := getGridNeighbours(grid, gp)

	for _, neighbour := range neighbours {
		if grid[neighbour.x][neighbour.y] <= grid[gp.x][gp.y] {
			return false
		}
	}

	return true
}
