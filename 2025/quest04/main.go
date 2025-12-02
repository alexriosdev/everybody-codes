package main

import (
	"everybody-codes/utils"
	"fmt"
	"math"
	"strings"
)

func main() {
	lines1, _ := utils.ReadLines("2025/quest04/input1.txt")
	lines2, _ := utils.ReadLines("2025/quest04/input2.txt")
	lines3, _ := utils.ReadLines("2025/quest04/input3.txt")
	fmt.Println("2025 Quest 04 Solution")
	fmt.Printf("Part 1: %v\n", part1(lines1))
	fmt.Printf("Part 2: %v\n", part2(lines2))
	fmt.Printf("Part 3: %v\n", part3(lines3))
}

func part1(lines []string) int {
	first, last := utils.StrToFloat(lines[0]), utils.StrToFloat(lines[len(lines)-1])
	return int((first / last) * 2025)
}

func part2(lines []string) int {
	first, last := utils.StrToFloat(lines[0]), utils.StrToFloat(lines[len(lines)-1])
	return int(math.Ceil((last * 10_000_000_000_000) / first))
}

func part3(lines []string) int {
	first, last := utils.StrToFloat(lines[0]), utils.StrToFloat(lines[len(lines)-1])
	for _, s := range lines[1 : len(lines)-1] {
		innerSplit := strings.Split(s, "|")
		left, right := utils.StrToFloat(innerSplit[0]), utils.StrToFloat(innerSplit[1])
		first *= right / left
	}
	return int((first * 100) / last)
}
