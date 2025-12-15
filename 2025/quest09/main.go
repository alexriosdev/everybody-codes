package main

import (
	"everybody-codes/utils"
	"fmt"
	"sort"
	"strings"
)

func main() {
	lines1, _ := utils.ReadLines("2025/quest09/input1.txt")
	lines2, _ := utils.ReadLines("2025/quest09/input2.txt")
	lines3, _ := utils.ReadLines("2025/quest09/input3.txt")
	fmt.Println("2025 Quest 09 Solution")
	fmt.Printf("Part 1: %v\n", part1(lines1))
	fmt.Printf("Part 2: %v\n", part2(lines2))
	fmt.Printf("Part 3: %v\n", part3(lines3))
}

func part1(lines []string) int {
	scales := Scales{}
	for _, line := range lines {
		split := strings.Split(line, ":")
		scales.Add(NewScale(split[1]))
	}
	scales.SortByChild()
	return scales.CalculateDegreesOfSimilarity()
}

func part2(lines []string) int {
	return len(lines)
}

func part3(lines []string) int {
	return len(lines)
}

type Scale struct {
	DNA     []rune
	IsChild int
}

func NewScale(s string) *Scale {
	return &Scale{DNA: []rune(s)}
}

type Scales []*Scale

func (s *Scales) Add(scale *Scale) {
	*s = append(*s, scale)
}

func (s *Scales) SortByChild() {
	a, b, c := (*s)[0], (*s)[1], (*s)[2]
	for i := range a.DNA {
		switch {
		case (a.DNA[i] == b.DNA[i]) && (b.DNA[i] == c.DNA[i]) && (a.DNA[i] == c.DNA[i]):
			a.IsChild++
			b.IsChild++
			c.IsChild++
		case (a.DNA[i] == b.DNA[i]) && (b.DNA[i] != c.DNA[i]) && (a.DNA[i] != c.DNA[i]):
			a.IsChild++
			b.IsChild++
		case (a.DNA[i] != b.DNA[i]) && (b.DNA[i] != c.DNA[i]) && (a.DNA[i] == c.DNA[i]):
			a.IsChild++
			c.IsChild++
		case (a.DNA[i] != b.DNA[i]) && (b.DNA[i] == c.DNA[i]) && (a.DNA[i] != c.DNA[i]):
			b.IsChild++
			c.IsChild++
		}
	}
	sort.Slice(*s, func(i, j int) bool {
		return (*s)[i].IsChild > (*s)[j].IsChild
	})
}

func (s *Scales) CalculateDegreesOfSimilarity() int {
	a, b, c := (*s)[0], (*s)[1], (*s)[2]
	degreeA, degreeB := 0, 0
	for i := range a.DNA {
		if a.DNA[i] == b.DNA[i] {
			degreeA++
		}
		if a.DNA[i] == c.DNA[i] {
			degreeB++
		}
	}
	return degreeA * degreeB
}
