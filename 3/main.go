// Luke Green
// github.com/jlgre
// Advent of code problem 2

package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func countTrees(slope []string, right int, down int) int {
	var count int

	lineLength := len(slope[0])

	at := 0

	for i, line := range slope {
		if i%down == 0 {
			if traverse(at%lineLength, line) {
				count++
			}
			at += right
		}
	}

	return count
}

func traverse(at int, line string) bool {
	if line[at] == '#' {
		return true
	}
	return false
}

func main() {
	raw, err := ioutil.ReadFile("slope.dat")

	if err != nil {
		log.Fatal(err)
	}

	skips := [][]int{
		{1, 1},
		{3, 1}, // Part 1
		{5, 1},
		{7, 1},
		{1, 2},
	}

	slope := strings.Split(string(raw), "\n")

	product := 1

	for _, skip := range skips {
		trees := countTrees(slope, skip[0], skip[1])
		product *= trees
		fmt.Println(trees)
	}

	fmt.Println(product)
}
