package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {

	day1a, day1b := Day1()
	day2a, day2b := Day2()
	day3a, day3b := Day3()
	day4a, day4b := Day4()

	fmt.Printf("Day 1: %v, %v\n", day1a, day1b)
	fmt.Printf("Day 2: %v, %v\n", day2a, day2b)
	fmt.Printf("Day 3: %v, %v\n", day3a, day3b)
	fmt.Printf("Day 4: %v, %v\n", day4a, day4b)

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

func ReadToStrings(fileName string) []string {
	f, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	var ret []string

	for scanner.Scan() {
		ret = append(ret, scanner.Text())
	}

	return ret
}
