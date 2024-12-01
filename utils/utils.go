package utils

import (
	"bytes"
	"io"
	"os"
	"strings"
)

/*
Parse a file into a list of strings (lines)
*/
func GetLinesFromFile(inFile string) []string {
	f, err := os.Open(inFile)
	if err != nil {
		panic(err)
	}
	bytes, err := io.ReadAll(f)
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(bytes), "\n")
	return lines
}

/*
Parse a file into a rectangular grid, and get width and height.
*/
func GetByteArrayNoNewlinesFromFile(inFile string) ([]byte, int, int) {
	f, err := os.Open(inFile)
	if err != nil {
		panic(err)
	}
	byteArray, err := io.ReadAll(f)
	if err != nil {
		panic(err)
	}
	lines := bytes.Split(byteArray, []byte{'\n'})
	width := len(lines[0])
	height := len(lines)
	// decrement height if input file ends with an empty newline
	if len(lines[len(lines)-1]) == 0 {
		height--
	}
	data := bytes.ReplaceAll(byteArray, []byte{'\n'}, []byte{})
	return data, width, height
}

/*
Go's standard library Abs() function only does floats, which is annoying.
*/
func IntAbs(val int) int {
	if val > 0 {
		return val
	} else {
		return (val ^ -1) + 1
	}
}
