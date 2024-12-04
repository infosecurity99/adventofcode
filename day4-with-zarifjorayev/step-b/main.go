package main

import (
	"bufio"
	"fmt"
	"os"
)

func readGridFromFile(filename string) ([]string, error) {
	var grid []string
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		grid = append(grid, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return grid, nil
}

func main() {
	filename := "grid.txt"

	grid, err := readGridFromFile(filename)
	if err != nil {
		fmt.Println("Error reading from file:", err)
		return
	}

	count := 0

	for r := 0; r < len(grid); r++ {
		for c := 0; c < len(grid[0]); c++ {
			if string(grid[r][c]) != "X" {
				continue
			}

			for _, dr := range []int{-1, 0, 1} {
				for _, dc := range []int{-1, 0, 1} {
					if dr == 0 && dc == 0 {
						continue
					}
					if r+3*dr < 0 || r+3*dr >= len(grid) || c+3*dc < 0 || c+3*dc >= len(grid[0]) {
						continue
					}

					if string(grid[r+dr][c+dc]) == "M" &&
						string(grid[r+2*dr][c+2*dc]) == "A" &&
						string(grid[r+3*dr][c+3*dc]) == "S" {
						count++
					}
				}
			}
		}
	}

	for r := 1; r < len(grid)-1; r++ {
		for c := 1; c < len(grid[0])-1; c++ {
			if string(grid[r][c]) != "A" {
				continue
			}

			corners := []string{
				string(grid[r-1][c-1]),
				string(grid[r-1][c+1]),
				string(grid[r+1][c+1]),
				string(grid[r+1][c-1]),
			}

			if corners[0] == "M" && corners[1] == "M" && corners[2] == "S" && corners[3] == "S" {
				count++
			} else if corners[0] == "M" && corners[1] == "S" && corners[2] == "S" && corners[3] == "M" {
				count++
			} else if corners[0] == "S" && corners[1] == "S" && corners[2] == "M" && corners[3] == "M" {
				count++
			} else if corners[0] == "S" && corners[1] == "M" && corners[2] == "M" && corners[3] == "S" {
				count++
			}
		}
	}

	fmt.Println(count)
}
