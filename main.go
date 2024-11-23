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
	day1a() //Complete
	day1b() //Not Complete

	fmt.Printf("\nEnd Of Code | Execution Duration: %s\n", time.Since(start))
}

// +++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ DAY 1A +++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
func day1a() {
	fmt.Println("Starting Day 1a")
	lines := returnlines("inputdata/day_01/day_01_test.txt") //Read In Inputs

	//Setup VARS For Today
	answer := 0

	//Loop Through Each Line
	for i := 0; i < len(lines); i++ {
		j := strings.IndexAny(lines[i], "0123456789")     //j is the index of the first instance of a number
		k := strings.LastIndexAny(lines[i], "0123456789") //k is the index of the second instance of a number
		firstchar := rune(lines[i][j])                    //Convert CHARS to RUNE
		lastchar := rune(lines[i][k])
		fmt.Printf("|First: %c | Last: %c | String: %s\n", firstchar, lastchar, lines[i]) //Debug Line
		linenumber := fmt.Sprint(string(firstchar), string(lastchar))                     //Using chars as a string combine them to make one string
		num, _ := strconv.Atoi(linenumber)                                                //Convert String To Number
		answer += num                                                                     //Add To running Total
	}

	fmt.Printf(">Answer To Day 1a: %d\n", answer)
}

// +++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ DAY 1B +++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
// Day1b is a copy of the day1a code but before the lies are parsed the replacements are made
func day1b() {
	fmt.Println("\nStarting Day 1b")
	lines := returnlines("inputdata/day_01/day_01_test.txt") //Read In Inputs

	//Setup VARS For Today
	answer := 0

	//Loop Through Each Line
	for i := 0; i < len(lines); i++ {

		//Check For First Instance Of a Number
		lines[i] = strings.Replace(lines[i], "one", "1", 1)
		lines[i] = strings.Replace(lines[i], "four", "4", 1)
		lines[i] = strings.Replace(lines[i], "five", "5", 1)
		lines[i] = strings.Replace(lines[i], "six", "6", 1)
		lines[i] = strings.Replace(lines[i], "seven", "7", 1)
		lines[i] = strings.Replace(lines[i], "eight", "8", 1)
		lines[i] = strings.Replace(lines[i], "nine", "9", 1)

		lines[i] = reversestring(lines[i])
		lines[i] = strings.Replace(lines[i], "eno", "1", 1)
		lines[i] = strings.Replace(lines[i], "ruof", "4", 1)
		lines[i] = strings.Replace(lines[i], "evif", "5", 1)
		lines[i] = strings.Replace(lines[i], "xis", "6", 1)
		lines[i] = strings.Replace(lines[i], "neves", "7", 1)
		lines[i] = strings.Replace(lines[i], "thgie", "8", 1)
		lines[i] = strings.Replace(lines[i], "enin", "9", 1)
		lines[i] = reversestring(lines[i])

		j := strings.IndexAny(lines[i], "0123456789")     //j is the index of the first instance of a number
		k := strings.LastIndexAny(lines[i], "0123456789") //k is the index of the second instance of a number
		firstchar := rune(lines[i][j])                    //Convert CHARS to RUNE
		lastchar := rune(lines[i][k])
		fmt.Printf("|First: %c | Last: %c | String: %s\n", firstchar, lastchar, lines[i]) //Debug Line
		linenumber := fmt.Sprint(string(firstchar), string(lastchar))                     //Using chars as a string combine them to make one string
		num, _ := strconv.Atoi(linenumber)                                                //Convert String To Number
		answer += num                                                                     //Add To running Total
	}

	fmt.Printf(">Answer To Day 1b: %d\n", answer)
}

// +++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ DAY 2A +++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
func day2a() {

}

// +++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ Supporting Functions +++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
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

func reversestring(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
