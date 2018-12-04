package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	input, _ := os.Open("input.txt")
	defer input.Close()

	scanner := bufio.NewScanner(input)

	input_slice := make([]string, 0, 100)

	//fill the slice
	for scanner.Scan() {
		item := scanner.Text()
		input_slice = append(input_slice, item)
	}
	//fmt.Println(input_slice[1])

	//do the action

	for i := 0; i < len(input_slice)-1; i++ {
		ref_table1 := strings.Split(input_slice[i], "")

		for j := 1; j < len(input_slice); j++ {
			ref_table2 := strings.Split(input_slice[j], "")
			common := make([]string, 0)
			for k := 0; k < len(ref_table1); k++ {
				if ref_table1[k] == ref_table2[k] {
					common = append(common, ref_table1[k])
				}
			}
			if len(common) == len(ref_table1)-1 {
				fmt.Println("Found: ", strings.Join(common, ""))
				return
			}

		}

	}
	fmt.Println("Not Found")
}
