package main

import (
	"everybody-codes/utils"
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"
)

func main() {
	lines1, _ := utils.ReadLines("2025/quest05/input1.txt")
	lines2, _ := utils.ReadLines("2025/quest05/input2.txt")
	lines3, _ := utils.ReadLines("2025/quest05/input3.txt")
	fmt.Println("2025 Quest 05 Solution")
	fmt.Printf("Part 1: %v\n", part1(lines1))
	fmt.Printf("Part 2: %v\n", part2(lines2))
	fmt.Printf("Part 2: %v\n", part3(lines3))
}

func part1(lines []string) int64 {
	fishbone := Fishbone{}
	for _, line := range lines {
		split := strings.Split(line, ":")
		for _, s := range strings.Split(split[1], ",") {
			fishbone.Add(utils.StrToInt(s))
		}
	}
	return fishbone.Quality()
}

func part2(lines []string) int64 {
	maxQuality, minQuality := int64(math.MinInt64), int64(math.MaxInt64)
	for _, line := range lines {
		split := strings.Split(line, ":")
		fishbone := Fishbone{}
		for _, s := range strings.Split(split[1], ",") {
			fishbone.Add(utils.StrToInt(s))
		}
		maxQuality = max(maxQuality, fishbone.Quality())
		minQuality = min(minQuality, fishbone.Quality())
	}
	return maxQuality - minQuality
}

func part3(lines []string) int64 {
	swords := Swords{}
	for _, line := range lines {
		split := strings.Split(line, ":")
		fishbone := Fishbone{ID: utils.StrToInt(split[0])}
		for _, s := range strings.Split(split[1], ",") {
			fishbone.Add(utils.StrToInt(s))
		}
		fishbone.GenerateLevels()
		swords.Add(fishbone)
	}
	swords.Sort()
	return swords.Checksum()
}

type Swords struct {
	Fishbones []Fishbone
}

func (s *Swords) Add(fishbone Fishbone) {
	s.Fishbones = append(s.Fishbones, fishbone)
}

func (s *Swords) Sort() {
	fishbones := s.Fishbones
	sort.Slice(fishbones, func(i, j int) bool {
		if fishbones[i].Quality() != fishbones[j].Quality() {
			return fishbones[i].Quality() > fishbones[j].Quality()
		} else if fishbones[i].Quality() == fishbones[j].Quality() {
			a, b := fishbones[i].Levels, fishbones[j].Levels
			for k, l := 0, 0; k < len(a) && l < len(b); k, l = k+1, k+1 {
				if a[k] != b[l] {
					return a[k] > b[l]
				}
			}
		}
		return fishbones[i].ID > fishbones[j].ID
	})
}

func (s *Swords) Checksum() int64 {
	sum := int64(0)
	for i, fishbone := range s.Fishbones {
		sum += int64((i + 1) * fishbone.ID)
	}
	return sum
}

type Fishbone struct {
	ID       int
	Segments []Segment
	Levels   []int
}

func (f *Fishbone) Add(num int) {
	for i := range f.Segments {
		if f.Segments[i].Add(num) {
			return
		}
	}
	f.Segments = append(f.Segments, Segment{Mid: num})
}

func (f *Fishbone) Quality() int64 {
	quality := strings.Builder{}
	for _, segment := range f.Segments {
		quality.WriteString(strconv.Itoa(segment.Mid))
	}
	return utils.StrToInt64(quality.String())
}

func (f *Fishbone) GenerateLevels() {
	sb := strings.Builder{}
	for _, segment := range f.Segments {
		if segment.Left > 0 {
			sb.WriteString(strconv.Itoa(segment.Left))
		}
		if segment.Mid > 0 {
			sb.WriteString(strconv.Itoa(segment.Mid))
		}
		if segment.Right > 0 {
			sb.WriteString(strconv.Itoa(segment.Right))
		}
		f.Levels = append(f.Levels, utils.StrToInt(sb.String()))
		sb.Reset()
	}
}

func (f *Fishbone) Print() {
	for _, segment := range f.Segments {
		fmt.Println(segment.Left, segment.Mid, segment.Right)
	}
}

type Segment struct {
	Mid, Left, Right int
}

func (s *Segment) Add(num int) bool {
	switch {
	case s.Mid == 0:
		s.Mid = num
		return true
	case num < s.Mid && s.Left == 0:
		s.Left = num
		return true
	case num > s.Mid && s.Right == 0:
		s.Right = num
		return true
	}
	return false
}
