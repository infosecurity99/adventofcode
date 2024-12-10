package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

func calculateAntinodesPartTwo(grid []string) int {
	antennaMap := make(map[rune][][2]int)
	for y, row := range grid {
		for x, char := range row {
			if unicode.IsLetter(char) || unicode.IsDigit(char) {
				antennaMap[char] = append(antennaMap[char], [2]int{y, x})
			}
		}
	}

	antinodeSet := make(map[[2]int]struct{})

	for _, positions := range antennaMap {
		n := len(positions)
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				if i == j {
					continue
				}

				y1, x1 := positions[i][0], positions[i][1]
				y2, x2 := positions[j][0], positions[j][1]

				antinodeSet[[2]int{y1, x1}] = struct{}{}
				antinodeSet[[2]int{y2, x2}] = struct{}{}

				dy, dx := y2-y1, x2-x1

				for step := 1; ; step++ {
					newY, newX := y1-step*dy, x1-step*dx
					if newY < 0 || newY >= len(grid) || newX < 0 || newX >= len(grid[0]) {
						break
					}
					antinodeSet[[2]int{newY, newX}] = struct{}{}
				}

				for step := 1; ; step++ {
					newY, newX := y2+step*dy, x2+step*dx
					if newY < 0 || newY >= len(grid) || newX < 0 || newX >= len(grid[0]) {
						break
					}
					antinodeSet[[2]int{newY, newX}] = struct{}{}
				}
			}
		}
	}

	return len(antinodeSet)
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening input file:", err)
		return
	}
	defer file.Close()

	var grid []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		grid = append(grid, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading input file:", err)
		return
	}

	result := calculateAntinodesPartTwo(grid)
	fmt.Println(":", result)
}
