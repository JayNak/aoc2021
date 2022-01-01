package aoc

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/jaynak/aoc2021/pkg/util"
)

func Day24(path string) (int, int) {

	lines := util.ReadToStrings(path)

	_, mods := read_instruction_set(lines)

	largest := find_valid_model_4(mods, true)
	smallest := find_valid_model_4(mods, false)

	// num, ok := find_valid_model_3(mods, 0)
	// num, ok := find_valid_model_2(inst, []int{0, 0, 0, 0})

	// max_model := find_valid_model(inst, []int{})
	// input := []int{1, 3, 5, 7, 9, 2, 4, 6, 8, 9, 9, 9, 9, 9}
	// reg := apply_monad_instructions(inst, input)

	// fmt.Println(reg)

	return largest, smallest
}

func find_valid_model_4(mods [][]int, largest bool) int {

	// If looking for the largest, look 1->9 to account for overwrites
	w_vals := []int{9, 8, 7, 6, 5, 4, 3, 2, 1}
	if largest {
		w_vals = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	}

	results := make(map[int][]int)

	desired_z := []int{0}
	for i := len(mods) - 1; i >= 0; i-- {
		new_desired_z := []int{}
		for _, w := range w_vals {
			for _, z := range desired_z {
				temp, ok := results[z]
				if !ok {
					temp = []int{}

				}
				new_zs := get_possible_z_in(mods[i], z, w)

				for _, new_z := range new_zs {
					results[new_z] = append([]int{w}, temp...)
					new_desired_z = append(new_desired_z, new_z)
				}
			}
		}
		desired_z = new_desired_z
	}

	str := ""
	for _, v := range results[0] {
		str += strconv.Itoa(v)
	}

	model, _ := strconv.Atoi(str)

	return model
}

func get_possible_z_in(mod []int, z_out int, w int) []int {

	ret := []int{}

	// mod 0=z, 1=x, 2=y

	// Reverse the apply_mod function
	try := z_out - w - mod[2]
	if try%26 == 0 {
		ret = append(ret, try/26*mod[0])
	}
	if w-mod[1] >= 0 && w-mod[1] < 26 {
		ret = append(ret, w-mod[1]+z_out*mod[0])
	}

	return ret
}

func find_valid_model_3(mods [][]int, z_in int) ([]int, bool) {

	for num := 9; num > 0; num-- {

		z_out := apply_mod(mods[0], num, z_in)
		// locals := apply_monad_instructions(instructions[0], registers, num)

		if len(mods) == 1 {
			if z_out == 0 {
				// This is a valid number
				return []int{num}, true
			}
		} else {
			nums, ok := find_valid_model_3(mods[1:], z_out)
			if ok {
				return append([]int{num}, nums...), true
			}
		}
	}

	// Fell through
	return []int{}, false
}

func apply_mod(mod []int, w_in int, z_in int) int {

	// mod 0=z, 1=x, 2=y
	if w_in == z_in%26+mod[1] {
		return z_in / mod[0]
	} else {
		return (z_in/mod[0])*26 + w_in + mod[2]
	}
}

type monad_inst struct {
	op  int
	reg []int
	val int
}

func find_valid_model_2(instructions [][]*monad_inst, registers []int) ([]int, bool) {

	for num := 9; num > 0; num-- {
		fmt.Printf("%v ", num)
		locals := apply_monad_instructions(instructions[0], registers, num)

		if len(instructions) == 1 {
			fmt.Printf("\n")
			if locals[3] == 0 {
				return []int{num}, true
			} else {
				return []int{}, false
			}
		} else {
			nums, ok := find_valid_model_2(instructions[1:], locals)
			if ok {
				return append([]int{num}, nums...), true
			}
		}
	}

	// Fell through
	return []int{}, false
}

// func find_valid_model(instructions []*monad_inst, prev []int) int {

// 	if len(prev) == 14 {
// 		ret := apply_monad_instructions(instructions, prev)

// 		if ret[3] == 0 {
// 			// This is a valid model number
// 			str := ""
// 			for _, i := range prev {
// 				str += strconv.Itoa(i)
// 			}

// 			val, _ := strconv.Atoi(str)

// 			return val
// 		}

// 		return 0
// 	}

// 	// Not at the end yet
// 	nums := prev
// 	nums = append(nums, 0)
// 	for i := 9; i > 0; i-- {
// 		nums[len(nums)-1] = i
// 		ret := find_valid_model(instructions, nums)
// 		if ret != 0 {
// 			return ret
// 		}
// 	}

// 	// Fell through
// 	return 0
// }

func apply_monad_instructions(instructions []*monad_inst, registers []int, try int) []int {

	temp := make([]int, 4)
	copy(temp, registers)

	ptr := 0

	for _, inst := range instructions {
		num := inst.val

		if inst.op == 0 {
			num = try
			ptr++
		} else {
			if len(inst.reg) == 2 {
				num = temp[inst.reg[1]]
			}
		}

		dest := inst.reg[0]

		// Do the op
		switch inst.op {
		case 0:
			temp[dest] = num
		case 1:
			temp[dest] += num
		case 2:
			temp[dest] *= num
		case 3:
			temp[dest] /= num
		case 4:
			temp[dest] = temp[dest] % num
		case 5:
			if temp[dest] == num {
				temp[dest] = 1
			} else {
				temp[dest] = 0
			}
		}
	}

	return temp
}

func read_instruction_set(lines []string) ([][]*monad_inst, [][]int) {

	r := regexp.MustCompile("([a-z]+) ([w-z])( ([0-9a-z-]+))?")
	inst_map := map[string]int{
		"imp": 0, "add": 1, "mul": 2, "div": 3, "mod": 4, "eql": 5,
	}
	register_map := map[string]int{
		"w": 0, "x": 1, "y": 2, "z": 3,
	}

	ret := [][]*monad_inst{}
	instructions := []*monad_inst{}

	count := 0
	mod := make([][]int, 14)
	mod[count] = make([]int, 3)

	for _, line := range lines {
		m := r.FindAllStringSubmatch(line, -1)

		inst := &monad_inst{
			op:  inst_map[m[0][1]],
			reg: []int{register_map[m[0][2]]},
		}

		if inst.op != 0 {

			// Process the second variable
			switch m[0][4] {
			case "w", "x", "y", "z":
				inst.reg = append(inst.reg, register_map[m[0][4]])
			default:
				num, _ := strconv.Atoi(m[0][4])
				inst.val = num

				if m[0][2] == "z" && m[0][1] == "div" {
					mod[count][0] = num
				} else if m[0][1] == "add" {
					mod[count][register_map[m[0][2]]] = num
				}
			}

		} else {
			if len(instructions) > 0 {
				ret = append(ret, instructions)
				instructions = []*monad_inst{}

				count++
				mod[count] = make([]int, 3)
			}
		}

		instructions = append(instructions, inst)
	}

	ret = append(ret, instructions)

	return ret, mod
}
