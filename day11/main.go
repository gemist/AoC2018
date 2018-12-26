package main

import (
	"fmt"
)

func main() {
	GSN := 9424 //input parameter

	//	maxPower := 0
	//	maxX := 0
	//	maxY := 0
	var matrix [300][300]int
	for X := 1; X <= 300; X++ {
		for Y := 1; Y <= 300; Y++ {
			rackID := X + 10
			powLevel := rackID * Y
			out := powLevel + GSN
			out = out * rackID
			digit := (out / 100) % 10
			power := digit - 5
			//	fmt.Println(power)
			matrix[Y-1][X-1] = power
		}
	}
	/*	for i := range matrix {
			fmt.Println(matrix[i])
		}
	*/
	maxPower := 0
	var maxX int
	var maxY int
	for i := 0; i < 277; i++ {
		for j := 0; j < 277; j++ {
			power := 0
			for x := 0; x < 3; x++ {
				for y := 0; y < 3; y++ {
					power += matrix[j+y][i+x]
				}
			}
			if power > maxPower {
				maxPower = power
				maxX = i + 1
				maxY = j + 1
			}

		}
	}
	fmt.Println("part1", maxX, ",", maxY)

	//part2 - much slower
	maxPower = 0
	var targetSize int
	for size := 0; size < 300; size++ {
		for i := 0; i < 300-size; i++ {
			for j := 0; j < 300-size; j++ {

				power := 0
				for x := 0; x < size; x++ {
					for y := 0; y < size; y++ {
						power += matrix[j+y][i+x]
					}
				}
				if power > maxPower {
					maxPower = power
					maxX = i + 1
					maxY = j + 1
					targetSize = size
				}
			}
		}
	}
	fmt.Println("part2", maxX, ",", maxY, ",", targetSize)
}
