/*
Copyright © 2024 Taylor Plewe tplewe@outlook.com
*/
package cmd

import (
	"aoc/probs"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "aoc",
	Short: `Run an Advent of Code problem; just provide day number (e.g. "1") with optional ".2" for part 2 of the problem`,
	Run:   GetDayArgsAndExecute,
	Args:  cobra.RangeArgs(1, 2),
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

var dayFuncs = []func(bool, string){
	probs.Day01,
	probs.Day02,
	probs.Day03,
	probs.Day04,
	probs.Day05,
	probs.Day06,
	probs.Day07,
	probs.Day08,
	probs.Day09,
	probs.Day10,
	probs.Day11,
	probs.Day12,
	probs.Day13,
	probs.Day14,
	probs.Day15,
	probs.Day16,
	probs.Day17,
	probs.Day18,
	probs.Day19,
	probs.Day20,
	probs.Day21,
	probs.Day22,
	probs.Day23,
	probs.Day24,
	probs.Day25,
}

func GetDayArgsAndExecute(cmd *cobra.Command, args []string) {
	parts := strings.Split(args[0], ".")
	if len(parts) == 2 && parts[1] != "2" {
		fmt.Println(`to run part two, format must be "<day>.2"`)
		return
	}

	day, err := strconv.Atoi(parts[0])
	if err != nil {
		panic(err)
	}

	if day > len(dayFuncs) {
		fmt.Println("day provided is too high; there isn't a function for the provided day.")
		return
	}
	var filename string
	if len(args) == 2 {
		filename = fmt.Sprintf("inputs/%02d%s.txt", day, args[1])
	} else {
		filename = fmt.Sprintf("inputs/%02d.txt", day)
	}
	dayFuncs[day-1](len(parts) > 1, filename)
}
