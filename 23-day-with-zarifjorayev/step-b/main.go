package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

type Graph map[string]map[string]bool

func addEdge(graph Graph, a, b string) {
	if graph[a] == nil {
		graph[a] = make(map[string]bool)
	}
	if graph[b] == nil {
		graph[b] = make(map[string]bool)
	}
	graph[a][b] = true
	graph[b][a] = true
}

func isClique(graph Graph, nodes []string) bool {
	for i := 0; i < len(nodes); i++ {
		for j := i + 1; j < len(nodes); j++ {
			if !graph[nodes[i]][nodes[j]] {
				return false
			}
		}
	}
	return true
}

func findMaxClique(graph Graph, nodes []string, clique []string, maxClique *[]string) {
	if len(clique) > len(*maxClique) {
		*maxClique = append([]string{}, clique...)
	}

	for i, node := range nodes {
		newClique := append(clique, node)
		if isClique(graph, newClique) {
			findMaxClique(graph, nodes[i+1:], newClique, maxClique)
		}
	}
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	graph := make(Graph)
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		nodes := strings.Split(line, "-")
		addEdge(graph, nodes[0], nodes[1])
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	var nodes []string
	for node := range graph {
		nodes = append(nodes, node)
	}

	var maxClique []string
	findMaxClique(graph, nodes, []string{}, &maxClique)

	sort.Strings(maxClique)
	fmt.Println("Password to the LAN party:", strings.Join(maxClique, ","))
}
