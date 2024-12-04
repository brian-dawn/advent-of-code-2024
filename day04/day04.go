package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type Direction struct {
	dx, dy int
}

func main() {

	var lines []string

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal("Error reading from stdin", err)
	}

	if len(lines) == 0 {
		log.Fatal("Need to pass input")
	}

	rows := len(lines)
	grid := make([][]byte, rows)
	for i := 0; i < rows; i++ {
		grid[i] = []byte(lines[i])
	}

	directions := []Direction{
		{0, 1},
		{0, -1},
		{1, 0},
		{-1, 0},
		{1, 1},
		{-1, -1},
		{1, -1},
		{-1, 1},
	}
	countXMAS := 0
	countMAS := 0
	for i := 0; i < rows; i++ {
		for j := 0; j < len(grid[i]); j++ {
			for _, dir := range directions {

				if findXMAS(grid, i, j, dir) {
					countXMAS++
				}

			}

			if findMAS(grid, i, j) {
				countMAS++
			}
		}
	}

	fmt.Println("Part1:", countXMAS)
	fmt.Println("Part2:", countMAS)
}

func findXMAS(grid [][]byte, startX, startY int, direction Direction) bool {
	rows := len(grid)
	cols := len(grid[0])

	for i := 0; i < len("XMAS"); i++ {

		x := startX + i*direction.dx
		y := startY + i*direction.dy

		if x < 0 || x >= rows || y < 0 || y >= cols {
			return false
		}

		if grid[x][y] != "XMAS"[i] {
			return false
		}
	}

	return true
}

func findMAS(grid [][]byte, startX, startY int) bool {

	rows := len(grid)
	cols := len(grid[0])

	if grid[startX][startY] != 'A' {
		return false
	}

	// Too close to boundaries.
	if startX == 0 || startX == cols-1 || startY == 0 || startY == rows-1 {
		return false
	}

	topLeft := grid[startX-1][startY-1]
	topRight := grid[startX+1][startY-1]

	botLeft := grid[startX-1][startY+1]
	botRight := grid[startX+1][startY+1]

	if topLeft == 'S' && topRight == 'S' && botLeft == 'M' && botRight == 'M' {
		return true
	}

	if topLeft == 'M' && topRight == 'M' && botLeft == 'S' && botRight == 'S' {
		return true
	}

	if topLeft == 'M' && topRight == 'S' && botLeft == 'M' && botRight == 'S' {
		return true
	}

	if topLeft == 'S' && topRight == 'M' && botLeft == 'S' && botRight == 'M' {
		return true
	}

	return false
}
