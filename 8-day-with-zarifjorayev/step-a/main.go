package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

func calculateAntinodes(grid []string) int {
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

				dy, dx := y2-y1, x2-x1
				midY, midX := y1-dy, x1-dx
				endY, endX := y2+dy, x2+dx

				if midY >= 0 && midY < len(grid) && midX >= 0 && midX < len(grid[0]) {
					antinodeSet[[2]int{midY, midX}] = struct{}{}
				}
				if endY >= 0 && endY < len(grid) && endX >= 0 && endX < len(grid[0]) {
					antinodeSet[[2]int{endY, endX}] = struct{}{}
				}
			}
		}
	}

	return len(antinodeSet)
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(":", err)
		return
	}
	defer file.Close()

	var grid []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		grid = append(grid, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(":", err)
		return
	}

	result := calculateAntinodes(grid)
	fmt.Println(":", result)
}
