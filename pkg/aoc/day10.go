package aoc

import (
	"sort"

	"github.com/jaynak/aoc2021/pkg/util"
)

func Day10(path string) (int, int) {

	lines := util.ReadToStrings(path)

	points := map[rune]int{
		')': 3,
		']': 57,
		'}': 1197,
		'>': 25137,
	}

	acpoints := map[int]int{
		int('('): 1,
		int('['): 2,
		int('{'): 3,
		int('<'): 4,
	}

	sum := 0
	acscores := []int{}
	for _, line := range lines {

		s := util.Stack{}
		ok := true
		for _, r := range line {
			switch r {
			case '[', '(', '{', '<':
				s.Push(int(r))
			case ']', '}', '>':
				// 2 int offset
				if prev := s.Pop(); prev != int(r)-2 {
					ok = false
				}
			case ')':
				// 1 int offset
				if prev := s.Pop(); prev != int(r)-1 {
					ok = false
				}
			}

			if !ok {
				sum += points[r]
				break
			}
		}

		if ok {
			// This is an incomplete line
			score := 0
			for len(s) > 0 {
				char := s.Pop()
				score = 5*score + acpoints[char]
			}
			acscores = append(acscores, score)
		}
	}

	// Find the middle score
	sort.Ints(acscores)
	mid := len(acscores) / 2

	return sum, acscores[mid]
}
