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

	dict_matrix := make(map[int][]int)
	scanner := bufio.NewScanner(input)

	x_max := 0
	y_max := 0

	for scanner.Scan() {
		line := strings.Split(scanner.Text(), " ")
		id, _ := strconv.Atoi(line[0][1:])
		loc := strings.Split(line[2], ",")
		left, _ := strconv.Atoi(loc[0])
		top, _ := strconv.Atoi(loc[1][0 : len(loc[1])-1])
		//	fmt.Println(line, line[2], left, top)
		place := strings.Split(line[3], "x")
		width, _ := strconv.Atoi(place[0])
		height, _ := strconv.Atoi(place[1])
		//	fmt.Println(line, line[3], width, height)
		x := left + width
		y := top + height

		if x > x_max {
			x_max = x
		}
		if y > y_max {
			y_max = y
		}

		dict_matrix[id] = []int{left, top, width, height}
	}

	//	fmt.Println(dict_matrix)
	//	fmt.Println(x_max, y_max)
	matrix := make([][]uint8, y_max)
	for i := range matrix {
		matrix[i] = make([]uint8, x_max)
	}

	for _, val := range dict_matrix {
		w := val[0] + val[2]
		h := val[1] + val[3]
		for i := val[0]; i < w; i++ {
			for j := val[1]; j < h; j++ {
				matrix[i][j]++
			}

		}
	}

	counter := 0
	for _, i := range matrix {
		for _, j := range i {
			if j > 1 {
				counter++
			}
		}
	}

	fmt.Println(counter)
}
