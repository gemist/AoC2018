package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	input, _ := os.Open("input.txt")
	defer input.Close()

	scanner := bufio.NewScanner(input)

	sum := 0
	for scanner.Scan() {
		item, _ := strconv.ParseInt(scanner.Text(), 10, 32)
		sum += int(item)
	}

	fmt.Println(sum)
}
