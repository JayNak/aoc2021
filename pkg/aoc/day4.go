package aoc

import (
	"strconv"
	"strings"

	"github.com/jaynak/aoc2021/pkg/util"
)

func Day4(path string) (int, int) {
	lines := util.ReadToStrings(path)

	// Create the boards
	boards := []*board{}
	nums := []int{}
	for i := 2; i < len(lines); i++ {
		if lines[i] == "" {
			// This is the end of a board
			boards = append(boards, NewBoard(nums))
			nums = []int{}
		} else {
			strnums := strings.Fields(lines[i])
			for _, v := range strnums {
				thisNum, err := strconv.Atoi(v)
				if err != nil {
					panic(err)
				}
				nums = append(nums, thisNum)
			}
		}
	}

	// Write the last board
	boards = append(boards, NewBoard(nums))
	winOrder := []int{}

	// Iterate through the numbers to call
	callerLineSplit := strings.Split(lines[0], ",")
	for _, v := range callerLineSplit {
		num, err := strconv.Atoi(v)
		if err != nil {
			panic(err)
		}

		for n, b := range boards {
			if b.HasWon() {
				continue
			}

			b.Mark(num)
			if b.HasWon() {
				b.WinningNum = num
				winOrder = append(winOrder, n)
			}
		}
	}

	return boards[winOrder[0]].Score(), boards[winOrder[len(winOrder)-1]].Score()
}

type board struct {
	nums       []int
	marked     []bool
	won        bool
	WinningNum int
}

func NewBoard(nums []int) *board {
	if len(nums) != 25 {
		panic("Invalid number of entries")
	}

	b := board{nums: nums, marked: make([]bool, 25)}

	return &b
}

func (b *board) Mark(num int) {

	for i, v := range b.nums {
		if v == num {
			b.marked[i] = true
		}
	}

	return
}

func (b *board) HasWon() bool {

	if b.won {
		return true
	}

	won := true

	// Horizontals
	for i := 0; i < 5; i++ {
		won = true
		for j := 0; j < 5; j++ {
			if !b.marked[i*5+j] {
				won = false
			}
		}
		if won {
			b.won = true
			return true
		}
	}

	// Verticals
	for i := 0; i < 5; i++ {
		won = true
		for j := 0; j < 5; j++ {
			if !b.marked[j*5+i] {
				won = false
			}
		}
		if won {
			b.won = true
			return true
		}
	}

	// Diagonals
	diag := [][]int{{0, 6, 12, 18, 24}, {4, 8, 12, 16, 20}}

	for _, seq := range diag {
		won = true
		for i := range seq {
			if !b.marked[i] {
				won = false
			}
		}

		if won {
			b.won = true
			return true
		}
	}

	// fell through all conditions
	return false
}

func (b *board) Score() int {

	sum := 0
	for i, v := range b.marked {
		if !v {
			sum += b.nums[i]
		}
	}

	return sum * b.WinningNum
}
