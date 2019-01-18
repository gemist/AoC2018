package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {

	//input
	numRecipe := 704321

	strNum := strconv.Itoa(numRecipe)

	subslice := make([]int, 0)
	for _, val := range strNum {
		e, _ := strconv.Atoi(string(val))
		subslice = append(subslice, e)
	}
	//	fmt.Println(subslice)
	//substract 1 starting from 0
	numRecipe += -1
	scoreboard := make([]int, 2)
	firstElf := 0
	secondElf := 1
	scoreboard[firstElf] = 3
	scoreboard[secondElf] = 7

	counter := 0
	part1 := false
	part2 := false

	for {
		sum := scoreboard[firstElf] + scoreboard[secondElf]
		//if it is greater/equal than 10, can not be greater than 100
		if sum > 9 {
			scoreboard = append(scoreboard, sum/10)
			sum -= (sum / 10) * 10
		}
		scoreboard = append(scoreboard, sum)

		//handle elf positions
		first := firstElf + scoreboard[firstElf] + 1
		firstElf = first - (first)/len(scoreboard)*len(scoreboard)
		second := secondElf + scoreboard[secondElf] + 1
		secondElf = second - (second)/len(scoreboard)*len(scoreboard)

		//part 1
		if part1 == false && numRecipe+10 < len(scoreboard) {
			str := make([]string, 0)
			for i := numRecipe + 1; i < numRecipe+11; i++ {
				str = append(str, strconv.Itoa(scoreboard[i]))
			}
			fmt.Println("part 1: ", strings.Join(str, ""))
			part1 = true
		}

		//part 2
		if part2 == false && len(scoreboard) > len(subslice) {
			if compare(scoreboard[counter:counter+len(subslice)], subslice) {
				fmt.Println("part2:", counter)
				part2 = true
			}
			counter++
		}
		if part1 && part2 {
			break
		}
		//	counter++

	}

}

// compare slices
// A nil argument is equivalent to an empty slice.
func compare(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}
