package aoc

import (
	"math"
	"strconv"
	"strings"

	"github.com/jaynak/aoc2021/pkg/util"
)

func Day15(path string) (int, int) {

	cave := util.ReadToIntArray(path)

	bigcave := BuildBiggerCave(cave, 5)

	// sml := FindPath(cave)
	sml2 := AStar(cave)

	// big := FindPath(bigcave)
	big2 := AStar(bigcave)

	// return sml, big
	return sml2, big2

}

type node struct {
	i       int
	j       int
	parent  *node
	f       int
	g       int
	visited bool
}

type agrid [][]*node

func (g agrid) GetNode(i int, j int) *node {
	if g[i][j] == nil {
		g[i][j] = &node{
			i:      i,
			j:      j,
			parent: nil,
			f:      0,
			g:      0,
		}
	}

	return g[i][j]
}

func (g agrid) GetFromSlug(slug string) *node {
	s := strings.Split(slug, "_")
	i, _ := strconv.Atoi(s[0])
	j, _ := strconv.Atoi(s[1])
	return g.GetNode(i, j)
}

func (g agrid) GetNeighbours(n *node, cave [][]int) []*node {

	size := len(g)
	ret := []*node{}

	try := [][]int{
		{n.i - 1, n.j},
		{n.i, n.j - 1},
		{n.i + 1, n.j},
		{n.i, n.j + 1},
	}

	for _, t := range try {
		if t[0] < 0 || t[1] < 0 || t[0] >= size || t[1] >= size {
			// Invalid, skip
			continue
		}

		ret = append(ret, g.GetNode(t[0], t[1]))
	}

	return ret
}

func slug(n *node) string {
	return strconv.Itoa(n.i) + "_" + strconv.Itoa(n.j)
}

func AStar(cave [][]int) int {

	// Cave is a square
	size := len(cave)
	g := make(agrid, size)
	for i := range g {
		g[i] = make([]*node, size)
	}

	openList := make(map[*node]bool)
	openList[g.GetNode(0, 0)] = true

	for len(openList) > 0 {

		// loop through the openlist to find the smallest f
		f := math.MaxInt32

		var curr_node *node
		for n := range openList {
			if n.f < f {
				curr_node = n
				f = n.f
			}
		}

		curr_node.visited = true
		delete(openList, curr_node)

		if curr_node.i == size-1 && curr_node.j == size-1 {
			// This is the node!
			break
		}

		children := g.GetNeighbours(curr_node, cave)

		for _, child := range children {
			if child.visited {
				continue
			}

			new_g := curr_node.g + cave[child.i][child.j]

			// Is this already on the openList ?
			if _, ok := openList[child]; !ok {
				openList[child] = true
				child.g = new_g
				child.parent = curr_node
				child.f = child.g
				// + int(math.Sqrt(math.Pow(float64(size-1-child.i), 2)+math.Pow(float64(size-1-child.j), 2)))
			} else {
				// If this is a shorter path take it
				if new_g < child.g {
					child.g = new_g
					child.parent = curr_node
					child.f = child.g
					// + int(math.Sqrt(math.Pow(float64(size-1-child.i), 2))+math.Pow(float64(size-1-child.j), 2))
				}
			}
		}
	}

	n := g.GetNode(size-1, size-1)

	// for n != nil {
	// 	fmt.Printf("[%v, %v]: %v\n", n.i, n.j, n.g)
	// 	n = n.parent
	// }

	return n.g
}

func FindPath(cave [][]int) int {

	// Cave is square
	size := len(cave)

	memo := make([][]int, size)
	for i := range memo {
		memo[i] = make([]int, size)
	}

	for i := size - 1; i >= 0; i-- {
		for j := size - 1; j >= 0; j-- {
			if i == size-1 && j == size-1 {
				// Base case
				memo[i][j] = cave[i][j]
				continue
			}

			if i == size-1 {
				// Can't check down
				memo[i][j] = cave[i][j] + memo[i][j+1]
				continue
			}

			if j == size-1 {
				// Can't check right
				memo[i][j] = cave[i][j] + memo[i+1][j]
				continue
			}

			memo[i][j] = cave[i][j] + min(memo[i+1][j], memo[i][j+1])
		}
	}

	return memo[0][0] - cave[0][0]
}

func min(x int, y int) int {
	if x < y {
		return x
	}
	return y
}

func BuildBiggerCave(cave [][]int, multiplier int) [][]int {

	size := len(cave)
	newsize := len(cave) * multiplier

	newcave := make([][]int, newsize)
	for i := range newcave {
		newcave[i] = []int{}
	}

	bumped := [][][]int{cave}

	for m := 0; m < multiplier; m++ {
		for n := 0; n < multiplier; n++ {

			if len(bumped) < m+n+1 {
				bumped = append(bumped, BumpCave(bumped[m+n-1]))
			}

			for i := range cave {
				newcave[i+m*size] = append(newcave[i+m*size], bumped[m+n][i]...)

			}
		}
	}

	return newcave
}

func BumpCave(cave [][]int) [][]int {
	size := len(cave)

	bumped := make([][]int, size)

	for i, row := range cave {
		bumped[i] = make([]int, size)

		for j, val := range row {
			if val == 9 {
				bumped[i][j] = 1
			} else {
				bumped[i][j] = val + 1
			}
		}
	}

	return bumped
}
