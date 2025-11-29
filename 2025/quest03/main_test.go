package main

import (
	"os"
	"testing"
)

func TestPart1(t *testing.T) {
	input1, _ := os.ReadFile("input_test1.txt")
	input2, _ := os.ReadFile("input1.txt")
	tests := []struct {
		input    []byte
		expected int
	}{
		{input1, 29},
		{input2, 2699},
	}
	for _, test := range tests {
		result := part1(test.input)
		if result != test.expected {
			t.Errorf("Result %v not equal to expected %v", result, test.expected)
		}
	}
}

func TestPart2(t *testing.T) {
	input1, _ := os.ReadFile("input_test2.txt")
	input2, _ := os.ReadFile("input2.txt")
	tests := []struct {
		input    []byte
		expected int
	}{
		{input1, 781},
		{input2, 213},
	}
	for _, test := range tests {
		result := part2(test.input)
		if result != test.expected {
			t.Errorf("Result %v not equal to expected %v", result, test.expected)
		}
	}
}

func TestPart3(t *testing.T) {
	input1, _ := os.ReadFile("input_test3.txt")
	input2, _ := os.ReadFile("input3.txt")
	tests := []struct {
		input    []byte
		expected int
	}{
		{input1, 3},
		{input2, 4060},
	}
	for _, test := range tests {
		result := part3(test.input)
		if result != test.expected {
			t.Errorf("Result %v not equal to expected %v", result, test.expected)
		}
	}
}

func BenchmarkPart1(b *testing.B) {
	input, _ := os.ReadFile("input1.txt")
	for n := 0; n < b.N; n++ {
		part1(input)
	}
}

func BenchmarkPart2(b *testing.B) {
	input, _ := os.ReadFile("input2.txt")
	for n := 0; n < b.N; n++ {
		part2(input)
	}
}

func BenchmarkPart3(b *testing.B) {
	input, _ := os.ReadFile("input3.txt")
	for n := 0; n < b.N; n++ {
		part3(input)
	}
}
