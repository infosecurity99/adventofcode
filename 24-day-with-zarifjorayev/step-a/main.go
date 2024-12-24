package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var wireValues = map[string]int{}

var gates = [][]string{}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error  :", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}
		parts := strings.Split(line, ": ")
		wire := parts[0]
		value, _ := strconv.Atoi(parts[1])
		wireValues[wire] = value
	}

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Fields(line)
		if len(parts) == 5 {
			gates = append(gates, parts)
		}
	}

	for len(gates) > 0 {
		pending := [][]string{}
		for _, gate := range gates {
			input1 := gate[0]
			op := gate[1]
			input2 := gate[2]
			output := gate[4]

			val1, ok1 := wireValues[input1]
			val2, ok2 := wireValues[input2]

			if ok1 && ok2 {
				switch op {
				case "AND":
					wireValues[output] = val1 & val2
				case "OR":
					wireValues[output] = val1 | val2
				case "XOR":
					wireValues[output] = val1 ^ val2
				}
			} else {
				pending = append(pending, gate)
			}
		}
		gates = pending
	}

	binaryOutput := ""
	for i := 0; ; i++ {
		wire := fmt.Sprintf("z%02d", i)
		val, ok := wireValues[wire]
		if !ok {
			break
		}
		binaryOutput = strconv.Itoa(val) + binaryOutput
	}

	decimalOutput, _ := strconv.ParseInt(binaryOutput, 2, 64)
	fmt.Println("l:", decimalOutput)
}
