package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
)

const (
	RIGHT = iota
	DOWN
	LEFT
	UP
)

var DIRS = map[int][2]int{
	RIGHT: {1, 0},
	DOWN:  {0, 1},
	LEFT:  {-1, 0},
	UP:    {0, -1},
}

type State struct {
	score, x, y, direction int
}

type MinHeap []State

func (h *MinHeap) Len() int           { return len(*h) }
func (h *MinHeap) Less(i, j int) bool { return (*h)[i].score < (*h)[j].score }
func (h *MinHeap) Swap(i, j int)      { (*h)[i], (*h)[j] = (*h)[j], (*h)[i] }

func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(State))
}

func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	item := old[n-1]
	*h = old[0 : n-1]
	return item
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	defer file.Close()

	var grid []string
	var start, end [2]int

	scanner := bufio.NewScanner(file)
	for y := 0; scanner.Scan(); y++ {
		line := scanner.Text()
		grid = append(grid, line)
		for x, c := range line {
			if c == 'S' {
				start = [2]int{x, y}
			} else if c == 'E' {
				end = [2]int{x, y}
			}
		}
	}

	// Part 1
	todo := &MinHeap{}
	heap.Init(todo)
	heap.Push(todo, State{0, start[0], start[1], RIGHT})

	seen := make(map[[3]int]bool)
	origins := make(map[State][]State)
	var lowestScore int

	for len(*todo) > 0 {
		current := heap.Pop(todo).(State)

		if lowestScore != 0 && current.score > lowestScore {
			break
		}

		if current.x == end[0] && current.y == end[1] {
			lowestScore = current.score
			continue
		}

		stateKey := [3]int{current.x, current.y, current.direction}
		if seen[stateKey] {
			continue
		}
		seen[stateKey] = true

		for i := 1; i < 4; i++ {
			nd := (current.direction + i) % 4
			newState := State{current.score + 1000, current.x, current.y, nd}
			heap.Push(todo, newState)
			origins[newState] = append(origins[newState], current)
		}

		dx, dy := DIRS[current.direction][0], DIRS[current.direction][1]
		nx, ny := current.x+dx, current.y+dy
		if grid[ny][nx] != '#' {
			newState := State{current.score + 1, nx, ny, current.direction}
			heap.Push(todo, newState)
			origins[newState] = append(origins[newState], current)
		}
	}

	fmt.Println("Part 1:", lowestScore)

	// Part 2
	good := make(map[[2]int]bool)
	todo = &MinHeap{}
	heap.Init(todo)
	for _, d := range []int{RIGHT, DOWN, LEFT, UP} {
		heap.Push(todo, State{lowestScore, end[0], end[1], d})
	}

	for len(*todo) > 0 {
		current := heap.Pop(todo).(State)
		good[[2]int{current.x, current.y}] = true
		for _, prevState := range origins[current] {
			heap.Push(todo, prevState)
		}
	}

	fmt.Println("Part 2:", len(good))
}
