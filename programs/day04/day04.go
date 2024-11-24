package day04

import (
	util "AOC23/programs"
	"fmt"
	"math"
	"strconv"
	"strings"
)

func Day4a() {
	fmt.Println("\nStarting Day 4a")
	//lines := uril.Returnlines("inputdata/day_04/day_04_test.txt") //Read In Inputs
	lines := util.Returnlines("inputdata/day_04/day_04_actual.txt") //Read In Inputs

	// Variables For Today
	var score float64
	score = 0

	//Logic
	var scratchoffdeck []Scratchoff
	for i := 0; i < len(lines); i++ { //Text Handeling Portion -- Extract Card Numbers, Winning Numbers, and My Numbers
		var thiscard Scratchoff
		thiscard.CardID = ExtractCardNumber(lines[i])
		thiscard.WinningNumbers = ExtractWinningNumbers(lines[i])
		thiscard.MyNumbers = ExtractMyNumbers(lines[i])
		scratchoffdeck = append(scratchoffdeck, thiscard)
	}
	for i := 0; i < len(scratchoffdeck); i++ { //Check How Many Winning Scores
		runningscore := 0
		for j := 0; j < len(scratchoffdeck[i].MyNumbers); j++ {
			for k := 0; k < len(scratchoffdeck[i].WinningNumbers); k++ {
				if scratchoffdeck[i].WinningNumbers[k] == scratchoffdeck[i].MyNumbers[j] {
					runningscore += 1
				}
			}
		}
		scratchoffdeck[i].Score = runningscore
	}

	for _, card := range scratchoffdeck { //Do Math To Convery Winning Numbers to Score
		if card.Score > 0 {
			score += math.Pow(2, float64(card.Score)-1)
		}
	}
	fmt.Printf(">Answer To Day 4a: %.0f\n", score)
}

type Scratchoff struct {
	CardID         int
	WinningNumbers []int
	MyNumbers      []int
	Score          int
}

func ExtractCardNumber(s string) int { //Extract The Card Number From The Card
	NumberString := s[5:8]
	var NumberInt int
	NumberInt, _ = strconv.Atoi(strings.TrimLeft(NumberString, " "))
	return int(NumberInt)
}

func ExtractWinningNumbers(s string) []int { //Extract The Winning Numbers from The Card
	var winningnumbers []int                         //This Is What Will Be Returned
	winningString := s[9:40]                         //String That Contains The Winning Numbesr
	splitString := strings.Split(winningString, " ") //Split The String On Every Space
	for _, item := range splitString {               //Last Operation Made Several Blank Entries - Remove them
		if item != "" {
			var itemconv int
			itemconv, _ = strconv.Atoi(item)
			winningnumbers = append(winningnumbers, itemconv)
		}
	}
	return winningnumbers
}
func ExtractMyNumbers(s string) []int { //Extract The Winning Numbers from The Card
	var MyNumbers []int                              //This Is What Will Be Returned
	winningString := s[42:116]                       //String That Contains The Winning Numbesr
	splitString := strings.Split(winningString, " ") //Split The String On Every Space
	for _, item := range splitString {               //Last Operation Made Several Blank Entries - Remove them
		if item != "" {
			var itemconv int
			itemconv, _ = strconv.Atoi(item)
			MyNumbers = append(MyNumbers, itemconv)
		}
	}
	return MyNumbers
}

// +++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ DAY 4b +++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
func Day4b() {
	fmt.Println("\nStarting Day 4b")

	fmt.Printf(">Answer To Day 4b: %d\n", 0)
}
