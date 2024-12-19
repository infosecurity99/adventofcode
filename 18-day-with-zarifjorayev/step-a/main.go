package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const N = 71

var DIRS = [4][2]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}

func readInput(filename string) ([][2]int, error) {
	var positions [][2]int
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		coords := strings.Split(line, ",")
		if len(coords) != 2 {
			continue
		}
		x, _ := strconv.Atoi(coords[0])
		y, _ := strconv.Atoi(coords[1])
		positions = append(positions, [2]int{x, y})
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return positions, nil
}

func bfsShortestPath(G [][]bool, start, end [2]int) (int, bool) {
	type State struct {
		r, c, d int
	}

	queue := []State{{start[0], start[1], 0}}
	seen := make(map[[2]int]bool)
	seen[start] = true

	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]

		r, c, d := curr.r, curr.c, curr.d

		if r == end[0] && c == end[1] {
			return d, true
		}

		for _, dir := range DIRS {
			rr, cc := r+dir[0], c+dir[1]
			if rr >= 0 && rr < N && cc >= 0 && cc < N && !seen[[2]int{rr, cc}] && G[rr][cc] {
				seen[[2]int{rr, cc}] = true
				queue = append(queue, State{rr, cc, d + 1})
			}
		}
	}

	return -1, false
}

func isPathAvailable(G [][]bool, start, end [2]int) bool {
	_, found := bfsShortestPath(G, start, end)
	return found
}

func main() {
	infile := "input.txt"
	if len(os.Args) > 1 {
		infile = os.Args[1]
	}

	positions, err := readInput(infile)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	G := make([][]bool, N)
	for i := range G {
		G[i] = make([]bool, N)
		for j := range G[i] {
			G[i][j] = true
		}
	}

	start := [2]int{0, 0}
	end := [2]int{70, 70}

	for i, pos := range positions {
		x, y := pos[0], pos[1]
		if 0 <= y && y < N && 0 <= x && x < N {
			G[y][x] = false
		}

		if i == 1023 {
			if dist, found := bfsShortestPath(G, start, end); found {
				fmt.Println(dist)
			} else {
				fmt.Println("1024.")
			}
		}

		if !isPathAvailable(G, start, end) {
			fmt.Printf("%d,%d\n", x, y)
			break
		}
	}
}
