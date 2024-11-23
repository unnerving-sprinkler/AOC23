package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	fmt.Println("Starting The AdventOfCode23 Script")
	day1b()

	fmt.Printf("End Of Code | Execution Duration: %s\n", time.Since(start))
}

func day1a() {
	fmt.Println("Starting Day 1")
	lines := returnlines("PuzzleInput/day1.txt") //Read In Inputs

	//Setup VARS For Today
	answer := 0

	//Loop Through Each Line
	for i := 0; i < len(lines); i++ {
		j := strings.IndexAny(lines[i], "0123456789")     //j is the index of the first instance of a number
		k := strings.LastIndexAny(lines[i], "0123456789") //k is the index of the second instance of a number
		firstchar := rune(lines[i][j])                    //Convert CHARS to RUNE
		lastchar := rune(lines[i][k])
		fmt.Printf("First: %c | Last: %c | String: %s\n", firstchar, lastchar, lines[i]) //Debug Line
		linenumber := fmt.Sprint(string(firstchar), string(lastchar))                    //Using chars as a string combine them to make one string
		num, _ := strconv.Atoi(linenumber)                                               //Convert String To Number
		answer += num                                                                    //Add To running Total
	}

	fmt.Printf("Answer To Day 1a: %d\n", answer)
}

func day1b() {
	fmt.Println("Starting Day 1")
	lines := returnlines("PuzzleInput/day1.txt") //Read In Inputs

	//Setup VARS For Today
	answer := 0

	//Loop Through Each Line
	for i := 0; i < len(lines); i++ {
		j := strings.IndexAny(lines[i], "0123456789")     //j is the index of the first instance of a number
		k := strings.LastIndexAny(lines[i], "0123456789") //k is the index of the second instance of a number
		firstchar := rune(lines[i][j])                    //Convert CHARS to RUNE
		lastchar := rune(lines[i][k])
		fmt.Printf("First: %c | Last: %c | String: %s\n", firstchar, lastchar, lines[i]) //Debug Line
		linenumber := fmt.Sprint(string(firstchar), string(lastchar))                    //Using chars as a string combine them to make one string
		num, _ := strconv.Atoi(linenumber)                                               //Convert String To Number
		answer += num                                                                    //Add To running Total
	}

	fmt.Printf("Answer To Day 1a: %d\n", answer)
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
