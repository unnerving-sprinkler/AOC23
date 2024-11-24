package main

import (
	day01 "AOC23/programs/day01"
	day02 "AOC23/programs/day02"
	day03 "AOC23/programs/day03"
	day04 "AOC23/programs/day04"
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	fmt.Println("Starting The AdventOfCode23 Script")
	day01.Day1a() //Complete
	day01.Day1b() //Complete
	day02.Day2a() //Complete
	day02.Day2b() //Complete
	day03.Day3a() //Complete
	day03.Day3b() //Not Complete
	day04.Day4a() //Complete
	day04.Day4b() //Not Complete

	fmt.Printf("\nEnd Of Code | Execution Duration: %s\n", time.Since(start))
}
