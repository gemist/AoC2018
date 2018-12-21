package main

import (
	"bufio"
	"container/ring"
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
	players := make([]int, numPlayers)

	marbleValue := 0
	marbleList := ring.New(1)

	marbleList.Value = marbleValue

	lastmarble *= 100

	for {
		marbleValue++

		if marbleValue > lastmarble {
			break
		}

		if marbleValue%23 == 0 {
			player_index := (marbleValue - 1) % numPlayers
			players[player_index] += marbleValue

			marbleList = marbleList.Move(-7)
			players[player_index] += marbleList.Value.(int)
			marbleList = marbleList.Prev()
			marbleList.Unlink(1)
			marbleList = marbleList.Next()
			continue
		}
		marbleList = marbleList.Next()

		r := ring.New(1)
		r.Value = marbleValue

		marbleList = marbleList.Link(r)
		marbleList = marbleList.Prev()
	}

	// print ring
	//	marbleList.Do(func(x interface{}) {
	//		fmt.Println(x)
	//	})

	max_score := 0
	for i := 0; i < len(players); i++ {
		if players[i] > max_score {
			max_score = players[i]
		}
	}

	fmt.Println(max_score)

}
