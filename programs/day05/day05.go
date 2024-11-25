package day05

import (
	util "AOC23/programs"
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

	//Convert All Maps To INT
	for i := 2; i < len(lines); i++ {
		if len(lines[i]) > 0 && slices.Contains(numbers, lines[i][:1]) { //Does This Line Contrain Info
			mapstring := (strings.Split(lines[i], " "))

			var thismap Map
			thismap.DestinationRangeStart, _ = strconv.Atoi(mapstring[0]) //Map Info To Structure
			thismap.SourceRangeStart, _ = strconv.Atoi(mapstring[1])
			thismap.RangeLength, _ = strconv.Atoi(mapstring[2])

			if i > 2 && i < 26 { //seed-to-soil map:
				maps[0] = append(maps[0], thismap)
			} else if i > 27 && i < 68 { //soil-to-fertilizer map:
				maps[1] = append(maps[1], thismap)
			} else if i > 69 && i < 84 { //fertilizer-to-water map:
				maps[2] = append(maps[2], thismap)
			} else if i > 85 && i < 132 { //water-to-light map:
				maps[3] = append(maps[3], thismap)
			} else if i > 133 && i < 169 { //light-to-temperature map:
				maps[4] = append(maps[4], thismap)
			} else if i > 170 && i < 199 { //temperature-to-humidity map:
				maps[5] = append(maps[5], thismap)
			} else if i > 200 { //humidity-to-location map:
				maps[6] = append(maps[6], thismap)
			}

		}

	}

	return 0, time.Since(start)
}

type Map struct {
	DestinationRangeStart int
	SourceRangeStart      int
	RangeLength           int
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
