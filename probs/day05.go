package probs

import (
	"aoc/utils"
	"fmt"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

type Page struct {
	id          int
	pagesBefore []int
	pagesAfter  []int
}

var pages map[int]*Page
var allPageIds []int

func Day05(isPartTwo bool, inFile string) {
	lines := utils.GetLinesFromFile(inFile)
	pages = make(map[int]*Page)
	allPageIds = make([]int, 0)

	isInSecondSection := false
	sum := 0
	incorrectSum := 0
	lineNum := 0
	for _, line := range lines {
		if len(line) == 0 {
			if !isInSecondSection {
				isInSecondSection = true
				continue
			} else {
				break
			}
		}

		if isInSecondSection {
			updateIds := strings.Split(line, ",")
			oldUpdateIds := make([]int, 0)
			for _, idStr := range updateIds {
				id, _ := strconv.Atoi(idStr)
				oldUpdateIds = append(oldUpdateIds, id)
			}
			middleInd := len(updateIds) / 2
			newUpdateIds := make([]int, 0)
			// i := 0
			for _, id := range allPageIds {
				if slices.Index(oldUpdateIds, id) != -1 {
					newUpdateIds = append(newUpdateIds, id)
					// if i == middleInd {
					// 	incorrectSum += id
					// 	//break
					// }
					// i++
				}
			}
			fmt.Println(lineNum)
			fmt.Println(oldUpdateIds)
			fmt.Println(newUpdateIds)
			fmt.Println()
			if slices.Equal(oldUpdateIds, newUpdateIds) {
				sum += newUpdateIds[middleInd]
			} else {
				incorrectSum += newUpdateIds[middleInd]
			}
			//incorrectSum += newUpdateIds[middleInd]
			// }
		} else {
			pageRuleRe := regexp.MustCompile(`(\d+)\|(\d+)`)
			match := pageRuleRe.FindStringSubmatch(line)
			left, _ := strconv.Atoi(match[1])
			right, _ := strconv.Atoi(match[2])
			if _, exists := pages[left]; !exists {
				pages[left] = &Page{
					left,
					make([]int, 0),
					make([]int, 0),
				}
			}
			if _, exists := pages[right]; !exists {
				pages[right] = &Page{
					right,
					make([]int, 0),
					make([]int, 0),
				}
			}
			pages[left].pagesAfter = append(pages[left].pagesAfter, right)
			pages[right].pagesBefore = append(pages[right].pagesBefore, left)

			// try 2
			leftInd := slices.Index(allPageIds, left)
			rightInd := slices.Index(allPageIds, right)
			if leftInd == -1 {
				if rightInd == -1 {
					allPageIds = append(allPageIds, left, right)
				} else {
					allPageIds = slices.Insert(allPageIds, rightInd, left)
				}
			} else {
				if rightInd == -1 {
					allPageIds = slices.Insert(allPageIds, leftInd+1, right)
				} else if rightInd < leftInd {
					allPageIds = append(allPageIds[:leftInd], allPageIds[leftInd+1:]...)
					allPageIds = slices.Insert(allPageIds, rightInd, left)
				}
			}
		}
		lineNum++
	}

	fmt.Println(allPageIds)
	fmt.Println("sum:", sum)
	fmt.Println("incorrect sum:", incorrectSum)
}
