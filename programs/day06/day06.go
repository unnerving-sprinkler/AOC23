package day06

import (
	util "AOC23/programs"
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
	print(lines)
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
