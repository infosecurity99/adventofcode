package main

import (
	"bufio"
	"fmt"
	"os"
)

var directions = []struct {
	deltaX, deltaY int
}{
	{1, 0},
	{-1, 0},
	{0, 1},
	{0, -1},
	{1, 1},
	{-1, -1},
	{1, -1},
	{-1, 1},
}

func checkDirection(grid [][]rune, x, y, dx, dy int) bool {
	word := "X-MAS"
	for i := 0; i < len(word); i++ {
		newX, newY := x+dx*i, y+dy*i
		if newX < 0 || newX >= len(grid) || newY < 0 || newY >= len(grid[0]) || grid[newX][newY] != rune(word[i]) {
			return false
		}
	}
	return true
}
func countOccurrences(grid [][]rune) int {
	count := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			for _, dir := range directions {
				if checkDirection(grid, i, j, dir.deltaX, dir.deltaY) {
					count++
				}
			}
		}
	}
	return count
}

func readGridFromFile(filename string) ([][]rune, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var grid [][]rune
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		grid = append(grid, []rune(line))
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return grid, nil
}

func main() {
	filename := "input.txt"

	grid, err := readGridFromFile(filename)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	result := countOccurrences(grid)
	fmt.Printf("T", result)
}
