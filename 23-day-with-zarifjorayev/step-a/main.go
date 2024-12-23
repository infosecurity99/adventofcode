package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func checkTriangle(graph map[string]map[string]bool, a, b, c string) bool {
	return graph[a][b] && graph[b][c] && graph[c][a]
}

func startsWithT(node string) bool {
	return strings.HasPrefix(node, "t")
}

func main() {

	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	graph := make(map[string]map[string]bool)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, "-")
		if len(parts) != 2 {
			fmt.Println("Invalid input format:", line)
			return
		}
		a, b := parts[0], parts[1]
		if graph[a] == nil {
			graph[a] = make(map[string]bool)
		}
		if graph[b] == nil {
			graph[b] = make(map[string]bool)
		}
		graph[a][b] = true
		graph[b][a] = true
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	nodes := make([]string, 0, len(graph))
	for node := range graph {
		nodes = append(nodes, node)
	}

	totalTriangles := 0
	tTriangles := 0
	for i := 0; i < len(nodes); i++ {
		for j := i + 1; j < len(nodes); j++ {
			for k := j + 1; k < len(nodes); k++ {
				a, b, c := nodes[i], nodes[j], nodes[k]
				if checkTriangle(graph, a, b, c) {
					totalTriangles++
					if startsWithT(a) || startsWithT(b) || startsWithT(c) {
						tTriangles++
					}
				}
			}
		}
	}

	fmt.Println("s:", totalTriangles)
	fmt.Println(" 't':", tTriangles)
}
