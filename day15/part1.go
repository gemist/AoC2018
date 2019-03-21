package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type player struct {
	name string
	x    int
	y    int
}

func inrange(player player, players []player, grid [][]string) (inrange [][]int) {
	inrange = make([][]int, 0)

	vector := [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
	for i, _ := range players {
		if player != players[i] && player.name != players[i].name {
			for _, vecval := range vector {
				if grid[players[i].y+vecval[1]][players[i].x+vecval[0]] == "." {
					inrange = append(inrange, []int{players[i].x + vecval[0], players[i].y + vecval[1]})
				}
			}
		}
	}
	return inrange
}

/*
func (f field) bfs(start point, target rune) map[point]bool {
	out := map[point]bool{}
	visited := map[point]bool{start: true}
	next := []point{start}
	cur := []point{} // to avoid new slices - just swap them

	for step := 0; len(next) > 0 && len(out) == 0; step++ { //steps
		cur, next = next, cur[:0]
		for _, p := range cur {
			for d := 0; d <= 3; d++ {
				x := p.x + (1-d)%2
				y := p.y + (2-d)%2
				t := point{y, x}
				if visited[t] {
					continue
				}
				switch f[y][x] {
				case target:
					out[t] = true
				case '.':
					next = append(next, t)
				}
				visited[t] = true
			}
		}
	}
	return out
}
*/

func main() {

	file, _ := os.Open("input1.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)

	var players []player
	grid := make([][]string, 0)
	line_num := 0
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), "")
		gridline := make([]string, 0)
		fmt.Println(line)
		for i, val := range line {
			player_bool := false
			//	var x, y int
			var player_name string

			if val == "G" {
				val = "."
				player_name = "G"
				player_bool = true
			} else if val == "E" {
				val = "."
				player_name = "E"
				player_bool = true
			}

			if player_bool == true {
				players = append(players, player{
					name: string(player_name),
					x:    int(i),
					y:    int(line_num),
				})
			}
			gridline = append(gridline, val)
		}
		grid = append(grid, gridline)
		line_num++
	}
	for i, _ := range grid {

		fmt.Println(grid[i])
	}

	fmt.Println(inrange(players[0], players, grid))
	fmt.Println(grid[1][0])
	fmt.Println(players)
}
