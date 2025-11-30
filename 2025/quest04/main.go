package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	input1, _ := os.ReadFile("2025/quest04/input1.txt")
	input2, _ := os.ReadFile("2025/quest04/input2.txt")
	input3, _ := os.ReadFile("2025/quest04/input3.txt")
	fmt.Println("2025 Quest 04 Solution")
	fmt.Printf("Part 1: %v\n", part1(input1))
	fmt.Printf("Part 2: %v\n", part2(input2))
	fmt.Printf("Part 3: %v\n", part3(input3))
}

func part1(input []byte) int {
	split := strings.Split(strings.TrimSpace(string(input)), "\n")
	first, last := StrToFloat(split[0]), StrToFloat(split[len(split)-1])
	return int((first / last) * 2025)
}

func part2(input []byte) int {
	split := strings.Split(strings.TrimSpace(string(input)), "\n")
	first, last := StrToFloat(split[0]), StrToFloat(split[len(split)-1])
	return int(math.Ceil((last * 10_000_000_000_000) / first))
}

func part3(input []byte) int {
	split := strings.Split(strings.TrimSpace(string(input)), "\n")
	first, last := StrToFloat(split[0]), StrToFloat(split[len(split)-1])
	for _, s := range split[1 : len(split)-1] {
		innerSplit := strings.Split(s, "|")
		left, right := StrToFloat(innerSplit[0]), StrToFloat(innerSplit[1])
		first *= right / left
	}
	return int((first * 100) / last)
}

func StrToFloat(s string) float64 {
	num, _ := strconv.ParseFloat(s, 64)
	return num
}
