package probs

import (
	"aoc/utils"
	"fmt"
)

var outcomeMap = map[string]int{
	"lose": 0,
	"tie":  3,
	"win":  6,
}
var scoreMap = map[rune]int{
	'X': 1,
	'Y': 2,
	'Z': 3,
}
var winCodeMap = map[rune]string{
	'X': "lose",
	'Y': "tie",
	'Z': "win",
}

func Day02(isPartTwo bool, inFile string) {
	lines := utils.GetLinesFromFile(inFile)
	total := 0

	for _, line := range lines {
		lChar := rune(line[0])
		rChar := rune(line[2])
		if isPartTwo {
			total += outcomeMap[winCodeMap[rChar]] + getShapeScoreUsed(lChar, rChar)
		} else {
			total += scoreMap[rChar] + getOutcome(lChar, rChar)
		}
	}

	fmt.Println("total:", total)
}

func getOutcome(lChar rune, rChar rune) int {
	switch lChar {
	case 'A':
		switch rChar {
		case 'X':
			return outcomeMap["tie"]
		case 'Y':
			return outcomeMap["win"]
		case 'Z':
			return outcomeMap["lose"]
		}
	case 'B':
		switch rChar {
		case 'X':
			return outcomeMap["lose"]
		case 'Y':
			return outcomeMap["tie"]
		case 'Z':
			return outcomeMap["win"]
		}
	case 'C':
		switch rChar {
		case 'X':
			return outcomeMap["win"]
		case 'Y':
			return outcomeMap["lose"]
		case 'Z':
			return outcomeMap["tie"]
		}
	}
	return 0
}

func getShapeScoreUsed(lChar rune, rChar rune) int {
	switch lChar {
	case 'A':
		switch rChar {
		case 'X':
			return scoreMap['Z']
		case 'Y':
			return scoreMap['X']
		case 'Z':
			return scoreMap['Y']
		}
	case 'B':
		switch rChar {
		case 'X':
			return scoreMap['X']
		case 'Y':
			return scoreMap['Y']
		case 'Z':
			return scoreMap['Z']
		}
	case 'C':
		switch rChar {
		case 'X':
			return scoreMap['Y']
		case 'Y':
			return scoreMap['Z']
		case 'Z':
			return scoreMap['X']
		}
	}
	return 0
}
