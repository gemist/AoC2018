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

	//create grid
	grid := make([][]int, max_grid)
	for i := range grid {
		grid[i] = make([]int, max_grid)
	}

	ref_value := 10000
	num_points := 0
	for j := 0; j < max_grid; j++ {
		for i := 0; i < max_grid; i++ {
			distance := 0
			for _, p_val := range points {
				distance += manhattanDistance(i, j, point{x: int(p_val.x), y: int(p_val.y)})
			}
			grid[i][j] = distance
			if distance < ref_value {
				num_points++
			}
		}

	}

	//	fmt.Println(grid)
	fmt.Println(num_points)

}
