package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	file, _ := os.Open("input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	pot_notes := make([]string, 0)
	no_pot_notes := make([]string, 0)
	var init_state string

	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, "initial state") {
			init_state = strings.TrimLeft(strings.Split(line, ":")[1], " ")
		} else if strings.Contains(line, "=>") {
			instruction := strings.Split(line, "=>")
			if strings.TrimLeft(instruction[1], " ") == "#" {
				pot_notes = append(pot_notes, strings.TrimRight(instruction[0], " "))
			} else {
				no_pot_notes = append(no_pot_notes, strings.TrimRight(instruction[0], " "))
			}

		}

	}

	//part 1
	old_state := strings.Join([]string{".", init_state, "."}, "")
	shift := 1 //added . at the beginning

	for generation := 0; generation < 20; generation++ {
		old_state, shift = nextGen(old_state, shift, pot_notes, no_pot_notes)
		//		fmt.Println(old_state)
	}

	fmt.Println("part1:", sum_pot(old_state, shift))

	//part 2
	num := 50000000000
	old_state = strings.Join([]string{".", init_state, "."}, "")
	old_shift := 1 //reinitialize shift
	var new_shift int
	var new_state string
	generation := 0
	var difference int

	for {
		generation++
		new_state, new_shift = nextGen(old_state, old_shift, pot_notes, no_pot_notes)

		difference = sum_pot(new_state, new_shift) - sum_pot(old_state, old_shift)
		if strings.Compare(old_state, new_state) == 0 {
			break
		}
		old_state = new_state
		old_shift = new_shift

	}

	fmt.Println("part2:", sum_pot(new_state, new_shift)+difference*(num-generation))
}

func nextGen(old_state string, shift int, pot_notes []string, no_pot_notes []string) (string, int) {
	new_list := make([]string, 0)
	var length int = 5
	var substring string
	for i := 0; i < len(old_state); i++ {
		subslice := make([]string, 0)
		if i == 0 {
			for j := 0; j < length; j++ {
				if j < 2 {
					subslice = append(subslice, ".")
				} else {
					subslice = append(subslice, string(old_state[i+j-2]))
				}
			}
		} else if i == 1 {
			for j := 0; j < length; j++ {
				if j < 1 {
					subslice = append(subslice, ".")
				} else {
					subslice = append(subslice, string(old_state[i+j-2]))
				}
			}

		} else if i == len(old_state)-2 {
			for j := 0; j < length; j++ {
				if j > 3 {
					subslice = append(subslice, ".")
				} else {
					subslice = append(subslice, string(old_state[i+j-2]))
				}
			}
		} else if i == len(old_state)-1 {
			for j := 0; j < length; j++ {
				if j > 2 {
					subslice = append(subslice, ".")
				} else {
					subslice = append(subslice, string(old_state[i+j-2]))
				}
			}

		} else {
			for j := 0; j < length; j++ {
				subslice = append(subslice, string(old_state[i+j-2]))
			}
		}
		substring = strings.Join(subslice, "")

		match := false

		for _, val := range pot_notes {
			if strings.Compare(substring, val) == 0 {
				new_list = append(new_list, "#")
				match = true
				break
			}
		}

		if match == false {
			for _, val := range no_pot_notes {

				if strings.Compare(substring, val) == 0 {
					new_list = append(new_list, ".")
					match = true
					break
				}

			}
		}
		// just in case
		if match == false {
			new_list = append(new_list, ".")
		}

	}

	for i, val := range new_list {
		if val == "#" {
			shift -= i
			break
		}
	}
	new := strings.Join([]string{".", strings.Trim(strings.Join(new_list, ""), "."), "."}, "")
	shift++

	return new, shift

}

func sum_pot(state string, shift int) int {
	count := 0
	for i := 0; i < len(state); i++ {
		if string(state[i]) == "#" {
			count += i - shift
		}
	}
	return count
}
