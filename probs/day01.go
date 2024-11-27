package probs

import (
	"aoc/utils"
	"fmt"
	"strconv"
)

var elvesInv []int

func Day01(isDayTwo bool, inFile string) {
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
}
