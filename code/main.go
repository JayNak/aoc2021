package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {

	day1a, day1b := Day1()

	fmt.Printf("Day 1: %v, %v\n", day1a, day1b)

}

func ReadToInts(fileName string) []int {
	f, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	var ret []int

	for scanner.Scan() {
		i, err := strconv.Atoi(scanner.Text())
		if err != nil {
			panic(err)
		}

		ret = append(ret, i)
	}

	return ret
}
