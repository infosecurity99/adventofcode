package main

import (
	"bufio"
	"fmt"
	"os"
)

// Function to check if the X-MAS pattern exists in the grid at a specific position
func checkPattern(grid [][]rune, i, j, n, m int) bool {
	// Check for the X-MAS pattern in four diagonal directions
	return ((i-1 >= 0 && j-1 >= 0 && grid[i-1][j-1] == 'M' && grid[i][j] == 'S' && grid[i+1][j+1] == 'M') || // top-left to bottom-right
		(i-1 >= 0 && j+1 < m && grid[i-1][j+1] == 'M' && grid[i][j] == 'S' && grid[i+1][j-1] == 'M')) // top-right to bottom-left
}

func countXmasPatterns(grid [][]rune) int {
	// Get grid dimensions
	n, m := len(grid), len(grid[0])
	count := 0

	// Loop through the grid to find "M" positions that can be the center of an X-MAS
	for i := 1; i < n-1; i++ {
		for j := 1; j < m-1; j++ {
			if grid[i][j] == 'S' && checkPattern(grid, i, j, n, m) {
				count++
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
	// Change this to the correct file path for your input file
	filename := "input.txt"

	// Read the grid from the file
	grid, err := readGridFromFile(filename)
	if err != nil {
		fmt.Println("Error reading the file:", err)
		return
	}

	// Count the X-MAS patterns
	result := countXmasPatterns(grid)

	// Print the result
	fmt.Printf("The number of X-MAS patterns is: %d\n", result)
}
