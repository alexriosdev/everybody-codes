package main

import (
	"everybody-codes/utils"
	"fmt"
)

func main() {
	lines1, _ := utils.ReadLines("2025/quest11/input1.txt")
	lines2, _ := utils.ReadLines("2025/quest11/input2.txt")
	lines3, _ := utils.ReadLines("2025/quest11/input3.txt")
	fmt.Println("2025 Quest 11 Solution")
	fmt.Printf("Part 1: %v\n", part1(lines1, 10))
	fmt.Printf("Part 2: %v\n", part2(lines2))
	fmt.Printf("Part 3: %v\n", part3(lines3))
}

func part1(lines []string, rounds int) int {
	birds := make([]int, len(lines))
	for i, line := range lines {
		birds[i] = utils.StrToInt(line)
	}
	isPhaseOne := true
	for r := 0; r < rounds; r++ {
		continuePhaseOne := false
		if isPhaseOne {
			for i := 0; i < len(birds)-1; i++ {
				if birds[i] > birds[i+1] {
					birds[i]--
					birds[i+1]++
					continuePhaseOne = true
				}
			}
			if continuePhaseOne {
				continue
			}
			isPhaseOne = false
		}
		for i := 0; i < len(birds)-1; i++ {
			if birds[i] < birds[i+1] {
				birds[i]++
				birds[i+1]--
			}
		}
	}
	checkSum := 0
	for i := 0; i < len(birds); i++ {
		checkSum += (i + 1) * birds[i]
	}
	return checkSum
}

func part2(lines []string) int {
	return len(lines)
}

func part3(lines []string) int {
	return len(lines)
}
