package main

import (
	"everybody-codes/utils"
	"testing"
)

func TestPart1(t *testing.T) {
	line1, _ := utils.ReadLine("input_test1.txt")
	line2, _ := utils.ReadLine("input1.txt")
	tests := []struct {
		input    string
		expected int
	}{
		{line1, 29},
		{line2, 2699},
	}
	for _, test := range tests {
		result := part1(test.input)
		if result != test.expected {
			t.Errorf("Result %v not equal to expected %v", result, test.expected)
		}
	}
}

func TestPart2(t *testing.T) {
	line1, _ := utils.ReadLine("input_test2.txt")
	line2, _ := utils.ReadLine("input2.txt")
	tests := []struct {
		input    string
		expected int
	}{
		{line1, 781},
		{line2, 213},
	}
	for _, test := range tests {
		result := part2(test.input)
		if result != test.expected {
			t.Errorf("Result %v not equal to expected %v", result, test.expected)
		}
	}
}

func TestPart3(t *testing.T) {
	line1, _ := utils.ReadLine("input_test3.txt")
	line2, _ := utils.ReadLine("input3.txt")
	tests := []struct {
		input    string
		expected int
	}{
		{line1, 3},
		{line2, 4060},
	}
	for _, test := range tests {
		result := part3(test.input)
		if result != test.expected {
			t.Errorf("Result %v not equal to expected %v", result, test.expected)
		}
	}
}

func BenchmarkPart1(b *testing.B) {
	line, _ := utils.ReadLine("input1.txt")
	for n := 0; n < b.N; n++ {
		part1(line)
	}
}

func BenchmarkPart2(b *testing.B) {
	line, _ := utils.ReadLine("input2.txt")
	for n := 0; n < b.N; n++ {
		part2(line)
	}
}

func BenchmarkPart3(b *testing.B) {
	line, _ := utils.ReadLine("input3.txt")
	for n := 0; n < b.N; n++ {
		part3(line)
	}
}
