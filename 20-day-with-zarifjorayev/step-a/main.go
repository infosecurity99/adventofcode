package main

import (
	"fmt"
	"math"
)

func getDataLines() []string {
	return []string{
		".....",
		"...S.",
		"...#.",
		".#..E",
	}
}

func getNeighbors1(spot [2]int, myMap map[[2]int]rune) [][2]int {
	myx, myy := spot[0], spot[1]
	nbs := [][2]int{
		{myx - 1, myy},
		{myx, myy - 1},
		{myx, myy + 1},
		{myx + 1, myy},
	}
	var validNbs [][2]int
	for _, s := range nbs {
		if val, exists := myMap[s]; exists && (val == '.' || val == 'S' || val == 'E') {
			validNbs = append(validNbs, s)
		}
	}
	return validNbs
}

func bfs(start [2]int, myMap map[[2]int]rune, getNeighbors func([2]int, map[[2]int]rune) [][2]int) map[[2]int]int {
	distMap := make(map[[2]int]int)
	workq := [][2]int{start}
	distMap[start] = 0

	for len(workq) > 0 {
		spot := workq[0]
		workq = workq[1:]
		for _, n := range getNeighbors(spot, myMap) {
			if _, visited := distMap[n]; !visited {
				distMap[n] = distMap[spot] + 1
				workq = append(workq, n)
			}
		}
	}
	return distMap
}

func main() {
	data := getDataLines()
	mymap := make(map[[2]int]rune)
	var startloc, endloc [2]int
	for xidx, row := range data {
		for yidx, ch := range row {
			mymap[[2]int{xidx, yidx}] = ch
			if ch == 'S' {
				startloc = [2]int{xidx, yidx}
			} else if ch == 'E' {
				endloc = [2]int{xidx, yidx}
			}
		}
	}

	distToEndMap := bfs(endloc, mymap, getNeighbors1)

	distFromStartMap := bfs(startloc, mymap, getNeighbors1)

	baseDist := distToEndMap[startloc]

	goodcheats := make(map[[2]int][2]int)
	for h := range distFromStartMap {
		for _, n := range getNeighbors1(h, mymap) {
			newDist := distFromStartMap[h] + 2 + distToEndMap[n]
			if baseDist-newDist >= 100 {
				goodcheats[h] = n
			}
		}
	}

	fmt.Println("Part 1:", len(goodcheats))

	getNeighbors2 := func(spot [2]int, myMap map[[2]int]rune) [][2]int {
		myx, myy := spot[0], spot[1]
		var nbs [][2]int
		for xoff := -20; xoff <= 20; xoff++ {
			for yoff := -20; yoff <= 20; yoff++ {
				cheatlen := int(math.Abs(float64(xoff))) + int(math.Abs(float64(yoff)))
				if cheatlen <= 20 {
					nbs = append(nbs, [2]int{myx + xoff, myy + yoff})
				}
			}
		}
		var validNbs [][2]int
		for _, s := range nbs {
			if val, exists := myMap[s]; exists && (val == '.' || val == 'S' || val == 'E') {
				validNbs = append(validNbs, s)
			}
		}
		return validNbs
	}

	goodcheats = make(map[[2]int][2]int)
	for h := range distFromStartMap {
		for _, fn := range getNeighbors2(h, mymap) {
			newDist := distFromStartMap[h] + distToEndMap[fn] + 2
			if baseDist-newDist >= 100 {
				goodcheats[h] = fn
			}
		}
	}

	fmt.Println("Part 2:", len(goodcheats))
}
