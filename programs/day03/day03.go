package day03

import (
	util "AOC23/programs"
	"slices"
	"strconv"
	"strings"
	"time"
)

// +++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ DAY 03A ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
func Day3a(m int) (int, time.Duration) {
	start := time.Now()
	var lines []string
	if m == 0 {
		lines = util.Returnlines("inputdata/day_03/day_03_test.txt")
	} else {
		lines = util.Returnlines("inputdata/day_03/day_03_actual.txt")
	}

	//Todays VARS
	symbols := "!@*+%#/$&--!@#$%^&*()="
	var numbers = []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "0"}
	var partnumbers []partnumber //Store Our Parts We Find (Will Index Later)
	var score int

	//Logic Section
	for i := 0; i < len(lines); i++ { //Index Through Each Line Of The Problem
		var currentnumber []string
		for j := 0; j < len(lines[i]); j++ { //Index Through Each Char Of That Line
			if slices.Contains(numbers, string(lines[i][j])) { //Check If This Char is a Number
				currentnumber = append(currentnumber, string(lines[i][j])) // If This is a Number Append That Number To Current Number Array
			} else {
				currentnumber = nil //If this CHAR is not a number Zero Out The Search Array
			}
			if ((j == len(lines[i])-1) || !(slices.Contains(numbers, string(lines[i][j+1])))) && len(currentnumber) > 0 { //Was This This End Of The Line Or Is Next CHAR Going To Be Not A Number

				newPart := partnumber{ //Generate a Part Info Sheet About This Part
					PN:        strings.Join(currentnumber, ""),
					POSXStart: int16(j) - int16(len(currentnumber)-1),
					POSXEnd:   int16(j) + 1,
					POSY:      int16(i),
					Score:     0,
					Valid:     false,
				}
				//fmt.Printf("PN || PN:%-5s || POSXStart:%-3d || POSXEnd:%-3d || POSY:%3d\n", newPart.PN, newPart.POSXStart, newPart.POSXEnd, newPart.POSY)
				partnumbers = append(partnumbers, newPart) // Append This Part Onto Our Running Parts List

			}
		}
	}
	for i := 0; i < len(partnumbers); i++ { //Convert The Scores Into INTS
		value, _ := strconv.Atoi(partnumbers[i].PN)
		partnumbers[i].Score = int16(value)
	}
	for i := 0; i < len(partnumbers); i++ { //Determine If Valid Part Number
		var surroundingchars []string
		fullLeft := true
		fullRight := true
		//Add Chars After The Number To Surrounding Chars
		if partnumbers[i].POSXEnd < 140 {
			fullRight = false
			charAfter := string(lines[partnumbers[i].POSY][partnumbers[i].POSXEnd])
			surroundingchars = append(surroundingchars, charAfter)
		}
		//Aff Chars After Number To Surrounding Chars
		if partnumbers[i].POSXStart > 1 {
			fullLeft = false
			charBefore := string(lines[partnumbers[i].POSY][(partnumbers[i].POSXStart)-1])
			surroundingchars = append(surroundingchars, charBefore)
		}
		//Add Chars Above Number To Surrounding Chars
		if partnumbers[i].POSY > 0 {
			leftchar := partnumbers[i].POSXStart
			rightchar := partnumbers[i].POSXEnd

			if !fullLeft {
				leftchar += -1
			}
			if !fullRight {
				rightchar += 1
			}

			charAbove := string(lines[partnumbers[i].POSY-1][leftchar:rightchar])
			surroundingchars = append(surroundingchars, charAbove)
		}
		//Add Chars Below Number To Surrounding Chars
		if partnumbers[i].POSY < int16(len(lines)-1) {
			leftchar := partnumbers[i].POSXStart
			rightchar := partnumbers[i].POSXEnd

			if !fullLeft {
				leftchar += -1
			}
			if !fullRight {
				rightchar += 1
			}

			charAbove := string(lines[partnumbers[i].POSY+1][leftchar:rightchar])
			surroundingchars = append(surroundingchars, charAbove)
		}
		allsurrounding := strings.Join(surroundingchars, "")
		partnumbers[i].Valid = strings.ContainsAny(allsurrounding, symbols)

	}
	for i := 0; i < len(partnumbers); i++ { //Score Points
		if partnumbers[i].Valid {
			score += int(partnumbers[i].Score)
		}
	}

	return score, time.Since(start)
}

type partnumber struct {
	PN        string
	POSXStart int16
	POSXEnd   int16
	POSY      int16
	Score     int16
	Valid     bool
}

// +++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ DAY 03B ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
func Day3b(m int) (int, time.Duration) {
	start := time.Now()
	//var lines []string
	if m == 0 {
		//lines = util.Returnlines("inputdata/day_03/day_03_test.txt")
	} else {
		//lines = util.Returnlines("inputdata/day_03/day_03_actual.txt")
	}

	return 0, time.Since(start)
}
