package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type C struct{ x, y int }

type Config struct {
	A, B, P C
}

func parseConfigs(filename string) []Config {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return nil
	}
	defer file.Close()

	var configs []Config
	var currentConfig Config
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			if currentConfig != (Config{}) {
				configs = append(configs, currentConfig)
			}
			currentConfig = Config{}
		} else {
			var x, y int
			if strings.HasPrefix(line, "Button A") {
				fmt.Sscanf(line, "Button A: X+%d, Y+%d", &x, &y)
				currentConfig.A = C{x, y}
			}
			if strings.HasPrefix(line, "Button B") {
				fmt.Sscanf(line, "Button B: X+%d, Y+%d", &x, &y)
				currentConfig.B = C{x, y}
			}
			if strings.HasPrefix(line, "Prize") {
				fmt.Sscanf(line, "Prize: X=%d, Y=%d", &x, &y)
				const offset = 10000000000000
				currentConfig.P = C{x + offset, y + offset}
			}
		}
	}
	if currentConfig != (Config{}) {
		configs = append(configs, currentConfig)
	}

	return configs
}

func solve(c Config) int {
	b := (c.A.x*c.P.y - c.A.y*c.P.x) / (c.A.x*c.B.y - c.A.y*c.B.x)
	a := (c.P.x - b*c.B.x) / c.A.x

	if a*c.A.x+b*c.B.x == c.P.x && a*c.A.y+b*c.B.y == c.P.y {
		return 3*a + b
	}

	return 0
}

func main() {
	infile := "input.txt"
	if len(os.Args) >= 2 {
		infile = os.Args[1]
	}

	configs := parseConfigs(infile)

	if configs == nil {
		fmt.Println("No found.")
		return
	}

	tokens := 0
	for _, config := range configs {
		tokens += solve(config)
	}

	fmt.Println(" tokens:", tokens)
}
