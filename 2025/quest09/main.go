package main

import (
	"everybody-codes/utils"
	"fmt"
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
		scales.Add(NewScale(split[0], split[1]))
	}
	child := scales.AssignFamily()
	return child.CalculateDegreeOfSimilarity()
}

func part2(lines []string) int {
	scales := Scales{}
	for _, line := range lines {
		split := strings.Split(line, ":")
		scales.Add(NewScale(split[0], split[1]))
	}
	sum := 0
	for i := 0; i < len(scales); i++ {
		for j := i + 1; j < len(scales); j++ {
			for k := j + 1; k < len(scales); k++ {
				newScales := Scales{scales[i], scales[j], scales[k]}
				if child := newScales.AssignFamily(); child != nil {
					sum += child.CalculateDegreeOfSimilarity()
				}
			}
		}
	}
	return sum
}

func part3(lines []string) int {
	scales := Scales{}
	for _, line := range lines {
		split := strings.Split(line, ":")
		scales.Add(NewScale(split[0], split[1]))
	}
	root := scales[0]
	maxDepth := 1
	for i := 0; i < len(scales); i++ {
		for j := i + 1; j < len(scales); j++ {
			for k := j + 1; k < len(scales); k++ {
				newScales := Scales{scales[i], scales[j], scales[k]}
				if child := newScales.AssignFamily(); child != nil && maxDepth < child.Depth {
					maxDepth = child.Depth
					root = child
				}
			}
		}
	}
	sum := 0
	visited := map[int]bool{}
	queue := []*Scale{root}
	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]
		if visited[curr.Number] {
			continue
		}
		visited[curr.Number] = true
		sum += curr.Number
		for child := range curr.Children {
			queue = append(queue, child)
		}
		if a := curr.ParentA; a != nil {
			queue = append(queue, a)
		}
		if b := curr.ParentB; b != nil {
			queue = append(queue, b)
		}
	}
	return sum
}

type Scale struct {
	Number, Depth    int
	DNA              []rune
	ParentA, ParentB *Scale
	Children         map[*Scale]bool
}

func NewScale(s1, s2 string) *Scale {
	return &Scale{
		Number:   utils.StrToInt(s1),
		Depth:    1,
		DNA:      []rune(s2),
		Children: map[*Scale]bool{},
	}
}

func (s *Scale) AssignParents(a, b *Scale) {
	s.ParentA = a
	s.ParentB = b
	s.ParentA.Children[s] = true
	s.ParentB.Children[s] = true
	s.Depth = s.ParentA.Depth + s.ParentB.Depth
}

func (s *Scale) CalculateDegreeOfSimilarity() int {
	degreeA, degreeB := 0, 0
	for i := range s.DNA {
		if s.DNA[i] == s.ParentA.DNA[i] {
			degreeA++
		}
		if s.DNA[i] == s.ParentB.DNA[i] {
			degreeB++
		}
	}
	return degreeA * degreeB
}

type Scales []*Scale

func (s *Scales) Add(scale *Scale) {
	*s = append(*s, scale)
}

func (s *Scales) AssignFamily() *Scale {
	a, b, c := (*s)[0], (*s)[1], (*s)[2]
	aCount, bCount, cCount := 0, 0, 0
	for i := range a.DNA {
		switch {
		case (a.DNA[i] == b.DNA[i]) && (b.DNA[i] == c.DNA[i]) && (a.DNA[i] == c.DNA[i]):
			aCount++
			bCount++
			cCount++
		case (a.DNA[i] == b.DNA[i]) && (b.DNA[i] != c.DNA[i]) && (a.DNA[i] != c.DNA[i]):
			aCount++
			bCount++
		case (a.DNA[i] != b.DNA[i]) && (b.DNA[i] != c.DNA[i]) && (a.DNA[i] == c.DNA[i]):
			aCount++
			cCount++
		case (a.DNA[i] != b.DNA[i]) && (b.DNA[i] == c.DNA[i]) && (a.DNA[i] != c.DNA[i]):
			bCount++
			cCount++
		}
	}
	switch {
	case aCount == len(a.DNA):
		a.AssignParents(b, c)
		return a
	case bCount == len(a.DNA):
		b.AssignParents(a, c)
		return b
	case cCount == len(a.DNA):
		c.AssignParents(a, b)
		return c
	}
	return nil
}
