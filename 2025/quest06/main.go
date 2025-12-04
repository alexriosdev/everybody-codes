package main

import (
	"everybody-codes/utils"
	"fmt"
)

func main() {
	line1, _ := utils.ReadLine("2025/quest06/input1.txt")
	lines2, _ := utils.ReadLines("2025/quest06/input2.txt")
	lines3, _ := utils.ReadLines("2025/quest06/input3.txt")
	fmt.Println("2025 Quest 06 Solution")
	fmt.Printf("Part 1: %v\n", part1(line1))
	fmt.Printf("Part 2: %v\n", part2(lines2))
	fmt.Printf("Part 3: %v\n", part3(lines3))
}

func part1(line string) int {
	sword := []rune{}
	for _, c := range line {
		if c == 'a' || c == 'A' {
			sword = append(sword, c)
		}
	}
	count := 0
	for i := 0; i < len(sword); i++ {
		if sword[i] == 'A' {
			for j := i + 1; j < len(sword); j++ {
				if sword[j] == 'a' {
					count++
				}
			}
		}
	}
	return count
}

func part2(lines []string) int {
	return len(lines)
}

func part3(lines []string) int {
	return len(lines)
}
