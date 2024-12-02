package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	increasing = iota
	decreasing
)

func isSafeWithDampener(report []int) bool {
	// If it's already safe no need to dampen it.
	if isSafe(report) {
		return true
	}

	for i := 0; i < len(report); i++ {
		// Create a new report with the ith element removed.
		modified := append([]int{}, report[:i]...)
		modified = append(modified, report[i+1:]...)

		if isSafe(modified) {
			return true
		}

	}

	return false

}

func isSafe(report []int) bool {

	increasing := true
	decreasing := true

	for i := 1; i < len(report); i++ {

		diff := report[i] - report[i-1]

		if diff < -3 || diff > 3 || diff == 0 {
			return false
		}

		if diff > 0 {
			decreasing = false
		}

		if diff < 0 {
			increasing = false
		}
	}

	return increasing || decreasing // Must be one of the two
}

func main() {

	scanner := bufio.NewScanner(os.Stdin)

	var safeCounterWithDampener, safeCounter int

	for scanner.Scan() {
		line := scanner.Text()

		fields := strings.Fields(line)

		var level []int
		for _, field := range fields {

			number, err := strconv.Atoi(field)
			if err != nil {
				continue
			}

			level = append(level, number)
		}

		if isSafe(level) {
			safeCounter++
		}

		if isSafeWithDampener(level) {
			safeCounterWithDampener++
		}

	}

	fmt.Println("Part1:", safeCounter)
	fmt.Println("Part2:", safeCounterWithDampener)

}
