package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	input, _ := os.Open("input.txt")
	defer input.Close()

	var numPlayers, lastmarble int
	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), " ")
		numPlayers, _ = strconv.Atoi(line[0])
		lastmarble, _ = strconv.Atoi(line[6])
	}

	marbleList := make([]int, 1)
	position := 0
	//lastmarble = lastmarble * 100

	players := make([]int, numPlayers)

	for i := 1; i <= lastmarble; i++ {

		player_index := (i - 1) % numPlayers

		//handle first marble
		if len(marbleList) < 2 {
			marbleList = append(marbleList, i)
			position++
			continue

		}

		if i%23 == 0 {
			position -= 7

			if position < 0 {
				position = len(marbleList) + position
			}

			players[player_index] += marbleList[position] + i
			marbleList = append(marbleList[:position], marbleList[position+1:]...)

		} else {

			position += 2
			if position > len(marbleList) {
				position = 1
			}
			temp := make([]int, 0)
			temp = append(temp, i)
			temp = append(temp, marbleList[position:]...)
			marbleList = append(marbleList[:position], temp...)

		}
	}

	max_score := 0
	for i := 0; i < len(players); i++ {
		if players[i] > max_score {
			max_score = players[i]
		}
	}
	fmt.Println(max_score)

}
