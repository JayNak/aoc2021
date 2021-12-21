package aoc

import (
	"regexp"
	"strconv"

	"github.com/jaynak/aoc2021/pkg/util"
)

func Day21(path string) (int, int) {

	lines := util.ReadToStrings(path)
	r := regexp.MustCompile("Player [1-2] starting position: ([0-9]+)")

	start_positions := []int{}

	for _, line := range lines {
		m := r.FindAllStringSubmatch(line, -1)
		n, err := strconv.Atoi(m[0][1])
		if err != nil {
			panic(err)
		}
		start_positions = append(start_positions, n)
	}

	val := simulate_game(start_positions[0], start_positions[1], 100, 1000)

	x, y := simulate_universes(0, 0, start_positions[0], start_positions[1])

	ret := x
	if y > x {
		ret = y
	}

	return val, int(ret)
}

func simulate_universes(curr_score int, other_score int, curr_pos int, other_pos int) (int64, int64) {
	if other_score >= 21 {
		return 0, 1
	}

	possibilities := map[int]int{3: 1, 4: 3, 5: 6, 6: 7, 7: 6, 8: 3, 9: 1}

	var wins, losses int64

	for k, v := range possibilities {
		new_pos := (curr_pos + k) % 10
		if new_pos == 0 {
			new_pos = 10
		}
		new_losses, new_wins := simulate_universes(other_score, curr_score+new_pos, other_pos, new_pos)
		wins += new_wins * int64(v)
		losses += new_losses * int64(v)
	}

	return wins, losses
}

func simulate_game(p1start int, p2start int, die int, target int) int {

	pos := []int{p1start, p2start}
	score := []int{0, 0}
	die_val := die
	turn := 0
	player := 1

	for score[0] < target && score[1] < target {

		player = turn % 2

		this_roll := 0
		for i := 0; i < 3; i++ {
			if die_val == die {
				die_val = 1
			} else {
				die_val++
			}
			this_roll += die_val
		}

		turn++

		pos[player] = pos[player] + this_roll%10
		if pos[player] > 10 {
			pos[player] -= 10
		}

		score[player] += pos[player]
	}

	for _, s := range score {
		if s < target {
			return s * (turn * 3)
		}
	}

	return 0
}
