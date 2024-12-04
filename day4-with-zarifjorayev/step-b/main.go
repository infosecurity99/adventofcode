package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fileName := "input.txt"
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	var l []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if line != "" {
			l = append(l, line)
		}
	}

	c := 0
	for i := 1; i < len(l)-1; i++ {
		for j := 1; j < len(l[i])-1; j++ {

			if string(l[i][j]) == "A" {

				if string(l[i-1][j-1]) == "M" && string(l[i+1][j+1]) == "M" &&
					string(l[i-1][j+1]) == "S" && string(l[i+1][j-1]) == "S" {
					c++
					fmt.Printf("Found X-MAS at: (%d, %d)\n", i, j)
				}
				if string(l[i-1][j+1]) == "M" && string(l[i+1][j-1]) == "M" &&
					string(l[i-1][j-1]) == "S" && string(l[i+1][j+1]) == "S" {
					c++
					fmt.Printf("Found X-MAS at: (%d, %d)\n", i, j)
				}
			}
		}
	}

	fmt.Printf("Total X-MAS matches: %d\n", c)
}
