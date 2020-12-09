// Luke Green - jlgre
// github.com/jlgre
// Advent of code 2020 - problem 1

package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func arraySum(arr []int, target int) int {
	for i, x := range arr {
		for _, y := range arr[i:] {
			if x+y == target {
				return x * y
			}
		}
	}

	return -1
}

func main() {
	var arr []int
	var target int
	target = 2020

	raw, err := ioutil.ReadFile("nums.dat")

	if err != nil {
		log.Fatal(err)
	}

	strarr := strings.Split(string(raw), "\n")
	for _, val := range strarr {
		ival, err := strconv.Atoi(val)
		if err != nil {
			log.Fatal(err)
		}
		arr = append(arr, ival)
	}

	value := arraySum(arr, target)
	fmt.Println(value)
}
