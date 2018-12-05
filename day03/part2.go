package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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

	unique_ID := make([]int, 0)
	for k, val := range dict_matrix {
		var sum uint8 = 0
		w := val[0] + val[2]
		h := val[1] + val[3]
		for i := val[0]; i < w; i++ {
			for j := val[1]; j < h; j++ {
				sum = sum + matrix[i][j] - 1
			}

		}
		if sum == 0 {
			unique_ID = append(unique_ID, k)

		}
	}
	//sort slice
	sort.Slice(unique_ID, func(i, j int) bool { return unique_ID[i] < unique_ID[j] })

	fmt.Println(unique_ID)
}
