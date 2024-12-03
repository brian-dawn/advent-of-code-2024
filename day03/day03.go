package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {

	var part1Total, part2Total int
	multiplyEnabled := true

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		line := scanner.Text()

		re := regexp.MustCompile(`don't\(\)|do\(\)|mul\((\d+),(\d+)\)`)

		matches := re.FindAllStringSubmatch(line, -1)
		if matches == nil {
			continue
		}

		for _, match := range matches {

			if match[0] == "do()" {
				multiplyEnabled = true
				continue
			}

			if match[0] == "don't()" {
				multiplyEnabled = false
				continue
			}

			a, err := strconv.Atoi(match[1])
			if err != nil {
				continue
			}

			b, err := strconv.Atoi(match[2])
			if err != nil {
				continue
			}

			part1Total += a * b
			if multiplyEnabled {
				part2Total += a * b
			}

		}
	}

	fmt.Println("Part1:", part1Total)
	fmt.Println("Part2:", part2Total)
}
