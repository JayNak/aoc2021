package aoc

func GetFunctions() []func(path string) (int, int) {

	fns := make([]func(path string) (int, int), 1)
	fns = append(fns, Day1, Day2, Day3, Day4, Day5, Day6, Day7, Day8, Day9, Day10)
	fns = append(fns, Day11)

	return fns
}
