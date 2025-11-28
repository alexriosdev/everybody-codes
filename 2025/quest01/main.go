package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	input, _ := os.ReadFile("2025/quest01/input.txt")
	fmt.Println("2025 Quest 01 Solution")
	fmt.Printf("Part 1: %v\n", part1(input))
}

func part1(input []byte) string {
	split := strings.Split(strings.TrimSpace(string(input)), "\n\n")
	names := strings.Split(split[0], ",")
	instructions := strings.Split(split[1], ",")
	i, n := 0, len(names)-1
	for _, instruction := range instructions {
		dir, steps := instruction[0], StrToInt(instruction[1:])
		operand := 1
		if dir == 'L' {
			operand = -1
		}
		i = i + (steps * operand)
		if i < 0 {
			i = 0
		}
		if i > n {
			i = n
		}
	}
	return names[i]
}

func StrToInt(s string) int {
	num, _ := strconv.Atoi(s)
	return num
}
