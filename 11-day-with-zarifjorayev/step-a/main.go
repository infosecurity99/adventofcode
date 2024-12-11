package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var DP = make(map[string]int)

func solve(x int, t int) int {
	key := fmt.Sprintf("%d-%d", x, t)

	if val, found := DP[key]; found {
		return val
	}

	var ret int
	if t == 0 {
		ret = 1
	} else if x == 0 {
		ret = solve(1, t-1)
	} else if len(strconv.Itoa(x))%2 == 0 {
		strX := strconv.Itoa(x)
		mid := len(strX) / 2
		left, _ := strconv.Atoi(strX[:mid])
		right, _ := strconv.Atoi(strX[mid:])
		ret = solve(left, t-1) + solve(right, t-1)
	} else {
		ret = solve(x*2024, t-1)
	}

	DP[key] = ret
	return ret
}

func solveAll(t int, D []int) int {
	total := 0
	for _, x := range D {
		total += solve(x, t)
	}
	return total
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	var D []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Fields(line)
		for _, part := range parts {
			num, err := strconv.Atoi(part)
			if err != nil {
				fmt.Println("Error converting string to int:", err)
				continue
			}
			D = append(D, num)
		}
	}

	result25 := solveAll(25, D)
	result75 := solveAll(75, D)

	fmt.Println("Result after 25 blinks:", result25)
	fmt.Println("Result after 75 blinks:", result75)
}
