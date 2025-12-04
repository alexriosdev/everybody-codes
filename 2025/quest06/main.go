package main

import (
	"everybody-codes/utils"
	"fmt"
)

func main() {
	line1, _ := utils.ReadLine("2025/quest06/input1.txt")
	line2, _ := utils.ReadLine("2025/quest06/input2.txt")
	lines3, _ := utils.ReadLines("2025/quest06/input3.txt")
	fmt.Println("2025 Quest 06 Solution")
	fmt.Printf("Part 1: %v\n", part1(line1))
	fmt.Printf("Part 2: %v\n", part2(line2))
	fmt.Printf("Part 3: %v\n", part3(lines3))
}

func part1(line string) int {
	categories := &Categories{}
	for _, c := range line {
		categories.Add(c)
	}
	return categories.SwordPairs()
}

func part2(line string) int {
	categories := &Categories{}
	for _, c := range line {
		categories.Add(c)
	}
	return categories.AllPairs()
}

func part3(lines []string) int {
	return len(lines)
}

type Categories struct {
	Sword, Archery, Magic []rune
}

func (c *Categories) Add(r rune) {
	switch r {
	case 'A', 'a':
		c.Sword = append(c.Sword, r)
	case 'B', 'b':
		c.Archery = append(c.Archery, r)
	default:
		c.Magic = append(c.Magic, r)
	}
}

func (c *Categories) AllPairs() int {
	return c.SwordPairs() + c.ArcheryPairs() + c.MagicPairs()
}

func (c *Categories) SwordPairs() int {
	return c.GetMentorNovicePairCount(c.Sword, 'A', 'a')
}

func (c *Categories) ArcheryPairs() int {
	return c.GetMentorNovicePairCount(c.Archery, 'B', 'b')
}

func (c *Categories) MagicPairs() int {
	return c.GetMentorNovicePairCount(c.Magic, 'C', 'c')
}

func (c *Categories) GetMentorNovicePairCount(category []rune, mentor, novice rune) int {
	count := 0
	for i := 0; i < len(category); i++ {
		if category[i] == mentor {
			for j := i + 1; j < len(category); j++ {
				if category[j] == novice {
					count++
				}
			}
		}
	}
	return count
}
