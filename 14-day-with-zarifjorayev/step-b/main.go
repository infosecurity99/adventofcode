package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

const X = 101
const Y = 103

var DIRS = [4][2]int{
	{-1, 0},
	{0, 1},
	{1, 0},
	{0, -1},
}

func ints(s string) []int {
	re := regexp.MustCompile("-?\\d+")
	matches := re.FindAllString(s, -1)
	var result []int
	for _, match := range matches {
		num, _ := strconv.Atoi(match)
		result = append(result, num)
	}
	return result
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var robots [][4]int

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}
		coords := ints(line)
		robots = append(robots, [4]int{coords[0], coords[1], coords[2], coords[3]})
	}

	var q1, q2, q3, q4 int
	var t int

	for t = 1; t < 1000000; t++ {
		var G [Y][X]rune
		for y := 0; y < Y; y++ {
			for x := 0; x < X; x++ {
				G[y][x] = '.'
			}
		}

		if t == 100 {
			q1, q2, q3, q4 = 0, 0, 0, 0
		}

		for i := 0; i < len(robots); i++ {
			px, py, vx, vy := robots[i][0], robots[i][1], robots[i][2], robots[i][3]
			px += vx
			py += vy
			px %= X
			py %= Y

			if px < 0 {
				px += X
			}
			if py < 0 {
				py += Y
			}

			robots[i] = [4]int{px, py, vx, vy}

			if px >= 0 && px < X && py >= 0 && py < Y {
				G[py][px] = '#'
			}

			if t == 100 {
				if px < X/2 && py < Y/2 {
					q1++
				}
				if px > X/2 && py < Y/2 {
					q2++
				}
				if px < X/2 && py > Y/2 {
					q3++
				}
				if px > X/2 && py > Y/2 {
					q4++
				}
			}
		}

		if t == 100 {
			fmt.Println(q1 * q2 * q3 * q4)
		}

		components := 0
		SEEN := make(map[[2]int]bool)

		for x := 0; x < X; x++ {
			for y := 0; y < Y; y++ {
				if G[y][x] == '#' && !SEEN[[2]int{x, y}] {
					sx, sy := x, y
					components++
					queue := [][2]int{{sx, sy}}
					for len(queue) > 0 {
						node := queue[0]
						queue = queue[1:]
						x2, y2 := node[0], node[1]
						if SEEN[[2]int{x2, y2}] {
							continue
						}
						SEEN[[2]int{x2, y2}] = true
						for _, dir := range DIRS {
							xx, yy := x2+dir[0], y2+dir[1]
							if xx >= 0 && xx < X && yy >= 0 && yy < Y && G[yy][xx] == '#' {
								queue = append(queue, [2]int{xx, yy})
							}
						}
					}
				}
			}
		}

		if components <= 200 {
			fmt.Println(t)
			for _, row := range G {
				fmt.Println(string(row[:]))
			}
			break
		}
	}
}
