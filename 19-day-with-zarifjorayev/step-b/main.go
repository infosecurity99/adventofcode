package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	scanner.Scan()
	towelPatterns := strings.Split(scanner.Text(), ", ")

	scanner.Scan()

	var designs []string
	for scanner.Scan() {
		designs = append(designs, scanner.Text())
	}

	totalWays := 0
	for _, design := range designs {
		totalWays += countWaysToCreateDesign(design, towelPatterns)
	}

	fmt.Println(totalWays)
}

func countWaysToCreateDesign(design string, towelPatterns []string) int {
	dp := make([]int, len(design)+1)
	dp[0] = 1

	for i := 1; i <= len(design); i++ {
		for _, pattern := range towelPatterns {

			if i >= len(pattern) && design[i-len(pattern):i] == pattern {
				dp[i] += dp[i-len(pattern)]
			}
		}
	}

	return dp[len(design)]
}
