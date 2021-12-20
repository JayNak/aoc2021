package aoc

import (
	"strconv"

	"github.com/jaynak/aoc2021/pkg/util"
)

func Day20(path string) (int, int) {

	lines := util.ReadToStrings(path)

	p1count := 0
	count := 0

	key := lines[0]
	image := read_initial_image(lines[2:])

	new_image := image
	for i := 0; i < 50; i++ {
		pad_with := '0'
		if i > 0 && new_image[0][0] == '1' {
			pad_with = '1'
		}
		new_image = enhance_image(new_image, key, pad_with)

		if i == 1 {
			for _, row := range new_image {
				for _, v := range row {
					if v == '1' {
						p1count++
					}
				}
			}
		}
	}

	for _, row := range new_image {
		for _, v := range row {
			if v == '1' {
				count++
			}
		}
	}

	return p1count, count
}

func enhance_image(image [][]rune, key string, pad_with rune) [][]rune {

	new_size := len(image) + 2
	output := make([][]rune, new_size)
	for i := range output {
		output[i] = make([]rune, new_size)
	}

	expanded := expand_image(image, 2, pad_with)

	for i := 1; i < len(expanded)-1; i++ {
		for j := 1; j < len(expanded)-1; j++ {
			valstring := ""
			for x := i - 1; x < i+2; x++ {
				for y := j - 1; y < j+2; y++ {
					valstring += string(expanded[x][y])
				}
			}
			// Interpret
			map_int, _ := strconv.ParseInt(valstring, 2, 64)
			if key[map_int] == '#' {
				output[i-1][j-1] = '1'
			} else {
				output[i-1][j-1] = '0'
			}
		}
	}

	return output
}

func read_initial_image(lines []string) [][]rune {

	ret := make([][]rune, len(lines))

	for i, line := range lines {
		ret[i] = make([]rune, len(line))

		for j, r := range line {
			if r == '#' {
				ret[i][j] = '1'
			} else {
				ret[i][j] = '0'
			}
		}
	}

	return ret
}

func expand_image(source [][]rune, expand_by int, pad_with rune) [][]rune {

	new_size := len(source) + 2*expand_by

	source = append(make([][]rune, expand_by), append(source, make([]rune, new_size))...)
	source = append(source, make([]rune, expand_by))

	for i := 0; i < expand_by; i++ {
		source[i] = make([]rune, new_size)
		source[len(source)-1-i] = make([]rune, new_size)
	}

	for i, row := range source {
		if i < expand_by || i >= len(source)-expand_by {
			// Fill the new rows with zeros
			for j := range row {
				row[j] = pad_with
			}
			continue
		}

		tmp := make([]rune, 2*expand_by)
		for i := range tmp {
			tmp[i] = pad_with
		}

		source[i] = append(tmp[:expand_by], source[i]...)
		source[i] = append(source[i], tmp[expand_by:]...)
	}

	return source
}
