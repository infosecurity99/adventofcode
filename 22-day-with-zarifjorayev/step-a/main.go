package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func nextSecret(secret int) int {

	secret ^= (secret * 64)
	secret %= 16777216

	secret ^= (secret / 32)
	secret %= 16777216

	secret ^= (secret * 2048)
	secret %= 16777216

	return secret
}

func main() {

	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("e:", err)
		return
	}
	defer file.Close()

	var initialSecrets []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}
		value, err := strconv.Atoi(line)
		if err != nil {
			fmt.Println("Error parsing number:", err)
			return
		}
		initialSecrets = append(initialSecrets, value)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	sumOf2000thSecrets := 0
	for _, secret := range initialSecrets {
		currentSecret := secret
		for i := 0; i < 2000; i++ {
			currentSecret = nextSecret(currentSecret)
		}
		sumOf2000thSecrets += currentSecret
	}

	fmt.Println("s:", sumOf2000thSecrets)
}
