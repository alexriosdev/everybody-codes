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
	r := 0
	for r < rounds {
		continuePhaseOne := false
		if isPhaseOne && birds.PhaseOne(&r, continuePhaseOne) {
			continue
		}
		isPhaseOne = false
		birds.PhaseTwo(&r)
	}
	return birds.Checksum()
}

func part2(lines []string) int {
	birds := NewBirds(lines)
	isPhaseOne := true
	rounds := 0
	for !birds.AllEqual() {
		continuePhaseOne := false
		if isPhaseOne && birds.PhaseOne(&rounds, continuePhaseOne) {
			continue
		}
		isPhaseOne = false
		birds.PhaseTwo(&rounds)
	}
	return rounds
}

func part3(lines []string) int {
	birds := NewBirds(lines)
	isPhaseOne := true
	rounds := 0
	for isPhaseOne {
		continuePhaseOne := false
		if !birds.PhaseOne(&rounds, continuePhaseOne) {
			isPhaseOne = false
		}
	}
	birds.PhaseTwoWithMath(&rounds)
	return rounds
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

func (b *Birds) PhaseOne(rounds *int, continuePhaseOne bool) bool {
	for i := 0; i < len(*b)-1; i++ {
		if (*b)[i] > (*b)[i+1] {
			(*b)[i]--
			(*b)[i+1]++
			continuePhaseOne = true
		}
	}
	if continuePhaseOne {
		*rounds++
	}
	return continuePhaseOne
}

func (b *Birds) PhaseTwo(rounds *int) {
	for i := 0; i < len(*b)-1; i++ {
		if (*b)[i] < (*b)[i+1] {
			(*b)[i]++
			(*b)[i+1]--
		}
	}
	*rounds++
}

func (b *Birds) PhaseTwoWithMath(rounds *int) {
	sum := 0
	for _, num := range *b {
		sum += num
	}
	mean := sum / len(*b)
	remainingRounds := 0
	for _, num := range *b {
		if num > mean {
			remainingRounds += num - mean
		}
	}
	*rounds += remainingRounds
}
