package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func isRowLegal(nums []int) bool {
	diffs := make([]int, len(nums)-1)
	for i := 0; i < len(nums)-1; i++ {
		diffs[i] = nums[i+1] - nums[i]
	}

	for _, diff := range diffs {
		absDiff := int(math.Abs(float64(diff)))
		if absDiff < 1 || absDiff > 3 {
			return false
		}
	}

	isIncreasing := true
	isDecreasing := true
	for _, diff := range diffs {
		if diff <= 0 {
			isIncreasing = false
		}
		if diff >= 0 {
			isDecreasing = false
		}
	}

	return isIncreasing || isDecreasing
}

func part2(nums []int) bool {
	for i := 0; i < len(nums); i++ {
		copyNums := append([]int{}, nums[:i]...)
		copyNums = append(copyNums, nums[i+1:]...)
		if isRowLegal(copyNums) {
			return true
		}
	}
	return false
}

func main() {
	file, err := os.Open("input2.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lines := [][]int{}

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}
		parts := strings.Fields(line)
		nums := make([]int, len(parts))
		for i, part := range parts {
			num, err := strconv.Atoi(part)
			if err != nil {
				panic(err)
			}
			nums[i] = num
		}
		lines = append(lines, nums)
	}

	perfectlySafeReports := 0
	safeReports := 0

	for _, levels := range lines {
		if isRowLegal(levels) {
			perfectlySafeReports++
			safeReports++
		} else if part2(levels) {
			safeReports++
		}
	}

	fmt.Println("Perfectly Safe Reports:", perfectlySafeReports)
	fmt.Println("Safe Reports:", safeReports)
}
