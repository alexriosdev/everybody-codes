package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	input1, _ := os.ReadFile("2025/quest05/input1.txt")
	input2, _ := os.ReadFile("2025/quest05/input2.txt")
	fmt.Println("2025 Quest 05 Solution")
	fmt.Printf("Part 1: %v\n", part1(input1))
	fmt.Printf("Part 2: %v\n", part2(input2))
}

func part1(input []byte) int64 {
	split := strings.Split(strings.TrimSpace(string(input)), ":")
	fishbone := Fishbone{}
	for _, s := range strings.Split(split[1], ",") {
		fishbone.Add(StrToInt(s))
	}
	return fishbone.Quality()
}

func part2(input []byte) int64 {
	lines := strings.Split(strings.TrimSpace(string(input)), "\n")
	maxQuality, minQuality := int64(math.MinInt64), int64(math.MaxInt64)
	for _, line := range lines {
		split := strings.Split(line, ":")
		fishbone := Fishbone{}
		for _, s := range strings.Split(split[1], ",") {
			fishbone.Add(StrToInt(s))
		}
		maxQuality = max(maxQuality, fishbone.Quality())
		minQuality = min(minQuality, fishbone.Quality())
	}
	return maxQuality - minQuality
}

type Fishbone struct {
	Segments []Segment
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
	return StrToInt64(quality.String())
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

func StrToInt(s string) int {
	num, _ := strconv.Atoi(s)
	return num
}

func StrToInt64(s string) int64 {
	num, _ := strconv.ParseInt(s, 10, 64)
	return num
}
