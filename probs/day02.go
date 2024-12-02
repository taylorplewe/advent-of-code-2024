package probs

import (
	"aoc/utils"
	"fmt"
	"regexp"
	"strconv"
)

func Day02(isPartTwo bool, inFile string) {
	lines := utils.GetLinesFromFile(inFile)

	totalSafe := 0
	for _, line := range lines {
		r := regexp.MustCompile(`\b(\d+)\b`)
		numStrs := r.FindAllString(line, 20)
		fmt.Println(numStrs)
		if isLineSafe(numStrs) {
			totalSafe++
		} else if isPartTwo {
			for i := range numStrs {
				slicedNumStrs := make([]string, 0)
				for j := range numStrs {
					if j != i {
						slicedNumStrs = append(slicedNumStrs, numStrs[j])
					}
				}
				if isLineSafe(slicedNumStrs, true) {
					totalSafe++
					break
				}
			}
		}
	}

	fmt.Println("num safe:", totalSafe)
}

func isLineSafe(numStrs []string, verbose ...bool) bool {
	if len(verbose) > 0 {
		fmt.Println("sliced line:", numStrs)
	}
	increasing := false
	lastNum := -1
	for index, numStr := range numStrs {
		num, _ := strconv.Atoi(numStr)
		if lastNum == -1 {
			lastNum = num
			continue
		} else {
			if lastNum == num {
				return false
			}
			if index == 1 {
				if num > lastNum {
					increasing = true
				} else {
					increasing = false
				}
			} else {
				if increasing && num < lastNum {
					return false
				} else if !increasing && num > lastNum {
					return false
				}
			}
			if utils.IntAbs(num-lastNum) > 3 {
				return false
			}
			lastNum = num
		}
	}
	return true
}
