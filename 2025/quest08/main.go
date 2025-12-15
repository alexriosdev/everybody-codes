package main

import (
	"everybody-codes/utils"
	"fmt"
	"strings"
)

func main() {
	line1, _ := utils.ReadLine("2025/quest08/input1.txt")
	line2, _ := utils.ReadLine("2025/quest08/input2.txt")
	line3, _ := utils.ReadLine("2025/quest08/input3.txt")
	fmt.Println("2025 Quest 08 Solution")
	fmt.Printf("Part 1: %v\n", part1(line1, 32))
	fmt.Printf("Part 2: %v\n", part2(line2, 256))
	fmt.Printf("Part 3: %v\n", part3(line3, 256))
}

func part1(line string, nails int) int {
	delta := nails / 2
	count := 0
	split := strings.Split(line, ",")
	for i := 0; i < len(split)-1; i++ {
		start, end := utils.StrToInt(split[i]), utils.StrToInt(split[i+1])
		if start > end {
			start, end = end, start
		}
		if end-start == delta {
			count++
		}
	}
	return count
}

func part2(line string, nails int) int {
	split := strings.Split(line, ",")
	threads := make([][]int, nails+1)
	for i := 0; i < len(split)-1; i++ {
		start, end := utils.StrToInt(split[i]), utils.StrToInt(split[i+1])
		if start > end {
			start, end = end, start
		}
		threads[start] = append(threads[start], end)
	}
	freq := make([]int, nails+1)
	sum := 0
	for start, ends := range threads {
		for _, end := range ends {
			for i := start + 1; i < end; i++ {
				sum += freq[i]
			}
		}
		for _, end := range ends {
			freq[end]++
		}
	}
	return sum
}

func part3(line string, nails int) int {
	split := strings.Split(line, ",")
	maxCuts := 0
	for i := 1; i < nails; i++ {
		for j := i + 1; j < nails; j++ {
			cuts := 0
			for k := 0; k < len(split)-1; k++ {
				start, end := utils.StrToInt(split[k]), utils.StrToInt(split[k+1])
				if start > end {
					start, end = end, start
				}
				if (start < i && i < end && end < j) ||
					(i < start && start < j && j < end) ||
					(i == start && j == end) {
					cuts++
				}
			}
			maxCuts = max(maxCuts, cuts)
		}
	}
	return maxCuts
}
