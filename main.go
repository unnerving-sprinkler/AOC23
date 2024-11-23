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
	day1b() //Complete
	day2a() //Not Complete

	fmt.Printf("\nEnd Of Code | Execution Duration: %s\n", time.Since(start))
}

// +++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ DAY 1A +++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
func day1a() {
	fmt.Println("Starting Day 1a")
	//lines := returnlines("inputdata/day_01/day_01_test.txt") //Read In Inputs
	lines := returnlines("inputdata/day_01/day_01_actual.txt") //Read In Inputs

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
func day1b() {
	fmt.Println("\nStarting Day 1b")
	//lines := returnlines("inputdata/day_01/day_01_test.txt") //Read In Inputs
	lines := returnlines("inputdata/day_01/day_01_actual.txt") //Read In Inputs

	//Setup VARS For Today
	answer := 0
	var digits = []string{"1", "2", "3", "4", "5", "6", "7", "8", "9"}
	var digitname = []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
	var reversedigitname = []string{"eno", "owt", "eerht", "ruof", "evif", "xis", "neves", "thgie", "enin"}

	//Loop Through Each Line
	for i := 0; i < len(lines); i++ {
		lines[i] = reversestring(lines[i]) //Flip Array In Reverse
		maxIndex := 0
		minValue := 99
		for j := 0; j < len(digits); j++ { //Loop Through To Find The Last (Now First) NumWord
			currentscore := strings.Index(lines[i], reversedigitname[j]) //Score Is The Index Of Where The Word Appears
			if currentscore > -1 {
				if currentscore < minValue {
					minValue = currentscore
					maxIndex = j
				}
			}
		}
		minNumValue := 99
		for j := 0; j < len(digits); j++ { //Loop Through To Find The Last (Now First) Number
			currentscore := strings.Index(lines[i], digits[j])
			if currentscore > -1 {
				if currentscore < minNumValue {
					minNumValue = currentscore
				}
			}
		}
		if minValue < minNumValue { //If There Was A Number Close To The End (Now Begining) Ignore, If Not Replace
			lines[i] = strings.Replace(lines[i], reversedigitname[maxIndex], digits[maxIndex], 1)
		}
		lines[i] = reversestring(lines[i])
		maxIndex = 0
		minValue = 99
		for j := 0; j < len(digits); j++ {
			currentscore := strings.Index(lines[i], digitname[j]) //Score Is The Index Of Where The Word Appears
			if currentscore > -1 {
				if currentscore < minValue {
					minValue = currentscore
					maxIndex = j
				}
			}
		}
		//Find The Index Of The Last Number(Not Word) In The Line
		minNumValue = 99
		for j := 0; j < len(digits); j++ {
			currentscore := strings.Index(lines[i], digits[j])
			if currentscore > -1 {
				if currentscore < minNumValue {
					minNumValue = currentscore
				}

			}
			print()
		}
		//If There Was a Number After The Word Ignore The Word
		if minValue < minNumValue {
			lines[i] = strings.Replace(lines[i], digitname[maxIndex], digits[maxIndex], 1) // Replace the number we know is at the end first
		} else {

		}

		fmt.Printf("%s | After Both Replace\n", lines[i])

		j := strings.IndexAny(lines[i], "0123456789")     //j is the index of the first instance of a number
		k := strings.LastIndexAny(lines[i], "0123456789") //k is the index of the second instance of a number
		firstchar := rune(lines[i][j])                    //Convert CHARS to RUNE
		lastchar := rune(lines[i][k])
		fmt.Printf("|First: %c | Last: %c | String: %s\n", firstchar, lastchar, lines[i]) //Debug Line
		fmt.Printf("%c%c\n", firstchar, lastchar)                                         //Debug Line
		linenumber := fmt.Sprint(string(firstchar), string(lastchar))                     //Using chars as a string combine them to make one string
		num, _ := strconv.Atoi(linenumber)                                                //Convert String To Number
		answer += num                                                                     //Add To running Total
	}

	fmt.Printf(">Answer To Day 1b: %d\n", answer)
}

// +++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ DAY 2A +++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
func day2a() {
	fmt.Println("\nStarting Day 2a")
	//lines := returnlines("inputdata/day_01/day_01_test.txt") //Read In Inputs
	lines := returnlines("inputdata/day_01/day_01_actual.txt") //Read In Inputs

	for i := 0; i < len(lines); i++ {
		print(lines[i])
	}
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
