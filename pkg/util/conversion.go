package util

import (
	"strconv"
	"strings"
)

func StringToIntSlice(s string, sep string) []int {
	n := []int{}

	slice := strings.Split(s, sep)
	for _, v := range slice {
		i, err := strconv.Atoi(v)
		if err != nil {
			panic(err)
		}
		n = append(n, i)
	}

	return n
}

func IntAbs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func StringSliceToIntGrid(lines []string) [][]int {
	var grid [][]int

	for _, line := range lines {
		row := []int{}
		for _, r := range line {
			row = append(row, int(r-'0'))
		}
		grid = append(grid, row)
	}

	return grid
}
