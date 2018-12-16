package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

type worker struct {
	currentJob string
	timeLeft   int
	working    bool
}

func contains(slice []string, ch string) bool {
	for _, val := range slice {
		if val == ch {
			return true
		}
	}
	return false
}

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

	//	fmt.Println(hashtable)
	//	fmt.Println("-----------------------")

	//set worktime
	worksteps := make(map[string]int)
	var ch byte
	for ch = 65; ch <= 90; ch++ {
		worksteps[string(ch)] = int(ch - 4) //int(ch - 64) //int(ch - 4)

	}

	//workers
	num_workers := 5
	var workers []*worker //slice of workers
	// Create a slice of workers
	for i := 0; i < num_workers; i++ {
		workers = append(workers, &worker{
			working: false,
		})
	}

	working_que := make([]string, 0)
	occupied_que := make([]string, 0)
	temp := make([]string, 0)
	result := make([]string, 0)

	var element string

	duration := 0

	//	for step := 0; step < 3; step++ {
	for len(hashtable) > 0 {
		for k, v := range hashtable {
			for count, i := range v {
				for _, element := range temp {
					if i == element {
						hashtable[k] = append(hashtable[k][:count], hashtable[k][count+1:]...)
					}
				}
			}

			if len(hashtable[k]) == 0 && contains(working_que, k) == false {
				working_que = append(working_que, k)
			}

		}
		temp = nil
		//	sort.Sort(sort.Reverse(sort.StringSlice(working_que)))
		sort.Sort(sort.StringSlice(working_que))
		completed := false

		for completed == false {
			//clean working que
			for _, ival := range occupied_que {
				for j, jval := range working_que {
					if jval == ival {
						working_que = append(working_que[:j], working_que[j+1:]...)
					}
				}
			}

			for _, val := range working_que {

				for k, _ := range workers {

					if workers[k].working == false { //&& contains(occupied_que, val) == false {
						workers[k].currentJob = val
						workers[k].timeLeft = worksteps[val]
						workers[k].working = true
						occupied_que = append(occupied_que, val)
						break
					}

				}

			}

			for k, _ := range workers {

				if workers[k].working == true {
					workers[k].timeLeft--
					if workers[k].timeLeft == 0 {
						temp = append(temp, workers[k].currentJob)
						workers[k].working = false
						for i, ival := range occupied_que {
							if workers[k].currentJob == ival {
								occupied_que = append(occupied_que[:i], occupied_que[i+1:]...)
								break
							}
						}

						completed = true

					}

				}
			}
			duration++

		}
		//	fmt.Println(hashtable)
		//	fmt.Println("working", working_que)
		//	fmt.Println("occupied", occupied_que)
		//	fmt.Println("temp", temp)
		for _, element := range temp {
			//		fmt.Println(element)
			delete(hashtable, element)
		}
		//	fmt.Println(hashtable)
		//	fmt.Println(duration)

		result = append(result, element)
		//	fmt.Println("........")
	}
	fmt.Println(duration)
}
