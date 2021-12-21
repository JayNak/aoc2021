package aoc

func GetFunctions() []func(path string) (int, int) {

	fns := make([]func(path string) (int, int), 1)
	fns = append(fns, Day1, Day2, Day3, Day4, Day5, Day6, Day7, Day8, Day9, Day10)
	fns = append(fns, Day11, Day12, Day13, Day14, Day15, Day16, Day17, Day18, Day19, Day20)
	fns = append(fns, Day21)

	return fns
}
