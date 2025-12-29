package main

import (
	"everybody-codes/utils"
	"fmt"
	"strings"
)

func main() {
	line1, _ := utils.ReadLine("2025/quest16/input1.txt")
	lines2, _ := utils.ReadLines("2025/quest16/input2.txt")
	lines3, _ := utils.ReadLines("2025/quest16/input3.txt")
	fmt.Println("2025 Quest 16 Solution")
	fmt.Printf("Part 1: %v\n", part1(line1))
	fmt.Printf("Part 2: %v\n", part2(lines2))
	fmt.Printf("Part 3: %v\n", part3(lines3))
}

func part1(line string) int {
	cols := 90
	sum := 0
	for _, s := range strings.Split(line, ",") {
		sum += cols / utils.StrToInt(s)
	}
	return sum
}

func part2(lines []string) int {
	return len(lines)
}

func part3(lines []string) int {
	return len(lines)
}
