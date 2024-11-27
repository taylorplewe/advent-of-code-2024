package utils

import (
	"bytes"
	"io"
	"os"
	"strings"
)

func GetLinesFromFile(inFile string) []string {
	f, err := os.Open(inFile)
	if err != nil {
		panic(err)
	}
	bytes, err := io.ReadAll(f)
	if err != nil {
		panic(err)
	}
	return strings.Split(string(bytes), "\n")
}

func GetByteArrayNoNewlinesFromFile(inFile string) []byte {
	f, err := os.Open(inFile)
	if err != nil {
		panic(err)
	}
	byteArray, err := io.ReadAll(f)
	if err != nil {
		panic(err)
	}
	return bytes.ReplaceAll(byteArray, []byte{'\n'}, []byte{})
}
