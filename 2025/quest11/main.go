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
	birds := NewBirds(lines)
	isPhaseOne := true
	for r := 0; r < rounds; r++ {
		continuePhaseOne := false
		if isPhaseOne && birds.PhaseOne(continuePhaseOne) {
			continue
		}
		isPhaseOne = false
		birds.PhaseTwo()
	}
	return birds.Checksum()
}

func part2(lines []string) int {
	birds := NewBirds(lines)
	isPhaseOne := true
	rounds := 0
	for ; !birds.AllEqual(); rounds++ {
		continuePhaseOne := false
		if isPhaseOne && birds.PhaseOne(continuePhaseOne) {
			continue
		}
		isPhaseOne = false
		birds.PhaseTwo()
	}
	return rounds
}

func part3(lines []string) int {
	return len(lines)
}

type Birds []int

func NewBirds(lines []string) Birds {
	birds := make([]int, len(lines))
	for i, line := range lines {
		birds[i] = utils.StrToInt(line)
	}
	return birds
}

func (b *Birds) Checksum() int {
	sum := 0
	for i := 0; i < len(*b); i++ {
		sum += (i + 1) * (*b)[i]
	}
	return sum
}

func (b *Birds) AllEqual() bool {
	first := (*b)[0]
	for i := 1; i < len(*b); i++ {
		if (*b)[i] != first {
			return false
		}
	}
	return true
}

func (b *Birds) PhaseOne(continuePhaseOne bool) bool {
	for i := 0; i < len(*b)-1; i++ {
		if (*b)[i] > (*b)[i+1] {
			(*b)[i]--
			(*b)[i+1]++
			continuePhaseOne = true
		}
	}
	return continuePhaseOne
}

func (b *Birds) PhaseTwo() {
	for i := 0; i < len(*b)-1; i++ {
		if (*b)[i] < (*b)[i+1] {
			(*b)[i]++
			(*b)[i+1]--
		}
	}
}
