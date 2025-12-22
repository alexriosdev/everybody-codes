package main

import (
	"everybody-codes/utils"
	"fmt"
)

func main() {
	lines1, _ := utils.ReadLines("2025/quest13/input1.txt")
	lines2, _ := utils.ReadLines("2025/quest13/input2.txt")
	lines3, _ := utils.ReadLines("2025/quest13/input3.txt")
	fmt.Println("2025 Quest 13 Solution")
	fmt.Printf("Part 1: %v\n", part1(lines1))
	fmt.Printf("Part 2: %v\n", part2(lines2))
	fmt.Printf("Part 3: %v\n", part3(lines3))
}

func part1(lines []string) int {
	dial := []int{1}
	for i, line := range lines {
		if i%2 == 0 {
			dial = append(dial, utils.StrToInt(line))
		}
	}
	for i := len(lines) - 1; i >= 0; i-- {
		if i%2 != 0 {
			dial = append(dial, utils.StrToInt(lines[i]))
		}
	}
	return dial[2025%len(dial)]
}

func part2(lines []string) int {
	return len(lines)
}

func part3(lines []string) int {
	return len(lines)
}
