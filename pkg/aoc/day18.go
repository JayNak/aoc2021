package aoc

import (
	"math"
	"strconv"

	"github.com/jaynak/aoc2021/pkg/util"
)

func Day18(path string) (int, int) {

	lines := util.ReadToStrings(path)

	var n *pair
	for _, line := range lines {
		num := ReadSnailfishNumber(line)
		if n == nil {
			n = num
		} else {
			n = n.Add(num)
			n = n.Reduce()
		}
	}

	max := 0

	for i, line := range lines {
		for j, line2 := range lines {
			if i == j {
				continue
			}

			p := ReadSnailfishNumber(line)
			p = p.Add(ReadSnailfishNumber(line2))
			p = p.Reduce()
			m := p.magnitude()

			if m > max {
				max = m
				// fmt.Printf("%v: %v + %v\n", m, line, line2)
			}
		}
	}

	return n.magnitude(), max
}

type pair struct {
	vals        []int
	pairs       []*pair
	parent      *pair
	descendants map[int][]*pair
}

func (p *pair) Print() string {
	ret := "["

	if p.vals[0] != -1 {
		ret += strconv.Itoa(p.vals[0])
	} else {
		ret += p.pairs[0].Print()
	}

	ret += ","

	if p.vals[1] != -1 {
		ret += strconv.Itoa(p.vals[1])
	} else {
		ret += p.pairs[1].Print()
	}

	ret += "]"

	return ret
}

func (p *pair) Add(other *pair) *pair {

	new_pair := &pair{
		vals:  []int{-1, -1},
		pairs: []*pair{p, other},
	}

	new_pair.descendants = make(map[int][]*pair)

	for _, new_child := range new_pair.pairs {
		new_child.parent = new_pair
		new_pair.descendants[1] = append(new_pair.descendants[1], new_child)
		for dep, children := range new_child.descendants {
			new_pair.descendants[dep+1] = append(new_pair.descendants[dep+1], children...)
		}

	}

	return new_pair
}

func find_four_deep_dfs(p *pair, depth int) *pair {

	var ret *pair

	for _, child := range p.pairs {
		if child != nil {
			if depth == 4 {
				return child
			} else {
				ret = find_four_deep_dfs(child, depth+1)
			}
		}

		if ret != nil {
			return ret
		}
	}

	return nil
}

func (p *pair) Reduce() *pair {

	// If any pair is nested inside four pairs, the leftmost such pair explodes.
	// If any regular number is 10 or greater, the leftmost such regular number splits.

	made_change := true

	for made_change == true {
		made_change = false

		// Switch to DFS to find the left-most 4 deep
		exp := find_four_deep_dfs(p, 1)

		// Check for 4 deep
		if exp != nil {

			exp_string := exp.Print()
			full_string := p.Print()
			idx := 0

			// walk through the string to find the correct instance of the string
			open := 0
			for i, r := range full_string {
				if r == '[' {
					open++
					if open == 5 {
						// This is the start
						idx = i
						break
					}
				} else if r == ']' {
					open--
				}
			}

			// idxx := strings.Index(full_string, exp_string)

			left, right := -1, -1
			done := false

			for i := idx; i >= 0; i-- {
				switch full_string[i] {
				case '[', ']', ',':
					// Skip
				default:
					// This is the first number
					if r := full_string[i-1]; r != '[' && r != ']' && r != ',' {
						left, _ = strconv.Atoi(string(full_string[i-1 : i+1]))
						new := left + exp.vals[0]
						full_string = full_string[:i-1] + strconv.Itoa(new) + full_string[i+1:]
					} else {
						left, _ = strconv.Atoi(string(full_string[i]))
						new := left + exp.vals[0]
						full_string = full_string[:i] + strconv.Itoa(new) + full_string[i+1:]
					}

					done = true
				}

				if done {
					break
				}
			}

			done = false

			for i := idx + len(exp_string); i < len(full_string); i++ {
				switch full_string[i] {
				case '[', ']', ',':
					// Skip
				default:
					// This is the first number
					if r := full_string[i+1]; r != '[' && r != ']' && r != ',' {
						right, _ = strconv.Atoi(string(full_string[i : i+2]))
						new := right + exp.vals[1]
						full_string = full_string[:i] + strconv.Itoa(new) + full_string[i+2:]
					} else {
						right, _ = strconv.Atoi(string(full_string[i]))
						new := right + exp.vals[1]
						full_string = full_string[:i] + strconv.Itoa(new) + full_string[i+1:]
					}

					done = true

				}

				if done {
					break
				}
			}

			if left >= 0 && left < 10 && left+exp.vals[0] > 9 {
				idx++
			}

			full_string = full_string[:idx] + "0" + full_string[idx+len(exp_string):]
			p = ReadSnailfishNumber(full_string)
			made_change = true

			// fmt.Printf("EXP: %v => %v\n", exp_string, full_string)

		}

		if !made_change {
			made_change = p.Split()
			// if made_change {
			// 	fmt.Printf("%v\n", p.Print())
			// }
		}
	}

	return p
}

func (p *pair) Split() bool {

	// Check left first
	if p.pairs[0] != nil {
		b := p.pairs[0].Split()
		if b {
			return true
		}
	}

	for i, v := range p.vals {
		if v > 9 {
			// fmt.Printf("SPL: %v => ", v)
			// Split this one
			new := &pair{
				vals: []int{
					int(math.Floor(float64(v) / 2)),
					int(math.Ceil(float64(v) / 2)),
				},
				pairs:       make([]*pair, 2),
				parent:      p,
				descendants: make(map[int][]*pair),
			}

			p.pairs[i] = new
			p.vals[i] = -1

			// Update descendants
			tmp := p
			i := 1
			for tmp != nil {
				tmp.descendants[i] = append(tmp.descendants[i], new)
				i++
				tmp = tmp.parent
			}

			return true
		}
	}

	// Check right last first
	if p.pairs[1] != nil {
		b := p.pairs[1].Split()
		if b {
			return true
		}
	}

	return false
}

func (p *pair) magnitude() int {
	mag := 0

	for i, v := range p.vals {
		if v != -1 {
			mag += (3 - i) * v
		} else {
			mag += (3 - i) * p.pairs[i].magnitude()
		}
	}

	return mag
}

func ReadSnailfishNumber(line string) *pair {

	open, mid := 0, 0

	// Skip the first character
	for pos, r := range line[1:] {
		switch r {
		case '[':
			open++
		case ']':
			open--
		case ',':
			if open == 0 {
				// This is the middle
				mid = pos
				break
			}
		}
	}

	parts := []string{
		line[1 : mid+1],
		line[mid+2 : len(line)-1],
	}

	p := &pair{
		vals:        []int{-1, -1},
		pairs:       make([]*pair, 2),
		descendants: make(map[int][]*pair),
	}

	for i, part := range parts {
		if _, err := strconv.Atoi(part); err == nil {
			p.vals[i], _ = strconv.Atoi(part)
		} else {
			p.pairs[i] = ReadSnailfishNumber(part)
			p.pairs[i].parent = p
			p.descendants[1] = append(p.descendants[1], p.pairs[i])

			for dep, children := range p.pairs[i].descendants {
				p.descendants[dep+1] = append(p.descendants[dep+1], children...)
			}
		}
	}

	return p

}
