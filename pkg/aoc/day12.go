package aoc

import (
	"regexp"
	"unicode"

	"github.com/jaynak/aoc2021/pkg/util"
)

func Day12(path string) (int, int) {

	lines := util.ReadToStrings(path)

	caves := MapCave(lines)

	n := caves["start"].CountPaths("end", []string{})
	m := caves["start"].CountPaths2("end", []string{}, false)

	return n, m
}

type cave struct {
	name    string
	adj     []*cave
	isLarge bool
}

func MapCave(lines []string) map[string]*cave {
	caves := make(map[string]*cave)

	r := regexp.MustCompile("([a-zA-Z]+)-([a-zA-Z]+)")

	for _, line := range lines {
		m := r.FindAllStringSubmatch(line, -1)

		for i := 1; i < 3; i++ {
			nm := m[0][i]
			if _, ok := caves[nm]; !ok {
				c := &cave{
					name:    nm,
					adj:     []*cave{},
					isLarge: unicode.IsUpper(rune(nm[0])),
				}
				caves[nm] = c
			}
		}

		caves[m[0][1]].adj = append(caves[m[0][1]].adj, caves[m[0][2]])
		caves[m[0][2]].adj = append(caves[m[0][2]].adj, caves[m[0][1]])
	}

	return caves
}

func (c *cave) CountPaths(target string, path []string) int {

	count := 0
	path2 := path
	path2 = append(path2, c.name)

	for _, adj := range c.adj {

		if adj.name == target {
			count++
		} else {
			if adj.isLarge {
				// fmt.Printf("From: %v to large %v\n", c.name, adj.name)
				count += adj.CountPaths(target, path2)
			} else if !InSlice(path2, adj.name) {
				// fmt.Printf("From: %v to small %v\n", c.name, adj.name)
				count += adj.CountPaths(target, path2)
			}
		}
	}

	return count

}

func InSlice(sl []string, val string) bool {
	for _, v := range sl {
		if v == val {
			return true
		}
	}

	return false
}

func (c *cave) CountPaths2(target string, path []string, secondsmall bool) int {

	count := 0
	path2 := path
	path2 = append(path2, c.name)

	for _, adj := range c.adj {

		if adj.name == target {
			count++
			continue
		}

		if adj.name == "start" {
			continue
		}

		if adj.isLarge {
			count += adj.CountPaths2(target, path2, secondsmall)
			continue
		}

		// At this point it's a small cave adjacent
		adjCount := CountInSlice(path2, adj.name)

		switch adjCount {
		case 0:
			count += adj.CountPaths2(target, path2, secondsmall)
		case 1:
			if !secondsmall {
				count += adj.CountPaths2(target, path2, true)
			}
		}
	}

	return count

}

func CountInSlice(sl []string, val string) int {
	cnt := 0
	for _, v := range sl {
		if v == val {
			cnt++
		}
	}

	return cnt
}
