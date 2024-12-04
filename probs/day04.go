package probs

import (
	"aoc/utils"
	"fmt"
)

var grid []byte
var width, height int

func Day04(isPartTwo bool, inFile string) {
	grid, width, height = utils.GetByteArrayNoNewlinesFromFile(inFile)

	numXmas := 0
	for i, b := range grid {
		if !isPartTwo && b == 'X' {
			numXmas += searchForXmasFromLoc(i)
		} else if isPartTwo && b == 'A' {
			numXmas += searchForXmasFromLoc2(i)
		}
	}

	fmt.Println("num xmas:", numXmas)
}

func searchForXmasFromLoc(loc int) int {
	num := 0
	if loc%width > 2 {
		num += checkLocsForXmas(loc, -1, -2, -3)
		if loc/width > 2 {
			num += checkLocsForXmas(loc, -(width + 1), -(width*2 + 2), -(width*3 + 3))
		}
		if loc/width < height-3 {
			num += checkLocsForXmas(loc, width-1, width*2-2, width*3-3)
		}
	}
	if loc%width < width-3 {
		num += checkLocsForXmas(loc, 1, 2, 3)
		if loc/width > 2 {
			num += checkLocsForXmas(loc, -(width - 1), -(width*2 - 2), -(width*3 - 3))
		}
		if loc/width < height-3 {
			num += checkLocsForXmas(loc, width+1, width*2+2, width*3+3)
		}
	}
	if loc/width > 2 {
		num += checkLocsForXmas(loc, -width, -width*2, -width*3)
	}
	if loc/width < height-3 {
		num += checkLocsForXmas(loc, width, width*2, width*3)
	}
	return num
}

func checkLocsForXmas(orig int, loc1 int, loc2 int, loc3 int) int {
	if grid[orig+loc1] == 'M' && grid[orig+loc2] == 'A' && grid[orig+loc3] == 'S' {
		return 1
	} else {
		return 0
	}
}

func getXBytes(loc int) []byte {
	return []byte{
		grid[loc-(width+1)],
		grid[loc-(width-1)],
		grid[loc+(width-1)],
		grid[loc+width+1],
	}
}
func searchForXmasFromLoc2(loc int) int {
	if loc/width == 0 || loc/width == height-1 || loc%width == 0 || loc%width == width-1 {
		return 0
	}

	bs := getXBytes(loc)

	if bs[0] == 'M' && bs[1] == 'M' && bs[2] == 'S' && bs[3] == 'S' {
		return 1
	}
	if bs[0] == 'M' && bs[1] == 'S' && bs[2] == 'M' && bs[3] == 'S' {
		return 1
	}
	if bs[0] == 'S' && bs[1] == 'S' && bs[2] == 'M' && bs[3] == 'M' {
		return 1
	}
	if bs[0] == 'S' && bs[1] == 'M' && bs[2] == 'S' && bs[3] == 'M' {
		return 1
	}
	return 0
}
