package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	input1, _ := os.ReadFile("2025/quest03/input1.txt")
	input2, _ := os.ReadFile("2025/quest03/input2.txt")
	input3, _ := os.ReadFile("2025/quest03/input3.txt")
	fmt.Println("2025 Quest 03 Solution")
	fmt.Printf("Part 1: %v\n", part1(input1))
	fmt.Printf("Part 2: %v\n", part2(input2))
	fmt.Printf("Part 3: %v\n", part3(input3))
}

func part1(input []byte) int {
	split := strings.Split(strings.TrimSpace(string(input)), ",")
	set := map[int]bool{}
	for _, s := range split {
		set[StrToInt(s)] = true
	}
	sum := 0
	for num := range set {
		sum += num
	}
	return sum
}

func part2(input []byte) int {
	split := strings.Split(strings.TrimSpace(string(input)), ",")
	set := map[int]bool{}
	for _, s := range split {
		set[StrToInt(s)] = true
	}
	uniqueNums := []int{}
	for num := range set {
		uniqueNums = append(uniqueNums, num)
	}
	sort.Ints(uniqueNums)
	sum := 0
	for i := 0; i < 20; i++ {
		sum += uniqueNums[i]
	}
	return sum
}

func part3(input []byte) int {
	split := strings.Split(strings.TrimSpace(string(input)), ",")
	freq := map[int]int{}
	for _, s := range split {
		freq[StrToInt(s)]++
	}
	count := 0
	for _, v := range freq {
		count = max(count, v)
	}
	return count
}

func StrToInt(s string) int {
	num, _ := strconv.Atoi(s)
	return num
}
