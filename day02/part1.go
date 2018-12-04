package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func getValues(arr []string) (two int, three int) {
	//Create a   dictionary of values for each element
	dict := make(map[string]int)
	for _, num := range arr {
		dict[num] = dict[num] + 1
	}

	two = 0
	three = 0
	for _, num := range dict {
		if num == 2 {
			two = 1
		}
		if num == 3 {
			three = 1
		}
	}
	//	fmt.Print(dict)
	return two, three
}

func main() {

	input, _ := os.Open("input.txt")
	defer input.Close()

	scanner := bufio.NewScanner(input)

	two_tot := 0
	three_tot := 0

	for scanner.Scan() {
		item := scanner.Text()
		table := strings.Split(item, "")

		two, three := getValues(table)
		two_tot += two
		three_tot += three
	}
	fmt.Println(two_tot * three_tot)
}
