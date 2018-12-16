package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	hashtable := make(map[string][]string)

	for scanner.Scan() {
		line := strings.Split(scanner.Text(), " ")
		hashtable[line[7]] = append(hashtable[line[7]], line[1])
		if _, exist := hashtable[line[1]]; exist == false {
			hashtable[line[1]] = make([]string, 0)
		}
	}

	temp := make([]string, 0)
	result := make([]string, 0)

	var element string

	for len(hashtable) > 0 {
		for k, v := range hashtable {
			for count, i := range v {
				if i == element {
					hashtable[k] = append(hashtable[k][:count], hashtable[k][count+1:]...)
				}
			}
			if len(hashtable[k]) == 0 {
				temp = append(temp, k)
			}
		}
		sort.Strings(temp)
		element = temp[0]
		delete(hashtable, element)
		result = append(result, element)
		temp = nil
	}
	fmt.Println(strings.Join(result, ""))
}
