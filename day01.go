package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {

	var left, right []int

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		line := scanner.Text()

		fields := strings.Fields(line)

		if len(fields) != 2 {
			continue
		}

		leftNum, err := strconv.Atoi(fields[0])
		if err != nil {
			continue
		}

		rightNum, err := strconv.Atoi(fields[1])
		if err != nil {
			continue
		}

		left = append(left, leftNum)
		right = append(right, rightNum)

	}

	if err := scanner.Err(); err != nil {
		log.Fatal("Error reading input:", err)
	}

	if len(left) != len(right) {
		log.Fatal("Left and right have different lengths.")
	}

	sort.Ints(left)
	sort.Ints(right)

	totalDistance := 0
	for i := 0; i < len(left); i++ {
		diff := int(math.Abs(float64(left[i] - right[i])))
		totalDistance += diff
	}

	fmt.Println(totalDistance)
}
