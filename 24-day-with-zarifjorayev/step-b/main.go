package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func Atoi(a string) int {
	i, err := strconv.Atoi(a)
	if err != nil {
		panic(err)
	}
	return i
}

func parseInput1(input string) (map[string]byte, [][4]string) {
	vals := make(map[string]byte)
	s := strings.Split(input, "\n\n")
	for _, s2 := range strings.Split(s[0], "\n") {
		s3 := strings.Split(s2, ": ")
		if s3[1] == "1" {
			vals[s3[0]] = 1
		} else {
			vals[s3[0]] = 0
		}
	}

	gates := make([][4]string, 0)
	for _, s2 := range strings.Split(s[1], "\n") {
		s3 := strings.Split(s2, " ")
		gates = append(gates, [4]string{s3[1], s3[0], s3[2], s3[4]})
	}

	return vals, gates
}

func p1(vals map[string]byte, gates [][4]string) (sum int) {
	for i := 0; i < len(gates); i++ {
		g := gates[i]

		x, ok := vals[g[1]]
		if !ok {
			gates = append(gates, g)
			continue
		}
		y, ok := vals[g[2]]
		if !ok {
			gates = append(gates, g)
			continue
		}
		if g[0] == "AND" {
			vals[g[3]] = x & y
		} else if g[0] == "OR" {
			vals[g[3]] = x | y
		} else if g[0] == "XOR" {
			vals[g[3]] = x ^ y
		}
	}

	for k, v := range vals {
		if k[0] == 'z' {
			sum += int(v) << Atoi(k[1:])
		}
	}

	return
}

type Connection struct {
	gate, to string
}

type Gate struct {
	op, a, b string
}

func parseInput2(input string) (map[string][]Connection, map[string]Gate) {
	s := strings.Split(input, "\n\n")

	connections := make(map[string][]Connection)
	gates := make(map[string]Gate)

	for _, s2 := range strings.Split(s[1], "\n") {
		s3 := strings.Split(s2, " ")
		connections[s3[0]] = append(connections[s3[0]], Connection{s3[1], s3[4]})
		connections[s3[2]] = append(connections[s3[2]], Connection{s3[1], s3[4]})

		gates[s3[4]] = Gate{s3[1], s3[0], s3[2]}
	}

	return connections, gates
}

func p2(conns map[string][]Connection, gates map[string]Gate) string {
	swaps := make([]string, 0)

	for i := 1; i <= 44; i++ {
		arg := "x"
		if i < 10 {
			arg += "0"
		}
		arg += strconv.Itoa(i)
		dest := "z" + arg[1:]

		g1 := conns[arg]
		if len(g1) != 2 {
			return fmt.Sprintf("Error: %v\n", g1)
		}

		var g2Name string
		if g1[0].gate == "XOR" {
			g2Name = g1[0].to
		}
		if g1[1].gate == "XOR" {
			g2Name = g1[1].to
		}
		g2 := conns[g2Name]

		if len(g2) != 2 {
			destGate := gates[dest]
			swaps = append(swaps, g2Name)
			if gates[destGate.a].op != "OR" {
				swaps = append(swaps, destGate.a)
			} else if gates[destGate.b].op != "OR" {
				swaps = append(swaps, destGate.b)
			} else {
				return fmt.Sprintf("Error: %v -> %v , %v\n", g1, gates[destGate.a], gates[destGate.b])
			}
			continue
		}

		if g2[0].gate == "XOR" && g2[0].to != dest {
			swaps = append(swaps, g2[0].to)
			swaps = append(swaps, dest)
		} else if g2[1].gate == "XOR" && g2[1].to != dest {
			swaps = append(swaps, g2[1].to)
			swaps = append(swaps, dest)
		}
	}

	slices.Sort(swaps)
	return strings.Join(swaps, ",")
}

func main() {
	input, _ := os.ReadFile("input.txt")
	inp := string(input)
	fmt.Printf("Part 1: %v\n", p1(parseInput1(inp)))
	fmt.Printf("Part 2: %v\n", p2(parseInput2(inp)))
}
