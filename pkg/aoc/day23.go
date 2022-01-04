package aoc

import (
	"fmt"
	"math"
	"strings"

	"github.com/jaynak/aoc2021/pkg/util"
)

func Day23(path string) (int, int) {

	// #D#C#B#A#
	// #D#B#A#C#

	lines := util.ReadToStrings(path)

	start := read_room(lines)
	target := rooms("00112233       ")

	cost, _ := a_star(start, target)

	inserts := []string{
		"  #D#C#B#A#",
		"  #D#B#A#C#",
		lines[3],
		lines[4],
	}

	new_lines := append(lines[:3], inserts...)

	start2 := read_room(new_lines)
	target2 := rooms("0000111122223333       ")

	cost2, _ := a_star(start2, target2)

	// for _, rm := range room_path2 {
	// 	rm.print()
	// }

	return cost, cost2
}

type rooms string

func (r rooms) room_size() int {
	room_size := 4
	if len(r) == 15 {
		room_size = 2
	}

	return room_size
}

func a_star(start rooms, target rooms) (int, []rooms) {

	f := make(map[rooms]int)
	f[start] = start.weight_to_complete()
	g := make(map[rooms]int)
	g[start] = 0
	path := make(map[rooms]rooms)

	open_set := make(map[rooms]int)
	open_set[start] = f[start]

	for len(open_set) > 0 {

		room := rooms("")
		min_f := math.MaxInt

		for k, v := range open_set {
			if v < min_f {
				min_f = v
				room = k
			}
		}

		cost_to_get_here := g[room]

		// Remove from the open set
		delete(open_set, room)

		if room == target {
			// Reconstruct Path
			tmp := target
			return_path := []rooms{target}
			for tmp != start {
				return_path = append([]rooms{path[tmp]}, return_path...)
				tmp = path[tmp]
			}

			return g[room], return_path
		}

		next := room.find_valid_moves()
		for next_room, move_cost := range next {
			val, ok := g[next_room]

			if !ok || (ok && move_cost+cost_to_get_here < val) {
				g[next_room] = move_cost + cost_to_get_here
				f[next_room] = g[next_room] + next_room.weight_to_complete()

				if _, ok := open_set[next_room]; !ok {
					open_set[next_room] = f[next_room]
					path[next_room] = room
				}
			}
		}
	}

	return -1, []rooms{}
}

func (r rooms) weight_to_complete() int {

	travel_cost := []int{1, 10, 100, 1000}
	cost := 0

	room_size := r.room_size()

	for spot, char := range r {
		if char == ' ' {
			continue
		}

		// tmp := '0'

		pod := int(char - '0')
		if pod == spot/room_size {
			// already in position
			continue
		}

		// Calculate distance to first spot in the room
		// May need to be more precise on this later

		// spot % room_size + 1 = position in room
		// Distance to other room = position + diff(position, target) * 2 + 1

		// If in room
		dist := spot%room_size + int(math.Abs(float64(spot/room_size-pod)))*2 + 2

		if spot >= room_size*4 {
			n := spot - (room_size * 4) - 2
			// Distance to room = diff(position, target) + 2

			if n < pod {
				dist = (pod - n) * 2
			} else {
				dist = (n-pod)*2 + 2
			}
			// dist = int(math.Abs(float64(n-pod)))*2 + 2

			if n == -2 || n == 4 {
				dist--
			}
		}

		cost += dist * travel_cost[pod]

	}

	return cost
}

func (r rooms) find_valid_moves() map[rooms]int {

	travel_cost := []int{1, 10, 100, 1000}
	room_size := r.room_size()

	ret := make(map[rooms]int)

	// Find open nodes

	// Pods in rooms to hallway (or to final room)
	for room := range travel_cost {
		loc := -1
		for spot := 0; spot < room_size; spot++ {
			pos := room*room_size + spot
			if r[pos] != ' ' {
				// This is the available pod
				loc = pos
				break

			}
		}

		if loc == -1 {
			continue
		}

		target_room := int(r[loc] - '0')

		// Target Room?
		target_spot, ok := r.target_room_open(target_room)

		if ok {

			// if this pod is already in the right spot
			if target_room == loc/room_size {
				continue
			}

			blocked := false
			// Moving to the left
			if target_room < room {
				// Is there an open path to the left ?
				for j := room; j >= target_room; j-- {
					if r[4*room_size+2+j] != ' ' {
						blocked = true
						break
					}
				}
			} else {
				// Moving to the right

				// Is there an open path to the right ?
				for j := room; j < target_room; j++ {
					if r[4*room_size+2+j] != ' ' {
						blocked = true
						break
					}
				}
			}

			if !blocked {
				dist := 2 + 2*int(math.Abs(float64(target_room-room))) + target_spot + loc%room_size
				new_room := r.swap(loc, target_room*room_size+target_spot)
				ret[new_room] = dist * travel_cost[target_room]
			}

		}

		// Left
		for i := 4*room_size + loc/room_size + 1; i >= 4*room_size; i-- {
			if r[i] != ' ' {
				// space is occupied
				break
			}

			new_room := r.swap(i, loc)

			// Travelling left, abs(room-hall)*2 (+1 if at the end)  hall is shifted, add position in room
			curr_room_no := loc / room_size
			shifted := i - 4*room_size - 2

			dist := int(math.Abs(float64(curr_room_no-shifted)))*2 + loc%room_size
			if shifted == -2 {
				dist--
			}

			ret[new_room] = dist * travel_cost[target_room]

		}

		// Right
		for i := 4*room_size + loc/room_size + 2; i < len(r); i++ {
			if r[i] != ' ' {
				// space is occupied
				break
			}

			new_room := r.swap(i, loc)

			// Travelling right
			curr_room_no := loc / room_size
			shifted := i - 4*room_size - 2

			dist := (int(math.Abs(float64(curr_room_no-shifted)))+1)*2 + loc%room_size
			if shifted == 4 {
				dist--
			}
			ret[new_room] = dist * travel_cost[target_room]

		}

	}

	// Pods in the hallway - 7 spots
	for i := room_size * 4; i < len(r); i++ {
		if r[i] == ' ' {
			continue
		}

		target_room := int(r[i] - '0')
		target_spot, ok := r.target_room_open(target_room)
		if !ok {
			continue
		}

		// Calculate the distance
		shifted := i - 4*room_size - 2

		dist := -1
		blocked := false
		if target_room <= shifted {
			// Is there an open path to the left ?
			for j := i - 1; j >= room_size*4+target_room+2; j-- {
				if r[j] != ' ' {
					blocked = true
					break
				}
			}

			// Going left
			dist = (shifted-target_room+1)*2 + target_spot
			if shifted == 4 {
				dist--
			}

		} else {

			// Is there an open path to the right ?
			for j := i + 1; j <= room_size*4+target_room+1; j++ {
				if r[j] != ' ' {
					blocked = true
					break
				}
			}
			dist = (target_room-shifted)*2 + target_spot
			if shifted == -2 {
				dist--
			}
		}

		if !blocked {
			new_room := r.swap(i, target_room*room_size+target_spot)
			ret[new_room] = dist * travel_cost[target_room]
		}

	}

	return ret

}

func (r rooms) target_room_open(room int) (int, bool) {
	room_size := r.room_size()
	target := byte(room + '0')
	open := -1

	for i := room_size - 1; i >= 0; i-- {
		spot := room_size*room + i
		if r[spot] != ' ' && r[spot] != target {
			return -1, false
		}

		if r[spot] == ' ' && open == -1 {
			open = spot % room_size
		}
	}

	return open, true
}

func (r rooms) swap(i int, j int) rooms {

	new := []rune(r)
	new[i], new[j] = new[j], new[i]
	return rooms(new)
}

func (r rooms) print() {

	room_size := r.room_size()
	o := 4 * room_size

	fmt.Println()
	fmt.Println("#############")
	fmt.Printf("#%v%v %v %v %v %v%v#\n", string(r[o]), string(r[o+1]), string(r[o+2]), string(r[o+3]), string(r[o+4]), string(r[o+5]), string(r[o+6]))
	fmt.Printf("###%v#%v#%v#%v###\n", string(r[0*room_size]), string(r[1*room_size]), string(r[2*room_size]), string(r[3*room_size]))

	for i := 1; i < room_size; i++ {
		fmt.Printf("  #%v#%v#%v#%v#  \n", string(r[0*room_size+i]), string(r[1*room_size+i]), string(r[2*room_size+i]), string(r[3*room_size+i]))
	}

	fmt.Println("  #########  ")
}

// -2 -1 0 1 2 3 4
// 89 A B C DE
//   0 2 4 6
//   1 3 5 7

// Small Rooms : 15
// Start   State : "CBAABDDC       "
// Desired State : "AABBCCDD       "

// spot % room_size + 1 = position in room
// Distance to other room = position + diff(position, target) * 2 + 1

// GH I J K LM
//   0 4 8 C
//   1 5 9 D
//   2 6 A E
//   3 7 B F

// Big Rooms : 23 characters
// Start State   : "CDDBACBABBADDACC       "
// Desired State : "AAAABBBBCCCCDDDD       "

func read_room(lines []string) rooms {

	mapping := map[rune]rune{
		'A': '0',
		'B': '1',
		'C': '2',
		'D': '3',
	}

	rms := make([][]rune, 4)

	for _, line := range lines {
		if line == "#############" || line == "#...........#" || line == "  #########" {
			continue
		}

		for i := range rms {
			rms[i] = append(rms[i], mapping[rune(line[3+i*2])])
		}
	}

	var sb strings.Builder

	for _, room := range rms {
		for _, r := range room {
			sb.WriteRune(r)
		}
	}

	sb.WriteString("       ")

	return rooms(sb.String())

}

// #############
// #...........#
// ###C#A#B#D###
//   #B#A#D#C#
//   #########
