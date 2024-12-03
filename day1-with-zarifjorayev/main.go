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
	fileName := "input.txt"

	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Faylni ochishda xato:", err)
		return
	}
	defer file.Close()

	var data [][]int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		nums := []int{}
		for _, str := range strings.Fields(line) {
			num, _ := strconv.Atoi(str)
			nums = append(nums, num)
		}
		data = append(data, nums)
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("Faylni o'qishda xato:", err)
		return
	}

	if len(data) == 0 || len(data[0]) < 2 {
		fmt.Println("Faylda noto'g'ri yoki yetarli ma'lumot yo'q")
		return
	}

	rows := len(data)
	cols := len(data[0])
	columns := make([][]int, cols)
	for i := 0; i < cols; i++ {
		columns[i] = make([]int, rows)
		for j := 0; j < rows; j++ {
			columns[i][j] = data[j][i]
		}
	}

	for _, col := range columns {
		sort.Ints(col)
	}

	absDiffSum := 0
	for i := 0; i < len(columns[0]); i++ {
		absDiffSum += int(math.Abs(float64(columns[0][i] - columns[1][i])))
	}
	fmt.Println(":", absDiffSum)

	left, right := columns[0], columns[1]
	result := 0
	for _, x := range left {
		count := 0
		for _, y := range right {
			if x == y {
				count++
			}
		}
		result += x * count
	}
	fmt.Println("sum:", result)
}

/*package main

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
*/
