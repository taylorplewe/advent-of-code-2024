package probs

import (
	"aoc/utils"
	"fmt"
	"slices"
	"strconv"
)

var elvesInv []int

func Day01(isPartTwo bool, inFile string) {
	lines := utils.GetLinesFromFile(inFile)
	elvesInv = make([]int, 0)
	elvesInv = append(elvesInv, 0)
	elfInd := 0
	highestInv := 0
	highestElfInd := 0

	for _, line := range lines {
		if len(line) == 0 {
			elfInd++
			elvesInv = append(elvesInv, 0)
		} else {
			num, _ := strconv.Atoi(line)
			elvesInv[elfInd] += num
			if elvesInv[elfInd] > highestInv {
				highestInv = elvesInv[elfInd]
				highestElfInd = elfInd
			}
		}
	}

	fmt.Println("highest inv elf:", highestElfInd+1)
	fmt.Println("highest inv", highestInv)

	if !isPartTwo {
		return
	}
	slices.Sort(elvesInv)
	slices.Reverse(elvesInv)
	fmt.Println("top three elves inv:", elvesInv[0]+elvesInv[1]+elvesInv[2])
}
