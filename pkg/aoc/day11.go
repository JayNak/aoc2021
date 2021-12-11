package aoc

import "github.com/jaynak/aoc2021/pkg/util"

func Day11(path string) (int, int) {
	grid := util.ReadToIntArray(path)

	og := NewOctoGrid(grid)

	flashes := 0
	allflash := 0
	flashcount := 0
	gencount := 1
	for allflash == 0 {
		thisgen := og.RunGeneration()
		flashcount += thisgen

		if thisgen == len(og.grid) {
			allflash = gencount
		}

		if gencount == 100 {
			flashes = flashcount
		}

		gencount++
	}

	return flashes, allflash
}

type Octogrid struct {
	grid       []int
	neighbours [][]int
	rowlen     int
}

func NewOctoGrid(input [][]int) Octogrid {

	// Given a 2-d grid, flatten it and build out the neighbours list
	size := len(input) * len(input[0])
	og := Octogrid{grid: make([]int, size), neighbours: make([][]int, size)}

	og.rowlen = len(input[0])
	for i, row := range input {
		for j, val := range row {
			og.grid[og.idx(i, j)] = val

			// Ugly neighbour things
			neigh := []int{}
			for m := i - 1; m <= i+1; m++ {
				if m < 0 || m >= len(input) {
					continue
				}

				for n := j - 1; n <= j+1; n++ {
					if n < 0 || n >= len(input[0]) {
						continue
					}
					if m == i && n == j {
						continue
					}

					neigh = append(neigh, og.idx(m, n))
				}
			}

			og.neighbours[og.idx(i, j)] = neigh

		}
	}

	return og
}

func (o Octogrid) idx(i int, j int) int {
	return i*o.rowlen + j
}

func (o Octogrid) RunGeneration() int {

	seen := make(map[int]bool)
	flashes := []int{}

	// increase the energy level for each
	for i := range o.grid {
		o.grid[i]++
		if o.grid[i] > 9 {
			flashes = append(flashes, i)
			seen[i] = true
		}
	}

	for len(flashes) > 0 {
		f2 := []int{}

		for _, n := range flashes {
			for _, j := range o.neighbours[n] {
				o.grid[j]++
				if _, ok := seen[j]; !ok && o.grid[j] > 9 {
					f2 = append(f2, j)
					seen[j] = true
				}
			}
		}

		flashes = f2
	}

	for i := range seen {
		o.grid[i] = 0
	}

	return len(seen)
}
