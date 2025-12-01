package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	input1, _ := os.ReadFile("2025/quest05/input1.txt")
	fmt.Println("2025 Quest 05 Solution")
	fmt.Printf("Part 1: %v\n", part1(input1))
}

func part1(input []byte) string {
	split := strings.Split(strings.TrimSpace(string(input)), ":")
	fishbone := Fishbone{}
	for _, s := range strings.Split(split[1], ",") {
		fishbone.Add(StrToInt(s))
	}
	return fishbone.Quality()
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

func (f *Fishbone) Quality() string {
	quality := strings.Builder{}
	for _, segment := range f.Segments {
		quality.WriteString(strconv.Itoa(segment.Mid))
	}
	return quality.String()
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
