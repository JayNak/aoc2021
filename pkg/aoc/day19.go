package aoc

import (
	"fmt"
	"math"
	"regexp"
	"strconv"

	"github.com/jaynak/aoc2021/pkg/util"
)

func Day19(path string) (int, int) {
	lines := util.ReadToStrings(path)

	ptr := 0
	count := 0
	scanners := []*scanner{}
	scanners_to_retry := []*scanner{}

	for i, line := range lines {
		if line == "" {
			// End of a line
			s := read_scanner_data(lines[ptr:i])

			if len(scanners) > 0 {
				b := scanners[0].merge_scanner(s)
				if !b {
					scanners_to_retry = append(scanners_to_retry, s)
				}
			}

			scanners = append(scanners, s)
			ptr = i + 1
		}
	}

	// Get the last one
	s := read_scanner_data(lines[ptr:])
	scanners = append(scanners, s)

	scan_list := scanners
	max := 0
	for 1 == 1 {
		root, leftover, manhattan_distance := reduce_scanners(scan_list)
		if manhattan_distance > max {
			max = manhattan_distance
		}
		count += len(root.beacons)

		if len(scan_list) != len(leftover) {
			scan_list = leftover
			if len(leftover) == 0 {
				break
			}
		} else {
			break
		}
	}

	return count, max
}

func reduce_scanners(scanners []*scanner) (*scanner, []*scanner, int) {

	root := scanners[0]
	merged := []*scanner{root}

	open := scanners[1:]
	next := []*scanner{}

	for true {
		next = []*scanner{}
		for _, s := range open {

			b := root.merge_scanner(s)
			if !b {
				next = append(next, s)
			} else {
				merged = append(merged, s)
			}
		}

		if len(next) == 0 || len(next) == len(open) {
			break
		}

		open = next
	}

	// calculate the largest manhattan distance
	max := 0
	for _, s1 := range merged {
		for _, s2 := range merged {
			manhattan_distance := 0
			for i := range s1.coords {
				manhattan_distance += int(math.Abs(float64(s1.coords[i]) - float64(s2.coords[i])))
			}
			if manhattan_distance > max {
				max = manhattan_distance
			}
		}
	}

	return root, next, max
}

type beacon struct {
	coords []int
	others map[float64]*beacon
}

type scanner struct {
	name    string
	coords  []int
	beacons []*beacon
}

func (s *scanner) Print() {
	fmt.Println(s.name)

	for _, b := range s.beacons {
		fmt.Println(b.coords)

		for k, v := range b.others {
			fmt.Printf("%v, %v\n", k, v.coords)
		}
	}
}

func (s *scanner) calculate_distances() *scanner {

	// Calculate distance between each one
	for _, b1 := range s.beacons {
		b1.others = make(map[float64]*beacon)
		for _, b2 := range s.beacons {
			if b1 == b2 {
				continue
			}

			// Calculate the distance between these 2 beacons
			dist := 0.0
			for i := range b1.coords {
				dist += math.Pow((float64(b1.coords[i]) - float64(b2.coords[i])), 2)
			}
			dist = math.Sqrt(dist)
			b1.others[dist] = b2
		}
	}

	return s
}

func (s *scanner) merge_scanner(other *scanner) bool {

	s.coords = []int{0, 0, 0}

	// Find matching beacons
	for _, b1 := range s.beacons {
		for _, b2 := range other.beacons {

			cnt := 0
			for dist := range b1.others {
				if _, ok := b2.others[dist]; ok {
					cnt++
				}
			}

			// fmt.Printf("%v <=> %v : %v\n", b1.coords, b2.coords, cnt)

			if cnt > 10 {
				// These are the same beacon

				// Align the second scanner (in terms of x, y, z)
				translator := []int{0, 0, 0}
				negate := []bool{false, false, false}
				s2_coords := []int{0, 0, 0}

				for dist := range b1.others {
					// Just need 1 matching one
					if _, ok := b2.others[dist]; !ok {
						continue
					}

					// we have a matching one
					b1_offsets := []int{}
					b2_offsets := []int{}
					for i := range b1.coords {
						b1_offsets = append(b1_offsets, b1.coords[i]-b1.others[dist].coords[i])
						b2_offsets = append(b2_offsets, b2.coords[i]-b2.others[dist].coords[i])
					}

					for i, b1_offset := range b1_offsets {
						for j, b2_offset := range b2_offsets {
							found := false
							if b2_offset == b1_offset {
								translator[j] = i
								// Subtract b2 from b1 to find center of s2
								s2_coords[j] = b1.coords[i] - b2.coords[j]
								found = true
							} else if int(math.Abs(float64(b2_offset))) == int(math.Abs(float64(b1_offset))) {
								// Add b2 to b1 to find center of s2
								translator[j] = i
								s2_coords[j] = b1.coords[i] + b2.coords[j]
								negate[j] = true
								found = true
							}

							if found {
								break
							}
						}

					}

					// We have the mapping at this point
					break
				}

				// fmt.Printf("Scanner %v is @ %v\n", other.name, s2_coords)

				for i := range s2_coords {
					other.coords[translator[i]] = s2_coords[i]
				}

				// Add the second scanner's beacons to the first one
				for _, b := range other.beacons {

					new := &beacon{
						coords: []int{0, 0, 0},
						others: make(map[float64]*beacon),
					}

					// x: change sign - s2
					// y: - s2
					// z: change sign + s2

					for i := range b.coords {
						num := b.coords[i]
						if negate[i] {
							num = -num
						}
						num += s2_coords[i]
						new.coords[translator[i]] = num
					}

					// Check to see if we already have this one
					match := true
					for _, b1 := range s.beacons {

						match = true
						for i := range b1.coords {
							if b1.coords[i] != new.coords[i] {
								match = false
								break
							}
						}

						if match {
							break
						}
					}

					if !match {
						// fmt.Printf("Adding %v\n", new.coords)
						s.beacons = append(s.beacons, new)
					}
				}

				s.calculate_distances()

				// Done
				return true
			}

		}
	}

	return false
}

func read_scanner_data(lines []string) *scanner {

	s := &scanner{
		coords: []int{0, 0, 0},
	}
	r := regexp.MustCompile("([0-9-])+")
	r_title := regexp.MustCompile("([0-9]+)")

	for _, line := range lines {

		if line[0:3] == "---" {
			// This is the name
			m_title := r_title.FindAllStringSubmatch(line, -1)
			s.name = m_title[0][0]
		} else {
			// This is a beacon line
			m := r.FindAllStringSubmatch(line, -1)

			b := &beacon{
				others: make(map[float64]*beacon),
			}
			for _, loc := range m {
				n, _ := strconv.Atoi(loc[0])

				b.coords = append(b.coords, n)
			}

			s.beacons = append(s.beacons, b)
		}
	}

	s = s.calculate_distances()

	return s
}
