package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// Contains tells whether slice/array a contains element x.
func Contains(a []int, x int) bool {
	for _, n := range a {
		if x == n {
			return true
		}
	}
	return false
}

func main() {
	input, _ := os.Open("input.txt")
	defer input.Close()

	scanner := bufio.NewScanner(input)

	input_slice := make([]int, 0, 100)

	//fill the slice
	for scanner.Scan() {
		item, _ := strconv.ParseInt(scanner.Text(), 10, 32)
		input_slice = append(input_slice, int(item))
	}

	//fmt.Println(input_slice)

	frequency := make([]int, 1, 100)
	frequency[0] = 0
	var i int = 0
	sum := 0
	//infinite loop
	for {
		sum += input_slice[i]
		if Contains(frequency, sum) == false {
			frequency = append(frequency, sum)
		} else {
			fmt.Println(sum)
			return
		}
		i = (i + 1) % len(input_slice)
	}
}
