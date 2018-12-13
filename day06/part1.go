package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type point struct {
	x int
	y int
}

// Abs returns the absolute value of x.
func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func manhattanDistance(x int, y int, p point) int {
	return Abs(x-p.x) + Abs(y-p.y)
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Print(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	max_grid := 0 //max x, y

	var points []point

	for scanner.Scan() {
		line := strings.Split(scanner.Text(), ",")
		x, _ := strconv.Atoi(line[0])
		if x > max_grid-1 {
			max_grid = x + 1
		}
		y, _ := strconv.Atoi(strings.TrimSpace(line[1]))
		if y > max_grid-1 {
			max_grid = y + 1
		}
		points = append(points, point{
			x: int(x),
			y: int(y),
		})

	}

	//	fmt.Println(points[2].y)

	//create grid
	grid := make([][]string, max_grid)
	for i := range grid {
		grid[i] = make([]string, max_grid)
	}
	//fill grid with points
	for _, p_val := range points {
		grid[p_val.x][p_val.y] = "*"

	}

	for j := 0; j < max_grid; j++ {
		for i := 0; i < max_grid; i++ {
			if grid[i][j] != "*" {

				counter := 0
				min_distance := 2 * max_grid // max distance = 2*max_grid
				s := "."
				for p, p_val := range points {

					distance := manhattanDistance(i, j, point{x: int(p_val.x), y: int(p_val.y)})
					if distance == min_distance {
						s = "."
						counter++
					}

					if distance < min_distance {
						min_distance = distance
						s = string(strconv.Itoa(p))
						counter = 0
					}

					grid[i][j] = s

				}
			}
		}
	}

	//	fmt.Println(grid)
	//calculate largest area
	//part1
	largestArea := 0
	for p, _ := range points {
		infinite := false
		area := 0
		for j := 0; j < max_grid; j++ {
			for i := 0; i < max_grid; i++ {
				if strconv.Itoa(p) == grid[i][j] {
					if i == 0 || j == 0 || i == max_grid-1 || j == max_grid-1 {
						infinite = true
					} else {
						area++
					}
				}
				if infinite {
					break
				}
			}
			if infinite {
				break
			}

		}

		if area > largestArea && infinite == false {
			largestArea = area

		}
	}

	fmt.Println(largestArea + 1) // +1 include also the point

}
