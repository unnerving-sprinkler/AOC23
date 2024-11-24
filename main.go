package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	fmt.Println("Starting The AdventOfCode23 Script")
	day1a() //Complete
	day1b() //Complete
	day2a() //Complete
	day2b() //Complete
	day3a() //Complete
	day3b() //Not Complete

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
		//fmt.Printf("|First: %c | Last: %c | String: %s\n", firstchar, lastchar, lines[i]) //Debug Line
		linenumber := fmt.Sprint(string(firstchar), string(lastchar)) //Using chars as a string combine them to make one string
		num, _ := strconv.Atoi(linenumber)                            //Convert String To Number
		answer += num                                                 //Add To running Total
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

		j := strings.IndexAny(lines[i], "0123456789")     //j is the index of the first instance of a number
		k := strings.LastIndexAny(lines[i], "0123456789") //k is the index of the second instance of a number
		firstchar := rune(lines[i][j])                    //Convert CHARS to RUNE
		lastchar := rune(lines[i][k])
		linenumber := fmt.Sprint(string(firstchar), string(lastchar)) //Using chars as a string combine them to make one string
		num, _ := strconv.Atoi(linenumber)                            //Convert String To Number
		answer += num                                                 //Add To running Total
	}

	fmt.Printf(">Answer To Day 1b: %d\n", answer)
}

// +++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ DAY 2A +++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
func day2a() {
	fmt.Println("\nStarting Day 2a")
	//lines := returnlines("inputdata/day_02/day_02_test.txt") //Read In Inputs
	lines := returnlines("inputdata/day_02/day_02_actual.txt") //Read In Inputs

	//Todays VARS
	puzzleanswer := 0

	for i := 0; i < len(lines); i++ {
		gamepossible := true
		gameinformationstring := lines[i][(strings.Index(lines[i], ":"))+2:] //Strip The Name Of The Game Off String
		gameinformationstring = gameinformationstring + string("; ")         //Put a Ending ; To Make Our Jobs Easier Later
		var rounds []string
		for {
			rounds = append(rounds, gameinformationstring[:(strings.Index(gameinformationstring, ";"))])    //Append Each Round Info To rounds
			gameinformationstring = gameinformationstring[(strings.Index(gameinformationstring, ";") + 2):] //Remove The Round we Just Counted From The Game
			if len(gameinformationstring) < 4 {                                                             //Break Loop When There Are No More Rounds
				break
			}
		}
		for j := 0; j < len(rounds); j++ {
			roundinfo := ExtractRoundInformation(rounds[j]) //For Each Game Extract Information About How Many Cubes
			if roundinfo[0] > 12 {
				gamepossible = false
			}
			if roundinfo[1] > 13 {
				gamepossible = false
			}
			if roundinfo[2] > 14 {
				gamepossible = false
			}
		}
		if gamepossible {
			puzzleanswer += ExtractGameNumber(lines[i])
		}
	}
	fmt.Printf(">Answer To Day 2a: %d\n", puzzleanswer)
}

// +++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ DAY 2B +++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
func day2b() {
	fmt.Println("\nStarting Day 2b")
	//lines := returnlines("inputdata/day_02/day_02_test.txt") //Read In Inputs
	lines := returnlines("inputdata/day_02/day_02_actual.txt") //Read In Inputs

	//Todays VARS
	puzzleanswer := 0

	for i := 0; i < len(lines); i++ {
		var gameminimumblocks [3]int
		gameinformationstring := lines[i][(strings.Index(lines[i], ":"))+2:] //Strip The Name Of The Game Off String
		gameinformationstring = gameinformationstring + string("; ")         //Put a Ending ; To Make Our Jobs Easier Later
		var rounds []string
		for {
			rounds = append(rounds, gameinformationstring[:(strings.Index(gameinformationstring, ";"))])    //Append Each Round Info To rounds
			gameinformationstring = gameinformationstring[(strings.Index(gameinformationstring, ";") + 2):] //Remove The Round we Just Counted From The Game
			if len(gameinformationstring) < 4 {                                                             //Break Loop When There Are No More Rounds
				break
			}
		}
		for j := 0; j < len(rounds); j++ {
			roundinfo := ExtractRoundInformation(rounds[j]) //For Each Game Extract Information About How Many Cubes
			if roundinfo[0] > gameminimumblocks[0] {
				gameminimumblocks[0] = roundinfo[0]
			}
			if roundinfo[1] > gameminimumblocks[1] {
				gameminimumblocks[1] = roundinfo[1]
			}
			if roundinfo[2] > gameminimumblocks[2] {
				gameminimumblocks[2] = roundinfo[2]
			}
		}
		puzzleanswer += (gameminimumblocks[0] * gameminimumblocks[1] * gameminimumblocks[2])
	}
	fmt.Printf(">Answer To Day 2b: %d\n", puzzleanswer)
}

// Return The Game Number (Used For Day 2)
func ExtractGameNumber(s string) int {
	withoutprefix, _ := strings.CutPrefix(s, "Game ") //Remove the Prefix Of The Line
	w := strings.Index(withoutprefix, ":")            //Determine Where the : Char Is
	gamenumberstring := withoutprefix[:w]             //Strip Text String Based On That Information
	gamenumber, _ := strconv.Atoi(gamenumberstring)   //Convert Text To int
	return gamenumber
}

// Returns Information About The Round Fed Into The Function:
// Output Is Generated In This Format | [red, green, blue]
func ExtractRoundInformation(s string) [3]int {
	var colorinfo [3]int
	s = s + string(",  ") // Add , to the end will help us out later
	var roundinfo []string
	for {
		roundinfo = append(roundinfo, s[:strings.Index(s, ",")]) //Information About 1 Cube
		s = s[strings.Index(s, ",")+2:]                          // Trim Info We Just Took Out Of s
		if len(s) < 2 {                                          //Break After We've Completed All Handfulls
			break
		}
	}
	//fmt.Printf("%s\n", roundinfo[0])
	for i := 0; i < len(roundinfo); i++ {
		resultred, foundred := strings.CutSuffix(roundinfo[i], " red")
		resultgreen, foundgreen := strings.CutSuffix(roundinfo[i], " green")
		resultblue, foundblue := strings.CutSuffix(roundinfo[i], " blue")

		if foundred {
			colorinfo[0], _ = strconv.Atoi(resultred)
		} else if foundgreen {
			colorinfo[1], _ = strconv.Atoi(resultgreen)
		} else if foundblue {
			colorinfo[2], _ = strconv.Atoi(resultblue)
		}
	}
	//fmt.Printf("%d|%d|%d", colorinfo[0], colorinfo[1], colorinfo[2])
	return colorinfo
}

// +++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ DAY 3A +++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
func day3a() {
	fmt.Println("\nStarting Day 3a")
	lines := returnlines("inputdata/day_03/day_03_test.txt") //Read In Inputs
	//lines := returnlines("inputdata/day_03/day_03_actual.txt") //Read In Inputs

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
	fmt.Printf(">Answer To Day 3a: %d\n", score)
}

type partnumber struct {
	PN        string
	POSXStart int16
	POSXEnd   int16
	POSY      int16
	Score     int16
	Valid     bool
}

// +++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ DAY 3B +++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
func day3b() {

}

// +++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ Supporting Functions +++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
func returnlines(filepath string) []string {
	var lines []string

	//Logic
	file, _ := os.Open(filepath)      //Open File Ignoring Errors
	scanner := bufio.NewScanner(file) //Scan File In
	for scanner.Scan() {
		lines = append(lines, scanner.Text()) //Append Each Line Into An Array
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
