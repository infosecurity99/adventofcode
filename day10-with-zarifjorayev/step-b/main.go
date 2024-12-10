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

// Dinamik dasturlash uchun yordamchi funksiya
func countDistinctTrails(topoMap [][]int, x, y int, memo [][]int) int {
	rows := len(topoMap)
	cols := len(topoMap[0])

	if memo[x][y] != -1 {
		return memo[x][y]
	}

	directions := [][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	trailCount := 0

	for _, dir := range directions {
		newX, newY := x+dir[0], y+dir[1]
		if isValid(newX, newY, rows, cols) && topoMap[newX][newY] == topoMap[x][y]+1 {
			trailCount += countDistinctTrails(topoMap, newX, newY, memo)
		}
	}

	// Agar 9 balandlikka yetgan bo'lsa, distinct yo‘lni birga qo‘shamiz
	if topoMap[x][y] == 9 {
		trailCount++
	}

	memo[x][y] = trailCount
	return trailCount
}

func main() {
	// Xarita ma'lumotlarini o'qiymiz
	topoMap := readMap("input.txt")

	rows := len(topoMap)
	cols := len(topoMap[0])
	totalRating := 0

	// Memoization massivini yaratamiz
	memo := make([][]int, rows)
	for i := range memo {
		memo[i] = make([]int, cols)
		for j := range memo[i] {
			memo[i][j] = -1
		}
	}

	// Trailheadlarni topib, ularning reytingini hisoblaymiz
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if topoMap[i][j] == 0 {
				totalRating += countDistinctTrails(topoMap, i, j, memo)
			}
		}
	}

	fmt.Printf("Trailheadlarning umumiy reytingi: %d\n", totalRating)
}
