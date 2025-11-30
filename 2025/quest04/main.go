package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	input1, _ := os.ReadFile("2025/quest04/input1.txt")
	fmt.Println("2025 Quest 04 Solution")
	fmt.Printf("Part 1: %v\n", part1(input1))
}

func part1(input []byte) int {
	split := strings.Split(strings.TrimSpace(string(input)), "\n")
	first, last := StrToFloat(split[0]), StrToFloat(split[len(split)-1])
	return int((first / last) * 2025)
}

func StrToFloat(s string) float64 {
	num, _ := strconv.ParseFloat(s, 64)
	return num
}
