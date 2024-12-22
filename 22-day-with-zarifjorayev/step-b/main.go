package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

var DIRS = [4][2]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}

func mix(x, y int) int {
	return x ^ y
}

func prune(x int) int {
	return x % 16777216
}

func prices(x int) []int {
	ans := []int{x}
	for i := 0; i < 2000; i++ {
		x = prune(mix(x, 64*x))
		x = prune(mix(x, x/32))
		x = prune(mix(x, x*2048))
		ans = append(ans, x)
	}
	return ans
}

func changes(P []int) []int {
	C := make([]int, len(P)-1)
	for i := 0; i < len(P)-1; i++ {
		C[i] = P[i+1] - P[i]
	}
	return C
}

func getScores(P, C []int) map[[4]int]int {
	ANS := make(map[[4]int]int)
	for i := 0; i < len(C)-3; i++ {
		pattern := [4]int{C[i], C[i+1], C[i+2], C[i+3]}
		if _, exists := ANS[pattern]; !exists {
			ANS[pattern] = P[i+4]
		}
	}
	return ANS
}

func main() {
	infile := "input.txt"
	if len(os.Args) >= 2 {
		infile = os.Args[1]
	}

	file, err := os.Open(infile)
	if err != nil {
		fmt.Printf("Error opening file %s: %v\n", infile, err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var D []string
	for scanner.Scan() {
		D = append(D, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		return
	}

	p1 := 0
	SCORE := make(map[[4]int]int)

	for _, line := range D {
		lineInt, err := strconv.Atoi(strings.TrimSpace(line))
		if err != nil {
			panic(err)
		}
		P := prices(lineInt)
		p1 += P[len(P)-1]

		for i := range P {
			P[i] %= 10
		}

		C := changes(P)
		S := getScores(P, C)
		for k, v := range S {
			if _, exists := SCORE[k]; !exists {
				SCORE[k] = v
			} else {
				SCORE[k] += v
			}
		}
	}

	fmt.Println(p1)

	maxScore := math.MinInt
	for _, v := range SCORE {
		if v > maxScore {
			maxScore = v
		}
	}

	fmt.Println(maxScore)
}
