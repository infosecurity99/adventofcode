package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fileName := "input.txt"

	file, err := os.Open(fileName)
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		return
	}
	defer file.Close()

	var l []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		l = append(l, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	c := 0
	for i := 0; i < len(l); i++ {
		for j := 0; j < len(l[i]); j++ {
			if strings.HasPrefix(l[i][j:], "XMAS") || strings.HasPrefix(l[i][j:], "SAMX") {
				c++
				fmt.Println(i, j)
			}
		}
	}

	lReversed := make([]string, len(l))
	for i := len(l) - 1; i >= 0; i-- {
		lReversed[len(l)-1-i] = l[i]
	}

	l2D := make([][]string, len(lReversed))
	for i := 0; i < len(lReversed); i++ {
		l2D[i] = strings.Split(lReversed[i], "")
	}

	for i := 0; i < len(l2D); i++ {
		for j := 0; j < len(l2D[i]); j++ {
			if strings.Join(l2D[i][j:], "") == "XMAS" || strings.Join(l2D[i][j:], "") == "SAMX" {
				c++
				fmt.Println(i, j)
			}
		}
	}

	for i := 0; i < len(l2D); i++ {
		for j := 0; j < len(l2D[i]); j++ {

			directions := [][2]int{
				{-1, -1},
				{1, 1},
				{-1, 1},
				{1, -1},
			}

			for _, dir := range directions {
				ci, cj := i, j
				s := ""
				for step := 0; step < 4; step++ {
					if ci >= 0 && ci < len(l2D) && cj >= 0 && cj < len(l2D[i]) {
						s += l2D[ci][cj]
						ci += dir[0]
						cj += dir[1]
					}
				}
				if s == "XMAS" {
					c++
					fmt.Println("diag", ci, cj)
				}
			}
		}
	}

	c2 := 0
	for i := 0; i < len(l2D); i++ {
		for j := 0; j < len(l2D[i]); j++ {
			if l2D[i][j] == "A" {
				r := []string{}
				for _, di := range []int{-1, 1} {
					for _, dj := range []int{-1, 1} {
						if i+di >= 0 && i+di < len(l2D) && j+dj >= 0 && j+dj < len(l2D[i]) {
							r = append(r, l2D[i+di][j+dj])
						}
					}
				}
				if strings.Count(strings.Join(r, ""), "M") == 2 && strings.Count(strings.Join(r, ""), "S") == 2 {
					if r[0] != r[len(r)-1] {
						c2++
					}
				}
			}
		}
	}

	fmt.Println("Count:", c)
	fmt.Println("Count2:", c2)
}
