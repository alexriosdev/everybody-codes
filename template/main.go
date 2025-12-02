package main

import (
	"everybody-codes/utils"
	"fmt"
)

func main() {
	lines1, _ := utils.ReadLines("202#/quest##/input1.txt")
	lines2, _ := utils.ReadLines("202#/quest##/input2.txt")
	lines3, _ := utils.ReadLines("202#/quest##/input3.txt")
	fmt.Println("202# Quest ## Solution")
	fmt.Printf("Part 1: %v\n", part1(lines1))
	fmt.Printf("Part 2: %v\n", part2(lines2))
	fmt.Printf("Part 3: %v\n", part3(lines3))
}

func part1(lines []string) int {
	return len(lines)
}

func part2(lines []string) int {
	return len(lines)
}

func part3(lines []string) int {
	return len(lines)
}
