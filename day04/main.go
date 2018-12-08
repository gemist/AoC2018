package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

// Event stores the timestamp and a slice of strings containing details of the event.
type event struct {
	datetime time.Time
	event    []string
}

type sleeplog struct {
	sleepminutes int
	timeFrame    []int
}

type timeSorted []event

func (p timeSorted) Len() int {
	return len(p)
}

func (p timeSorted) Less(i, j int) bool {
	return p[i].datetime.Before(p[j].datetime)
}

func (p timeSorted) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func find(s []string, x string) bool {
	for _, n := range s {

		if x == n {
			return true
		}
	}
	return false
}

func main() {

	var id string
	var foundid string
	var start time.Time
	var end time.Time
	var index int
	var max int

	var events []event
	input, _ := os.Open("input.txt")
	defer input.Close()

	scanner := bufio.NewScanner(input)

	for scanner.Scan() {
		var datetime strings.Builder
		line := strings.Split(scanner.Text(), " ")
		datetime.WriteString(line[0][1:])
		datetime.WriteString(" ")
		datetime.WriteString(line[1][:5])
		t, _ := time.Parse("2006-01-02 15:04", datetime.String())

		events = append(events, event{
			datetime: t,
			event:    line[2:],
		})
	}
	sort.Sort(timeSorted(events))

	guard := make(map[string]*sleeplog)

	for _, i := range events {
		if find(i.event, "Guard") {
			id = i.event[1][1:]
			_, exists := guard[id]
			if !exists {
				guard[id] = &sleeplog{
					timeFrame:    make([]int, 60),
					sleepminutes: 0,
				}
			}
		}

		if find(i.event, "falls") {
			start = i.datetime
		}

		if find(i.event, "wakes") {
			end = i.datetime
			e := end.Sub(start)
			s_ind := int(time.Duration(start.Minute()))
			d_add := int(e / time.Minute)
			for j := s_ind; j < s_ind+d_add; j++ {
				guard[id].timeFrame[j]++
				guard[id].sleepminutes++
			}
		}
	}
	/*	for ind, _ := range guard {
			fmt.Println(ind, guard[ind], guard[ind].sleepminutes)
		}
	*/
	longest_sleep := 0

	for k, v := range guard {
		if v.sleepminutes > longest_sleep {
			longest_sleep = v.sleepminutes
			foundid = k
		}
	}

	for i, v := range guard[foundid].timeFrame {
		if v > max {
			max = v
			index = i
		}
	}

	guard_num, _ := strconv.Atoi(foundid)

	fmt.Println("part1: ", guard_num*index)

	//part 2
	max_freq := 0
	for k, guard_val := range guard {
		for i, timeSlice_val := range guard_val.timeFrame {
			if timeSlice_val > max_freq {
				max_freq = timeSlice_val
				index = i
				foundid = k
			}
		}
	}
	guard_num, _ = strconv.Atoi(foundid)

	fmt.Println("part2: ", guard_num*index)
}
