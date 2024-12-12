package main

import (
	"bufio"
	"container/list"
	"fmt"
	"log"
	"os"
)

var DIRS = [][2]int{
	{-1, 0}, // up
	{0, 1},  // right
	{1, 0},  // down
	{0, -1}, // left
}

func readInput(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return lines, nil
}

func main() {
	infile := "input.txt"
	if len(os.Args) > 1 {
		infile = os.Args[1]
	}

	// Step 1: Read input data
	grid, err := readInput(infile)
	if err != nil {
		log.Fatalf("Error reading input: %v", err)
	}

	R := len(grid)
	C := len(grid[0])

	// Ensure all rows are the same length
	for i, row := range grid {
		if len(row) != C {
			log.Fatalf("Inconsistent row length at row %d: expected %d, got %d", i, C, len(row))
		}
	}

	seen := make(map[[2]int]bool)
	p1, p2 := 0, 0

	// Step 2: Process the grid
	for r := 0; r < R; r++ {
		for c := 0; c < C; c++ {
			if seen[[2]int{r, c}] {
				continue
			}

			// Flood fill
			Q := list.New()
			Q.PushBack([2]int{r, c})
			area, perim := 0, 0
			perimMap := make(map[[2]int]map[[2]int]struct{})

			for Q.Len() > 0 {
				elem := Q.Front()
				Q.Remove(elem)
				cell := elem.Value.([2]int)
				r2, c2 := cell[0], cell[1]

				if seen[cell] {
					continue
				}
				seen[cell] = true
				area++

				for _, dir := range DIRS {
					rr, cc := r2+dir[0], c2+dir[1]
					if rr < 0 || rr >= R || cc < 0 || cc >= C || grid[rr][cc] != grid[r2][c2] {
						perim++
						if _, exists := perimMap[dir]; !exists {
							perimMap[dir] = make(map[[2]int]struct{})
						}
						perimMap[dir][cell] = struct{}{}
					} else {
						Q.PushBack([2]int{rr, cc})
					}
				}
			}

			// Calculate sides
			sides := 0
			for _, positions := range perimMap {
				seenPerim := make(map[[2]int]bool)
				for pos := range positions {
					if seenPerim[pos] {
						continue
					}
					sides++
					Q.PushBack(pos)
					for Q.Len() > 0 {
						elem := Q.Front()
						Q.Remove(elem)
						cell := elem.Value.([2]int)
						r2, c2 := cell[0], cell[1]

						if seenPerim[cell] {
							continue
						}
						seenPerim[cell] = true

						for _, dir := range DIRS {
							rr, cc := r2+dir[0], c2+dir[1]
							if _, exists := positions[[2]int{rr, cc}]; exists {
								Q.PushBack([2]int{rr, cc})
							}
						}
					}
				}
			}

			p1 += area * perim
			p2 += area * sides
		}
	}

	// Step 3: Output results
	fmt.Printf("p1: %d\n", p1)
	fmt.Printf("p2: %d\n", p2)
}
