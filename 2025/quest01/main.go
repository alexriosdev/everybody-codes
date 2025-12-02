package main

import (
	"everybody-codes/utils"
	"fmt"
	"strings"
)

func main() {
	input1, _ := utils.ReadLines("2025/quest01/input1.txt")
	input2, _ := utils.ReadLines("2025/quest01/input2.txt")
	input3, _ := utils.ReadLines("2025/quest01/input3.txt")
	fmt.Println("2025 Quest 01 Solution")
	fmt.Printf("Part 1: %v\n", part1(input1))
	fmt.Printf("Part 2: %v\n", part2(input2))
	fmt.Printf("Part 3: %v\n", part3(input3))
}

func part1(lines []string) string {
	return navigateNames(lines, false, false)
}

func part2(lines []string) string {
	return navigateNames(lines, true, false)
}

func part3(lines []string) string {
	return navigateNames(lines, true, true)
}

func navigateNames(lines []string, isCircle, canSwap bool) string {
	names := strings.Split(lines[0], ",")
	instructions := strings.Split(lines[1], ",")
	i, n := 0, len(names)
	for _, instruction := range instructions {
		dir, steps := instruction[0], utils.StrToInt(instruction[1:])
		operand := 1
		if dir == 'L' {
			operand = -1
		}
		if !isCircle {
			i = i + (steps * operand)
			if i < 0 {
				i = 0
			}
			if i > n-1 {
				i = n - 1
			}
			continue
		}
		i = (i + (steps * operand)) % n
		if i < 0 {
			i += n
		}
		if canSwap {
			names[0], names[i] = names[i], names[0]
			i = 0
		}
	}
	return names[i]
}
