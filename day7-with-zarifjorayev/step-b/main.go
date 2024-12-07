package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func evaluate(target int, numbers []int, current int, index int) bool {
	if index == len(numbers) {
		return current == target
	}

	if evaluate(target, numbers, current+numbers[index], index+1) {
		return true
	}

	if evaluate(target, numbers, current*numbers[index], index+1) {
		return true
	}

	concat := concatNumbers(current, numbers[index])
	if evaluate(target, numbers, concat, index+1) {
		return true
	}

	return false
}

func concatNumbers(a, b int) int {
	return a*powerOf10(b) + b
}

func powerOf10(n int) int {
	count := 1
	for n > 0 {
		count *= 10
		n /= 10
	}
	return count
}

func parseInput(filename string) ([][]int, []int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var targets []int
	var equations [][]int

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, ":")
		if len(parts) != 2 {
			return nil, nil, fmt.Errorf("invalid input format")
		}

		target, err := strconv.Atoi(strings.TrimSpace(parts[0]))
		if err != nil {
			return nil, nil, err
		}
		targets = append(targets, target)

		numberStrings := strings.Fields(parts[1])
		var numbers []int
		for _, ns := range numberStrings {
			num, err := strconv.Atoi(ns)
			if err != nil {
				return nil, nil, err
			}
			numbers = append(numbers, num)
		}
		equations = append(equations, numbers)
	}

	if err := scanner.Err(); err != nil {
		return nil, nil, err
	}

	return equations, targets, nil
}

func main() {
	equations, targets, err := parseInput("input.txt")
	if err != nil {
		fmt.Printf("Error reading input: %v\n", err)
		return
	}

	total := 0

	for i, numbers := range equations {
		if evaluate(targets[i], numbers, numbers[0], 1) {
			total += targets[i]
		}
	}

	fmt.Printf("Total calibration result: %d\n", total)
}
