package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	input1, _ := os.ReadFile("2025/quest03/input1.txt")
	fmt.Println("2025 Quest 03 Solution")
	fmt.Printf("Part 1: %v\n", part1(input1))
}

func part1(input []byte) int {
	split := strings.Split(string(input), ",")
	uniqueNums := map[int]bool{}
	for _, s := range split {
		uniqueNums[StrToInt(s)] = true
	}
	sum := 0
	for num := range uniqueNums {
		sum += num
	}
	return sum
}

func StrToInt(s string) int {
	num, _ := strconv.Atoi(s)
	return num
}
