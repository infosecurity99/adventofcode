package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"slices"
	"strconv"
	"strings"
)

var inputFile = flag.String("inputFile", "input.txt", ".")

const maxX int = 101
const maxY int = 103

type coord struct {
	x, y int
}

type robot struct {
	pos, vel coord
}

func (r *robot) tick() {
	r.pos.x += r.vel.x
	r.pos.y += r.vel.y
	r.pos.x %= maxX
	r.pos.y %= maxY
	if r.pos.x < 0 {
		r.pos.x += maxX
	}
	if r.pos.y < 0 {
		r.pos.y += maxY
	}
}

func clear(positions map[coord]int) {
	for k := range positions {
		delete(positions, k)
	}
}

func main() {
	flag.Parse()
	bytes, err := ioutil.ReadFile(*inputFile)
	if err != nil {
		return
	}
	contents := string(bytes)
	split := strings.Split(contents, "\n")

	var robots []robot
	for _, s := range split[:len(split)-1] {
		var r robot
		parts := strings.Split(s, " ")
		pos := strings.Split(parts[0][2:], ",")
		vel := strings.Split(parts[1][2:], ",")
		r.pos.x, _ = strconv.Atoi(pos[0])
		r.pos.y, _ = strconv.Atoi(pos[1])
		r.vel.x, _ = strconv.Atoi(vel[0])
		r.vel.y, _ = strconv.Atoi(vel[1])
		robots = append(robots, r)
	}

	var ul, ur, dl, dr int
	robotsA := slices.Clone(robots)
	for _, r := range robotsA {
		for i := 0; i < 100; i++ {
			r.tick()
		}
		if r.pos.x == maxX/2 || r.pos.y == maxY/2 {
			continue
		}
		if r.pos.x < maxX/2 {
			if r.pos.y < maxY/2 {
				ul++
			} else {
				dl++
			}
		} else {
			if r.pos.y < maxY/2 {
				ur++
			} else {
				dr++
			}
		}
	}
	fmt.Println("Safety Factor:", ul*ur*dl*dr)

	positions := make(map[coord]int)
	for i := 1; ; i++ {
		clear(positions)
		for i := range robots {
			robots[i].tick()
			positions[robots[i].pos]++
		}

	outer:
		for loc := range positions {
			for xDelta := -1; xDelta <= 1; xDelta++ {
				for yDelta := -1; yDelta <= 1; yDelta++ {
					if positions[coord{loc.x + xDelta, loc.y + yDelta}] == 0 {
						continue outer
					}
				}
			}
			fmt.Println(" time:", i)
			visualise(positions)
			return
		}
	}
}

func visualise(positions map[coord]int) {
	for y := 0; y < maxY; y++ {
		for x := 0; x < maxX; x++ {
			if positions[coord{x, y}] != 0 {
				fmt.Printf("*")
			} else {
				fmt.Printf(" ")
			}
		}
		fmt.Println()
	}
	fmt.Println()
}
