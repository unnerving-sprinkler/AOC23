package main

import (
	day01 "AOC23/programs/day01"
	day02 "AOC23/programs/day02"
	day03 "AOC23/programs/day03"
	day04 "AOC23/programs/day04"
	day05 "AOC23/programs/day05"
	"fmt"
)

/*


 */

func main() {
	fmt.Println("Starting The AdventOfCode23 Script")

	// CHOOSE SETTINGS
	m := 0 // 0 = Test Input | 1 = Actual Input

	day1aans, day1atime := day01.Day1a(m) //Complete
	day1bans, day1btime := day01.Day1b(m) //Complete
	day2aans, day2atime := day02.Day2a(m) //Complete
	day2bans, day2btime := day02.Day2b(m) //Complete
	day3aans, day3atime := day03.Day3a(m) //Complete
	day3bans, day3btime := day03.Day3b(m) //Not Started
	day4aans, day4atime := day04.Day4a(m) //Complete
	day4bans, day4btime := day04.Day4b(m) //Not Started
	day5aans, day5atime := day05.Day5a(m) //Not Started
	day5bans, day5btime := day05.Day5b(m) //Not Started

	fmt.Printf("Day 01a | Answer: %-11d | Taking: %s\n", day1aans, day1atime)
	fmt.Printf("Day 01b | Answer: %-11d | Taking: %s\n", day1bans, day1btime)
	fmt.Printf("Day 02a | Answer: %-11d | Taking: %s\n", day2aans, day2atime)
	fmt.Printf("Day 02b | Answer: %-11d | Taking: %s\n", day2bans, day2btime)
	fmt.Printf("Day 03a | Answer: %-11d | Taking: %s\n", day3aans, day3atime)
	fmt.Printf("Day 03b | Answer: %-11d | Taking: %s\n", day3bans, day3btime)
	fmt.Printf("Day 04a | Answer: %-11d | Taking: %s\n", day4aans, day4atime)
	fmt.Printf("Day 04b | Answer: %-11d | Taking: %s\n", day4bans, day4btime)
	fmt.Printf("Day 05a | Answer: %-11d | Taking: %s\n", day5aans, day5atime)
	fmt.Printf("Day 05b | Answer: %-11d | Taking: %s\n", day5bans, day5btime)
}
