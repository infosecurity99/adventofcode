package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	var lines [][]int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.Fields(scanner.Text())
		if len(line) > 0 {
			var nums []int
			for _, s := range line {
				n, err := strconv.Atoi(s)
				if err != nil {
					panic(err)
				}
				nums = append(nums, n)
			}
			lines = append(lines, nums)
		}
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}
	transposed := transpose(lines)
	sum := 0
	for _, nums := range transposed {
		sort.Ints(nums)
		for i := 0; i < len(nums)/2; i++ {
			sum += int(math.Abs(float64(nums[len(nums)-1-i] - nums[i])))
		}
	}
	fmt.Println(sum)
}
func transpose(matrix [][]int) [][]int {
	if len(matrix) == 0 {
		return nil
	}
	rows, cols := len(matrix), len(matrix[0])
	result := make([][]int, cols)
	for i := range result {
		result[i] = make([]int, rows)
	}
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			result[j][i] = matrix[i][j]
		}
	}
	return result
}
