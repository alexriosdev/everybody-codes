package main

import (
	"everybody-codes/utils"
	"fmt"
	"strings"
)

func main() {
	line1, _ := utils.ReadLine("2025/quest16/input1.txt")
	line2, _ := utils.ReadLine("2025/quest16/input2.txt")
	line3, _ := utils.ReadLine("2025/quest16/input3.txt")
	fmt.Println("2025 Quest 16 Solution")
	fmt.Printf("Part 1: %v\n", part1(line1))
	fmt.Printf("Part 2: %v\n", part2(line2))
	fmt.Printf("Part 3: %v\n", part3(line3))
}

func part1(line string) int64 {
	spell := []int{}
	for _, s := range strings.Split(line, ",") {
		spell = append(spell, utils.StrToInt(s))
	}
	return countBlocks(spell, 90)
}

func part2(line string) int {
	wall := []int{}
	for _, s := range strings.Split(line, ",") {
		wall = append(wall, utils.StrToInt(s))
	}
	spell := getSpell(wall)
	result := 1
	for _, charm := range spell {
		result *= charm
	}
	return result
}

func part3(line string) int64 {
	target := int64(202520252025000)
	wall := []int{}
	for _, s := range strings.Split(line, ",") {
		wall = append(wall, utils.StrToInt(s))
	}
	spell := getSpell(wall)
	left, right := int64(0), target
	for left < right {
		mid := left + (right-left)/2
		blocks := countBlocks(spell, mid)
		if blocks <= target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return left
}

func countBlocks(spell []int, cols int64) int64 {
	sum := int64(0)
	for _, charm := range spell {
		sum += cols / int64(charm)
	}
	return sum
}

func getSpell(wall []int) []int {
	spell := []int{}
	for i := range wall {
		if wall[i] == 0 {
			continue
		}
		spell = append(spell, i+1)
		for j := i; j < len(wall); j += i + 1 {
			wall[j]--
		}
	}
	return spell
}
