package day05

import (
	util "AOC23/programs"
	"fmt"
	"slices"
	"strconv"
	"strings"
	"time"
)

// +++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ DAY 05A ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
func Day5a(m int) (int, time.Duration) {
	start := time.Now()
	var lines []string
	if m == 0 {
		lines = util.Returnlines("inputdata/day_05/day_05_test.txt")
	} else {
		lines = util.Returnlines("inputdata/day_05/day_05_actual.txt")
	}

	// Todays Vars
	var seeds []int
	var maps [7][]Map
	var numbers = []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "0"}

	seedstring := (strings.Split(lines[0][7:], " ")) //Split Seed Line
	for i := 0; i < len(seedstring); i++ {           //Convet Items To Type INT
		itemconv, _ := strconv.Atoi(seedstring[i])
		seeds = append(seeds, itemconv)
	}

	for i := 2; i < len(lines); i++ {
		if len(lines[i]) > 0 && slices.Contains(numbers, lines[i][:1]) {
			mapstring := (strings.Split(lines[i]))

			var thismap Map

			for j := 0; j < len(mapstring); j++ {
				itemconv, _ := strconv.Atoi(seedstring[i])

			}
			fmt.Printf("%s\n", (lines[i]))
		}

	}

	return 0, time.Since(start)
}

type Map struct {
	DestinationRangeStart int
	SourceRangeStart      []int
	RangeLength           []int
}

// +++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ DAY 05B ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
func Day5b(m int) (int, time.Duration) {
	start := time.Now()
	//var lines []string
	if m == 0 {
		//lines = util.Returnlines("inputdata/day_04/day_04_test.txt")
	} else {
		//lines = util.Returnlines("inputdata/day_04/day_04_actual.txt")
	}
	return 0, time.Since(start)
}

// +++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ UTILS ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
