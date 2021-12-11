package util

import (
	"bufio"
	"os"
	"strconv"
)

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

func ReadToIntArray(fileName string) [][]int {
	f, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	var ret [][]int

	for scanner.Scan() {
		line := []int{}
		for _, r := range scanner.Text() {
			i := int(r) - '0'
			line = append(line, i)
		}

		ret = append(ret, line)
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
