package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func pr(s string) {
	fmt.Println(s)
}

func solve(part2 bool, D string) int {
	var A []struct{ pos, size, fileID int }
	var SPACE []struct{ pos, size int }
	var FINAL []interface{}
	pos := 0
	fileID := 0

	for i := 0; i < len(D); i++ {
		c := string(D[i])
		if i%2 == 0 {
			fileSize, _ := strconv.Atoi(c)
			if part2 {
				A = append(A, struct{ pos, size, fileID int }{pos, fileSize, fileID})
			}
			for j := 0; j < fileSize; j++ {
				FINAL = append(FINAL, fileID)
				if !part2 {
					A = append(A, struct{ pos, size, fileID int }{pos, 1, fileID})
				}
				pos++
			}
			fileID++
		} else {
			freeSize, _ := strconv.Atoi(c)
			SPACE = append(SPACE, struct{ pos, size int }{pos, freeSize})
			for j := 0; j < freeSize; j++ {
				FINAL = append(FINAL, nil)
				pos++
			}
		}
	}

	for i := len(A) - 1; i >= 0; i-- {
		block := A[i]
		for j := 0; j < len(SPACE); j++ {
			space := SPACE[j]
			if space.pos < block.pos && block.size <= space.size {
				for k := 0; k < block.size; k++ {
					if FINAL[block.pos+k] == block.fileID {
						FINAL[block.pos+k] = nil
						FINAL[space.pos+k] = block.fileID
					}
				}
				SPACE[j] = struct{ pos, size int }{space.pos + block.size, space.size - block.size}
				break
			}
		}
	}

	ans := 0
	for i, c := range FINAL {
		if c != nil {
			fileID, ok := c.(int)
			if ok {
				ans += i * fileID
			}
		}
	}
	return ans
}

func main() {
	infile := "input.txt"
	if len(os.Args) >= 2 {
		infile = os.Args[1]
	}
	file, err := os.Open(infile)
	if err != nil {
		fmt.Println(":", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	D := scanner.Text()

	p1 := solve(false, D)
	p2 := solve(true, D)

	pr(fmt.Sprintf("%d", p1))
	pr(fmt.Sprintf("%d", p2))
}
