package main

import (
	"bufio"
	"fmt"
	"os"
)

func pr(s interface{}) {
	fmt.Println(s)
}

func main() {
	infile := "input.txt"

	file, err := os.Open(infile)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	var G []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		G = append(G, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	R := len(G)
	C := len(G[0])

	var sr, sc int
	for r := 0; r < R; r++ {
		for c := 0; c < C; c++ {
			if G[r][c] == '^' {
				sr, sc = r, c
				break
			}
		}
	}

	var p1, p2 int

	for o_r := 0; o_r < R; o_r++ {
		for o_c := 0; o_c < C; o_c++ {
			r, c := sr, sc
			d := 0
			SEEN := make(map[[3]int]bool)
			SEEN_RC := make(map[[2]int]bool)
			for {
				if SEEN[[3]int{r, c, d}] {
					p2++
					break
				}
				SEEN[[3]int{r, c, d}] = true
				SEEN_RC[[2]int{r, c}] = true

				directions := [][2]int{
					{-1, 0}, {0, 1}, {1, 0}, {0, -1},
				}

				dr, dc := directions[d][0], directions[d][1]
				rr, cc := r+dr, c+dc

				if rr < 0 || rr >= R || cc < 0 || cc >= C {
					if G[o_r][o_c] == '#' {
						p1 = len(SEEN_RC)
					}
					break
				}

				if G[rr][cc] == '#' || (rr == o_r && cc == o_c) {
					d = (d + 1) % 4
				} else {
					r, c = rr, cc
				}
			}
		}
	}

	pr(p1)
	pr(p2)
}
