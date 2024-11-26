package day06

import (
	util "AOC23/programs"
	"fmt"
	"strconv"
	"strings"
	"time"
)

// +++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ DAY 06A ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
func Day6a(m int) (int, time.Duration) {
	start := time.Now()
	var lines []string
	if m == 0 {
		lines = util.Returnlines("inputdata/day_06/day_06_test.txt")
	} else {
		lines = util.Returnlines("inputdata/day_06/day_06_actual.txt")
	}

	//Todays Vars
	var times []int

	timestrings := strings.Split(lines[0][10:], " ") //Create Time Strings
	for _, str := range timestrings {                //Convert To INTS
		if str == "" {
			asint, _ := strconv.Atoi(str)
			times = append(times, asint)
		}
	}

	fmt.Printf("%s", timestrings[0])

	return 5, time.Since(start)
}

// +++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ DAY 06B ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
func Day6b(m int) (int, time.Duration) {
	start := time.Now()
	//var lines []string
	if m == 0 {
		//lines = util.Returnlines("inputdata/day_06/day_06_test.txt")
	} else {
		//lines = util.Returnlines("inputdata/day_06/day_06_actual.txt")
	}

	return 5, time.Since(start)
}

// +++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ UTILS ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
