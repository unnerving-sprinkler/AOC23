package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("Starting The AdventOfCode23 Script")
	day1a()
}

func day1a() {
	fmt.Println("Starting Day 1")
	lines := returnlines("PuzzleInput/day1.txt")
	//fmt.Println(lines)

	for i := 0; i < len(lines); i++ { //Loop Through Each Entry
		j := strings.IndexAny(lines[i], "0123456789") //j is the first index of any number
		println(lines[i])
		println(j)
		char := lines[i][j]
		fmt.Printf("The Value Of The First Char is: %c", char)
		println("next")

	}

}

func returnlines(filepath string) []string {
	//Create VARS
	var lines []string

	//Logic
	file, _ := os.Open(filepath) //open file and ignore the errors
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}
