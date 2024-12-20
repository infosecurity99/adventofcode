package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

type Point struct {
	x, y int
}

var DIRS = []Point{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}

func readInput(filename string) (map[Point]byte, Point) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	coords := make(map[Point]byte)
	var start Point
	scanner := bufio.NewScanner(file)
	y := 0
	for scanner.Scan() {
		line := scanner.Text()
		for x, c := range line {
			point := Point{x, y}
			coords[point] = byte(c)
			if c == 'S' {
				start = point
			}
		}
		y++
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	return coords, start
}

func bfs(coords map[Point]byte, start Point) map[Point]int {
	dist := make(map[Point]int)
	queue := []Point{start}
	dist[start] = 0

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]
		currentDist := dist[current]

		for _, dir := range DIRS {
			next := Point{current.x + dir.x, current.y + dir.y}
			if _, visited := dist[next]; !visited && coords[next] != '#' {
				dist[next] = currentDist + 1
				queue = append(queue, next)
			}
		}
	}

	return dist
}

func calculateResults(coords map[Point]byte, dist map[Point]int) (int, int) {
	s1, s2 := 0, 0
	for c1, dist1 := range dist {
		for c2, dist2 := range dist {
			diff := int(math.Abs(float64(c2.x-c1.x)) + math.Abs(float64(c2.y-c1.y)))
			if dist2-dist1-diff >= 100 {
				if diff <= 2 {
					s1++
				}
				if diff <= 20 {
					s2++
				}
			}
		}
	}
	return s1, s2
}

func main() {
	infile := "input.txt"
	if len(os.Args) >= 2 {
		infile = os.Args[1]
	}

	coords, start := readInput(infile)
	dist := bfs(coords, start)
	s1, s2 := calculateResults(coords, dist)
	fmt.Println(s1)
	fmt.Println(s2)
}
