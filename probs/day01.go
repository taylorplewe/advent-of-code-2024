package probs

import (
	"aoc/utils"
	"fmt"
	"regexp"
	"slices"
	"strconv"
)

var leftNums []int
var rightNums []int
var rightNumOccurrences map[int]int

func Day01(isPartTwo bool, inFile string) {
	lines := utils.GetLinesFromFile(inFile)
	leftNums = make([]int, 0)
	rightNums = make([]int, 0)
	rightNumOccurrences = make(map[int]int)

	for _, line := range lines {
		reg := regexp.MustCompile(`(\d+)\s*(\d+)`)
		matches := reg.FindStringSubmatch(line)
		if len(matches) < 3 {
			continue
		}

		lNum, _ := strconv.Atoi(matches[1])
		rNum, _ := strconv.Atoi(matches[2])
		leftNums = append(leftNums, lNum)
		rightNums = append(rightNums, rNum)
		if _, exists := rightNumOccurrences[rNum]; !exists {
			rightNumOccurrences[rNum] = 0
		}
		rightNumOccurrences[rNum]++

	}

	if !isPartTwo {
		slices.Sort(leftNums)
		slices.Sort(rightNums)

		totalDiff := 0
		for i := range leftNums {
			if rightNums[i] > leftNums[i] {
				totalDiff += rightNums[i] - leftNums[i]
			} else {
				totalDiff += leftNums[i] - rightNums[i]
			}
		}

		// fmt.Println(leftNums)
		// fmt.Println(rightNums)
		fmt.Println("diff", totalDiff)
	} else {
		totalDiff := 0
		for _, lNum := range leftNums {
			totalDiff += lNum * rightNumOccurrences[lNum]
		}
		fmt.Println("part two total:", totalDiff)
	}
}
