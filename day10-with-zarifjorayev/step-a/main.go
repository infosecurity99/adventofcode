package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func readMap(fileName string) [][]int {
	file, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var topoMap [][]int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		row := []int{}
		for _, c := range line {
			height, _ := strconv.Atoi(string(c))
			row = append(row, height)
		}
		topoMap = append(topoMap, row)
	}
	return topoMap
}

func isValid(x, y, rows, cols int) bool {
	return x >= 0 && x < rows && y >= 0 && y < cols
}

func findTrailScore(topoMap [][]int, startX, startY int) int {
	rows := len(topoMap)
	cols := len(topoMap[0])
	visited := make([][]bool, rows)
	for i := range visited {
		visited[i] = make([]bool, cols)
	}

	directions := [][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	stack := [][2]int{{startX, startY}}
	visited[startX][startY] = true
	reachable9s := 0

	for len(stack) > 0 {
		current := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		x, y := current[0], current[1]

		if topoMap[x][y] == 9 {
			reachable9s++
		}

		for _, dir := range directions {
			newX, newY := x+dir[0], y+dir[1]
			if isValid(newX, newY, rows, cols) &&
				!visited[newX][newY] &&
				topoMap[newX][newY] == topoMap[x][y]+1 {
				stack = append(stack, [2]int{newX, newY})
				visited[newX][newY] = true
			}
		}
	}

	return reachable9s
}

func main() {
	topoMap := readMap("input.txt")

	rows := len(topoMap)
	cols := len(topoMap[0])
	totalScore := 0

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if topoMap[i][j] == 0 {
				totalScore += findTrailScore(topoMap, i, j)
			}
		}
	}

	fmt.Printf(": %d\n", totalScore)
}
