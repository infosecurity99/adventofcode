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

	re := regexp.MustCompile(`mul\((\d+),(\d+)\)`)

	matches := re.FindAllStringSubmatch(data, -1)

	total := 0

	for _, match := range matches {
		x, errX := strconv.Atoi(match[1])
		y, errY := strconv.Atoi(match[2])

		if errX != nil || errY != nil {
			continue
		}

		total += x * y
	}

	fmt.Println("The sum of the results of all valid multiplications is:", total)
}
