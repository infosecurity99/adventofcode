package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf(" file: %v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lines := []string{}
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line != "" {
			lines = append(lines, line)
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatalf("e: %v", err)
	}

	if len(lines) < 4 {
		log.Fatalf("lines.")
	}

	A := parseRegister(lines[0], "Register A")
	B := parseRegister(lines[1], "Register B")
	C := parseRegister(lines[2], "Register C")

	if !strings.HasPrefix(lines[3], "Program:") {
		log.Fatalf(": '. Found: %s", lines[3])
	}
	programParts := strings.SplitN(lines[3], ": ", 2)
	if len(programParts) < 2 {
		log.Fatalf("'.")
	}
	programStr := strings.Split(programParts[1], ",")
	program := make([]int, len(programStr))
	for i, val := range programStr {
		val = strings.TrimSpace(val)
		program[i], err = strconv.Atoi(val)
		if err != nil {
			log.Fatalf("xn: %v", err)
		}
	}

	for A = 1; ; A++ {

		output := []string{}
		instructionPointer := 0

		for instructionPointer < len(program) {
			opcode := program[instructionPointer]
			if instructionPointer+1 >= len(program) {
				break
			}
			operand := program[instructionPointer+1]

			switch opcode {
			case 0:
				denominator := 1 << getComboOperandValue(operand, A, B, C)
				if denominator != 0 {
					A /= denominator
				} else {
					A = 0
				}
			case 1:
				B ^= operand
			case 2:
				B = getComboOperandValue(operand, A, B, C) % 8
			case 3:
				if A != 0 {
					instructionPointer = operand
					continue
				}
			case 4:
				B ^= C
			case 5:
				outputValue := getComboOperandValue(operand, A, B, C) % 8
				output = append(output, strconv.Itoa(outputValue))
			case 6:
				denominator := 1 << getComboOperandValue(operand, A, B, C)
				if denominator != 0 {
					B = A / denominator
				} else {
					B = 0
				}
			case 7:
				denominator := 1 << getComboOperandValue(operand, A, B, C)
				if denominator != 0 {
					C = A / denominator
				} else {
					C = 0
				}
			default:
				log.Fatalf("Invalid opcode: %d", opcode)
			}

			instructionPointer += 2
		}

		if strings.Join(output, ",") == strings.Join(programStr, ",") {
			fmt.Printf(" A is: %d\n", A)
			break
		}
	}
}

func parseRegister(line, label string) int {
	parts := strings.Split(line, ": ")
	if len(parts) != 2 || parts[0] != label {
		log.Fatalf("Invalid format for %s: %s", label, line)
	}
	value := strings.TrimSpace(parts[1])
	val, err := strconv.Atoi(value)
	if err != nil {
		log.Fatalf("Failed to parse %s: %v", label, err)
	}
	return val
}

func getComboOperandValue(operand int, A, B, C int) int {
	switch operand {
	case 0, 1, 2, 3:
		return operand
	case 4:
		return A
	case 5:
		return B
	case 6:
		return C
	default:
		log.Fatalf("Invalid combo operand: %d", operand)
		return 0
	}
}
