// Luke Green
// github.com/jlgre
// Advent of Code day 2

package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

type Password struct {
	min      int
	max      int
	letter   string
	password string
}

func parseString(input string) Password {
	var min int
	var max int
	var letter string
	var pass string

	var err error

	var min_str string
	var max_str string

	for _, char := range input {
		if min == 0 {
			if char == '-' {
				min, err = strconv.Atoi(min_str)
				continue
			}

			min_str += string(char)
		} else if max == 0 {
			if char == ' ' {
				max, err = strconv.Atoi(max_str)
				continue
			}

			max_str += string(char)
		} else if letter == "" {
			letter = string(char)
		} else if char != ':' && char != ' ' {
			pass += string(char)
		}

		if err != nil {
			log.Fatal(err)
		}
	}

	return Password{
		min, max, letter, pass,
	}
}

func countOccurences(letter string, password string) int {
	var count int

	for _, char := range password {
		if letter == string(char) {
			count++
		}
	}

	return count
}

func checkPasswordsPart1(passwords []Password) []Password {
	var valid []Password

	for _, password := range passwords {
		count := countOccurences(password.letter, password.password)
		if count >= password.min && count <= password.max {
			valid = append(valid, password)
		}
	}

	return valid
}

func checkPasswordsPart2(passwords []Password) []Password {
	var valid []Password

	for _, password := range passwords {
		inPositionOne := string(password.password[password.min-1]) == password.letter
		inPositionTwo := string(password.password[password.max-1]) == password.letter

		if (inPositionOne || inPositionTwo) && !(inPositionOne && inPositionTwo) {
			valid = append(valid, password)
		}
	}

	return valid
}

func main() {
	var passwords []Password

	raw, err := ioutil.ReadFile("passwords.dat")
	if err != nil {
		log.Fatal(err)
	}

	strarr := strings.Split(string(raw), "\n")

	for _, val := range strarr {
		passwords = append(passwords, parseString(val))
	}

	fmt.Println(len(checkPasswordsPart1(passwords)))
	fmt.Println(len(checkPasswordsPart2(passwords)))
}
