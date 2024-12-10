package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func funcParse(x string) []int {
	re := regexp.MustCompile(`-?\d+`)
	matches := re.FindAllString(x, -1)
	result := make([]int, len(matches))
	for i, m := range matches {
		val, _ := strconv.Atoi(m)
		result[i] = val
	}
	return result
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	var dep []([]int)
	var updates [][]int

	scanner := bufio.NewScanner(file)
	section := 0
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			section++
			continue
		}
		if section == 0 {
			dep = append(dep, funcParse(line))
		} else {
			updates = append(updates, funcParse(line))
		}
	}

	d := make(map[int]map[int]struct{})
	for _, pair := range dep {
		from := pair[0]
		to := pair[1]

		if _, exists := d[to]; !exists {
			d[to] = make(map[int]struct{})
		}
		d[to][from] = struct{}{}
	}

	c := 0
	for _, update := range updates {
		f := true
		for x := 0; x < len(update); x++ {
			if _, exists := d[update[x]]; exists {
				for _, depPage := range update[x+1:] {
					if _, found := d[update[x]][depPage]; found {
						f = false
						break
					}
				}
			}
			if !f {
				break
			}
		}
		if f {
			c += update[len(update)/2]
		}
	}

	fmt.Println("", c)

	c2 := 0
	for _, update := range updates {
		f := true
		for f {
			f = false
			for x := 0; x < len(update); x++ {
				for y := x + 1; y < len(update); y++ {
					from := update[x]
					to := update[y]
					if _, exists := d[from][to]; !exists {
						f = true
						update = append([]int{from}, update...)
						break
					}
				}
				if f {
					break
				}
			}
		}
		for _, page := range update {
			c2 += page
		}
	}

	fmt.Println(":", c2-c)
}
