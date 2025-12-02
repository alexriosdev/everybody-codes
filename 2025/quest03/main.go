package main

import (
	"everybody-codes/utils"
	"fmt"
	"sort"
	"strings"
)

func main() {
	line1, _ := utils.ReadLine("2025/quest03/input1.txt")
	line2, _ := utils.ReadLine("2025/quest03/input2.txt")
	line3, _ := utils.ReadLine("2025/quest03/input3.txt")
	fmt.Println("2025 Quest 03 Solution")
	fmt.Printf("Part 1: %v\n", part1(line1))
	fmt.Printf("Part 2: %v\n", part2(line2))
	fmt.Printf("Part 3: %v\n", part3(line3))
}

func part1(line string) int {
	split := strings.Split(line, ",")
	set := map[int]bool{}
	for _, s := range split {
		set[utils.StrToInt(s)] = true
	}
	sum := 0
	for num := range set {
		sum += num
	}
	return sum
}

func part2(line string) int {
	split := strings.Split(line, ",")
	set := map[int]bool{}
	for _, s := range split {
		set[utils.StrToInt(s)] = true
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

func part3(line string) int {
	split := strings.Split(line, ",")
	freq := map[int]int{}
	for _, s := range split {
		freq[utils.StrToInt(s)]++
	}
	count := 0
	for _, v := range freq {
		count = max(count, v)
	}
	return count
}
