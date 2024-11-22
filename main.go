package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Starting The AdventOfCode23 Script")
	day1()
}

func day1() {
	fmt.Println("Starting Day 1")

	lines := returnlines("PuzzleInput/day1.txt")
	fmt.Println(lines[1])

}

func returnlines(filepath string) []string {
	file, _ := os.Open(filepath) //open file and ignore the errors
	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}
