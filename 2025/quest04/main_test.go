package main

import (
	"everybody-codes/utils"
	"testing"
)

func TestPart1(t *testing.T) {
	lines1, _ := utils.ReadLines("input_test1.txt")
	lines2, _ := utils.ReadLines("input_test2.txt")
	lines3, _ := utils.ReadLines("input1.txt")
	tests := []struct {
		input    []string
		expected int
	}{
		{lines1, 32400},
		{lines2, 15888},
		{lines3, 10175},
	}
	for _, test := range tests {
		result := part1(test.input)
		if result != test.expected {
			t.Errorf("Result %v not equal to expected %v", result, test.expected)
		}
	}
}

func TestPart2(t *testing.T) {
	lines1, _ := utils.ReadLines("input_test3.txt")
	lines2, _ := utils.ReadLines("input_test4.txt")
	lines3, _ := utils.ReadLines("input2.txt")
	tests := []struct {
		input    []string
		expected int
	}{
		{lines1, 625000000000},
		{lines2, 1274509803922},
		{lines3, 1523178807948},
	}
	for _, test := range tests {
		result := part2(test.input)
		if result != test.expected {
			t.Errorf("Result %v not equal to expected %v", result, test.expected)
		}
	}
}

func TestPart3(t *testing.T) {
	lines1, _ := utils.ReadLines("input_test5.txt")
	lines2, _ := utils.ReadLines("input_test6.txt")
	lines3, _ := utils.ReadLines("input3.txt")
	tests := []struct {
		input    []string
		expected int
	}{
		{lines1, 400},
		{lines2, 6818},
		{lines3, 422511306938},
	}
	for _, test := range tests {
		result := part3(test.input)
		if result != test.expected {
			t.Errorf("Result %v not equal to expected %v", result, test.expected)
		}
	}
}

func BenchmarkPart1(b *testing.B) {
	lines, _ := utils.ReadLines("input1.txt")
	for n := 0; n < b.N; n++ {
		part1(lines)
	}
}

func BenchmarkPart2(b *testing.B) {
	lines, _ := utils.ReadLines("input2.txt")
	for n := 0; n < b.N; n++ {
		part2(lines)
	}
}

func BenchmarkPart3(b *testing.B) {
	lines, _ := utils.ReadLines("input3.txt")
	for n := 0; n < b.N; n++ {
		part3(lines)
	}
}
