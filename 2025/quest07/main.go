package main

import (
	"fmt"
	"os"
	"slices"
	"strings"
)

func main() {
	input1, _ := os.ReadFile("2025/quest07/input1.txt")
	input2, _ := os.ReadFile("2025/quest07/input2.txt")
	input3, _ := os.ReadFile("2025/quest07/input3.txt")
	fmt.Println("2025 Quest 07 Solution")
	fmt.Printf("Part 1: %v\n", part1(input1))
	fmt.Printf("Part 2: %v\n", part2(input2))
	fmt.Printf("Part 3: %v\n", part3(input3))
}

func part1(input []byte) string {
	sections := strings.Split(strings.TrimSpace(string(input)), "\n\n")
	replacer := strings.NewReplacer(">", "", ",", "", " ", "")
	rules := map[rune][]rune{}
	for _, s := range strings.Split(sections[1], "\n") {
		runes := []rune(replacer.Replace(s))
		rules[runes[0]] = append(rules[runes[0]], runes[1:]...)
	}
	for _, name := range strings.Split(sections[0], ",") {
		if isValidName(name, rules) {
			return name
		}
	}
	return ""
}

func part2(input []byte) string {
	return ""
}

func part3(input []byte) string {
	return ""
}

func isValidName(name string, rules map[rune][]rune) bool {
	for i := 0; i < len(name)-1; i++ {
		a, b := rune(name[i]), rune(name[i+1])
		if !slices.Contains(rules[a], b) {
			return false
		}
	}
	return true
}
