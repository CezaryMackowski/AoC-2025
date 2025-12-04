package day_4

import (
	"bufio"
	"os"
	"testing"
)

func getInput(filename string) []string {
	f, err := os.Open(filename)
	if err != nil {
		return nil
	}
	defer f.Close()

	var lines []string

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}

	return lines
}

func TestDay4(t *testing.T) {
	inputRotations := getInput("../inputs/day_4.txt")

	var tests = []struct {
		lines          []string
		expectedResult int
	}{
		{[]string{"..@@.@@@@.", "@@@.@.@.@@", "@@@@@.@.@@", "@.@@@@..@.", "@@.@@@@.@@", ".@@@@@@@.@", ".@.@.@.@@@", "@.@@@.@@@@", ".@@@@@@@@.", "@.@.@@@.@."}, 13},
		{inputRotations, 1441},
	}

	for _, tt := range tests {
		t.Run("Day 4", func(t *testing.T) {
			result := day4(tt.lines)

			if result != tt.expectedResult {
				t.Errorf("Wrong - got %d, expected %d", result, tt.expectedResult)
			}
		})
	}
}

func TestDay4Extra(t *testing.T) {
	inputRotations := getInput("../inputs/day_4.txt")

	var tests = []struct {
		lines          []string
		expectedResult int
	}{
		{[]string{"..@@.@@@@.", "@@@.@.@.@@", "@@@@@.@.@@", "@.@@@@..@.", "@@.@@@@.@@", ".@@@@@@@.@", ".@.@.@.@@@", "@.@@@.@@@@", ".@@@@@@@@.", "@.@.@@@.@."}, 43},
		{inputRotations, 9050},
	}

	for _, tt := range tests {
		t.Run("Day 4 (Extra)", func(t *testing.T) {
			result := day4Extra(tt.lines)

			if result != tt.expectedResult {
				t.Errorf("Wrong - got %d, expected %d", result, tt.expectedResult)
			}
		})
	}
}
