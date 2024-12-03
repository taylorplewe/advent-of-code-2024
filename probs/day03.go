package probs

import (
	"fmt"
	"io"
	"os"
	"regexp"
	"strconv"
)

func Day03(isPartTwo bool, inFile string) {
	file, _ := os.Open(inFile)
	defer file.Close()

	fileBytes, _ := io.ReadAll(file)
	fileStr := string(fileBytes)

	mulRe := regexp.MustCompile(`mul\((\d+),(\d+)\)`)
	matches := mulRe.FindAllStringSubmatch(fileStr, 10000)

	var sum int = 0
	for _, match := range matches {
		num1, _ := strconv.Atoi(match[1])
		num2, _ := strconv.Atoi(match[2])
		sum += num1 * num2
	}

	if isPartTwo {
		sum = 0
		doRe := regexp.MustCompile(`do\(\)`)
		dontRe := regexp.MustCompile(`don't\(\)`)
		combinedReText := fmt.Sprintf("%s|%s|%s", doRe.String(), dontRe.String(), mulRe.String())
		combinedRe := regexp.MustCompile(combinedReText)

		fileStrClip := fileStr
		nextMulInd := combinedRe.FindStringIndex(fileStrClip)
		mulOn := true
		for nextMulInd != nil {
			clip := fileStrClip[nextMulInd[0]:nextMulInd[1]]
			if doRe.MatchString(clip) {
				mulOn = true
			} else if dontRe.MatchString(clip) {
				mulOn = false
			} else if mulRe.MatchString(clip) && mulOn {
				m := mulRe.FindStringSubmatch(clip)
				num1, _ := strconv.Atoi(m[1])
				num2, _ := strconv.Atoi(m[2])
				sum += num1 * num2
			}
			fileStrClip = fileStrClip[nextMulInd[1]:]
			nextMulInd = combinedRe.FindStringIndex(fileStrClip)
		}
		fmt.Println("pt 2: ", sum)
		return
	}

	fmt.Println(sum)
}
