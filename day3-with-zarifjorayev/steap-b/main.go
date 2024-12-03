package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
	"strconv"
)

func main() {
	filename := "input.txt"
	input, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal("Error reading file:", err)
	}

	data := string(input)

	re := regexp.MustCompile(`mul\((\d+),(\d+)\)|do\(\)|don't\(\)`)

	matches := re.FindAllStringSubmatch(data, -1)

	total := 0
	enabled := true

	for _, match := range matches {
		if match[0] == "do()" {
			enabled = true
		} else if match[0] == "don't()" {
			enabled = false
		} else if enabled {
			x, errX := strconv.Atoi(match[1])
			y, errY := strconv.Atoi(match[2])

			if errX != nil || errY != nil {
				continue
			}

			total += x * y
		}
	}

	fmt.Println("The sum of the results of all enabled multiplications is:", total)
}
