package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type point struct {
	x  int
	y  int
	vx int
	vy int
}

func Abs(x int) int {
	if x < 0 {
		x = -x
	}
	return x
}

func move(points *[]point) {
	for i, _ := range *points {
		(*points)[i].x += (*points)[i].vx
		(*points)[i].y += (*points)[i].vy
	}
}

func plot(points *[]point) {
	minX, maxX, minY, maxY := (*points)[0].x, (*points)[0].x, (*points)[0].y, (*points)[0].y
	for i, _ := range *points {
		if (*points)[i].x > maxX {
			maxX = (*points)[i].x
		}
		if (*points)[i].x < minX {
			minX = (*points)[i].x
		}
		if (*points)[i].y > maxY {
			maxY = (*points)[i].y
		}
		if (*points)[i].y < minY {
			minY = (*points)[i].y
		}
	}
	lenX := int(maxX - minX + 1)
	lenY := int(maxY - minY + 1)

	array := make([][]string, lenY)
	for i := range array {
		array[i] = make([]string, lenX)
	}

	for i := 0; i < lenY; i++ {
		for j := 0; j < lenX; j++ {
			array[i][j] = "."
		}
	}

	for i, _ := range *points {
		array[(*points)[i].y-minY][(*points)[i].x-minX] = "#"
	}

	for i := 0; i < lenY; i++ {

		fmt.Println(array[i])
	}
}

func messageFound(points *[]point) bool {

	count := 0
	found := false
	var distX, distY int

	for i, _ := range *points {
		for j, _ := range *points {
			if i != j {
				distX = Abs((*points)[i].x - (*points)[j].x)
				distY = Abs((*points)[i].y - (*points)[j].y)
			}
			if distX <= 1 && distY <= 1 {
				count++
				break
			}
		}

	}
	if count == len(*points) {
		found = true
	}

	return found
}

func main() {
	file, _ := os.Open("input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var points []point

	for scanner.Scan() {

		line := strings.Split(scanner.Text(), "=")
		//coordinate
		coor := strings.Split(line[1], ",")
		x, _ := strconv.Atoi(strings.TrimLeft(coor[0], "< "))
		y, _ := strconv.Atoi(strings.Split(strings.TrimLeft(coor[1], " "), ">")[0])
		//velocities
		vel := strings.Split(strings.Replace(line[2], " ", "", -1), ",")
		vx, _ := strconv.Atoi(strings.TrimLeft(vel[0], "<"))
		vy, _ := strconv.Atoi(strings.TrimRight(vel[1], "> "))

		points = append(points, point{
			x:  x,
			y:  y,
			vx: vx,
			vy: vy,
		})

	}

	counter := 0
	for {
		counter++

		move(&points)
		found := messageFound(&points)

		if found {
			break
		}
	}
	plot(&points)
	fmt.Println("number of seconds", counter)
}
