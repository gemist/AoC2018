package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

type cart struct {
	x     int
	y     int
	vx    int
	vy    int
	state string
}

func collision(carts []cart, cart_val cart, cart_ind int) (bool, int) {
	cool := false
	var index int
	//point := make([]int, 0)
	for i, _ := range carts {
		if i != cart_ind {
			if carts[i].x == cart_val.x && carts[i].y == cart_val.y {
				index = i
				cool = true
			}

		}
		if cool {
			break
		}

	}
	return cool, index
}

func move(carts *[]cart, grid [][]string) {
	for i, _ := range *carts {
		(*carts)[i].x += (*carts)[i].vx
		(*carts)[i].y += (*carts)[i].vy
		if grid[(*carts)[i].y][(*carts)[i].x] == "+" {
			if (*carts)[i].state == "l" {
				if (*carts)[i].vx == 0 {
					(*carts)[i].vx = (*carts)[i].vy
					(*carts)[i].vy = 0
				} else {
					(*carts)[i].vy = -(*carts)[i].vx
					(*carts)[i].vx = 0
				}
				(*carts)[i].state = "s"

			} else if (*carts)[i].state == "s" {
				(*carts)[i].state = "r"

			} else if (*carts)[i].state == "r" {
				if (*carts)[i].vx == 0 {
					(*carts)[i].vx = -(*carts)[i].vy
					(*carts)[i].vy = 0
				} else {
					(*carts)[i].vy = (*carts)[i].vx
					(*carts)[i].vx = 0
				}
				(*carts)[i].state = "l"

			}
		} else if grid[(*carts)[i].y][(*carts)[i].x] == "\\" {
			temp := (*carts)[i].vx
			(*carts)[i].vx = (*carts)[i].vy
			(*carts)[i].vy = temp
		} else if grid[(*carts)[i].y][(*carts)[i].x] == "/" {
			temp := (*carts)[i].vx
			(*carts)[i].vx = -(*carts)[i].vy
			(*carts)[i].vy = -temp

		}
	}
}

func main() {
	file, _ := os.Open("input.txt")
	defer file.Close()

	var carts []cart
	grid := make([][]string, 0)

	scanner := bufio.NewScanner(file)
	line_num := 0
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), "")
		gridline := make([]string, 0)
		for i, val := range line {
			cart_bool := false
			var vx, vy int
			if val == ">" {
				vx = 1
				vy = 0
				val = "-"
				cart_bool = true
			} else if val == "<" {
				vx = -1
				vy = 0
				val = "-"
				cart_bool = true
			} else if val == "v" {
				vx = 0
				vy = 1
				val = "|"
				cart_bool = true
			} else if val == "^" {
				vx = 0
				vy = -1
				val = "|"
				cart_bool = true
			}
			if cart_bool == true {
				carts = append(carts, cart{
					x:     int(i),
					y:     int(line_num),
					vx:    int(vx),
					vy:    int(vy),
					state: "l",
				})
			}
			gridline = append(gridline, val)

		}
		grid = append(grid, gridline)
		line_num++
	}

	//fmt.Println(carts)
	//fmt.Println("-----------")
	/*for i, _ := range grid {
		fmt.Println(grid[i])
	}*/

	tick := 0

	first := true

	for {
		tick++

		sort.Slice(carts, func(i, j int) bool {
			if carts[i].y < carts[j].y {
				return true
			} else if carts[i].y > carts[j].y {
				return false
			}
			return carts[i].x < carts[j].x
		})

		old_carts := make([]cart, 0)
		old_carts = append(old_carts, carts...)

		move(&carts, grid)

		indexes := make([]int, 0)

		for ind, cart_val := range carts {
			//check collision with already moved carts
			cool, _ := collision(carts, cart_val, ind)
			if cool {
				indexes = append(indexes, ind)
			}
			//check collision with not yet moved carts
			old_cool, index := collision(old_carts[ind:], cart_val, 0)
			if old_cool {
				indexes = append(indexes, ind)
				indexes = append(indexes, ind+index)

			}

		}

		//delete collided carts
		if len(indexes) > 0 {
			new_carts := make([]cart, 0)
			for ind, cart_val := range carts {
				ind_bool := true
				for _, val := range indexes {
					if ind == val {
						if first {
							fmt.Println("Part1: ", cart_val.x, ",", cart_val.y)
							first = false
						}

						ind_bool = false
					}
				}
				if ind_bool {
					new_carts = append(new_carts, cart_val)
				}
			}
			carts = new_carts

		}

		if len(carts) < 2 {

			fmt.Println("Part2:", carts[0].x, ",", carts[0].y)
			break
		}

	}

}
