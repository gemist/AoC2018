package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type node struct {
	metadata []int
	children []*node
}

func readNode(licenseData *[]int) *node {

	nchildren := (*licenseData)[0]
	nmetadata := (*licenseData)[1]
	*licenseData = (*licenseData)[2:]
	children := make([]*node, 0)

	for i := 0; i < nchildren; i++ {

		children = append(children, readNode(licenseData))
	}

	metadata := make([]int, 0)
	for i := 0; i < nmetadata; i++ {
		metadata = append(metadata, (*licenseData)[i])
	}
	*licenseData = (*licenseData)[nmetadata:]

	return &node{metadata, children}
}

func traverse(root *node, f func(*node)) {
	f(root)
	for _, child := range root.children {
		traverse(child, f)
	}
}

func nodeValue(node *node) int {
	sum := 0
	if len(node.children) == 0 {
		for _, metadata := range node.metadata {
			sum += metadata
		}
	} else {

		for _, metadata := range node.metadata {
			index := metadata - 1
			if len(node.children) <= index {
				continue
			}
			sum += nodeValue(node.children[index])
		}
	}
	return sum
}

func main() {
	file, _ := os.Open("input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	license := make([]int, 0)

	for scanner.Scan() {
		line := strings.Split(strings.TrimSpace(scanner.Text()), " ")
		for _, val := range line {
			s, _ := strconv.Atoi(val)
			license = append(license, s)
		}
	}
	sum := 0
	root := readNode(&license)
	//	fmt.Println(root.metadata)
	traverse(root, func(node *node) {
		for _, metadata := range node.metadata {
			sum += metadata
		}
	})

	fmt.Println("part 1: ", sum)

	fmt.Println("part 2:", nodeValue(root))
}
