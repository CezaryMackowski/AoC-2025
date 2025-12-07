package day_6

import (
	"bufio"
	"os"
	"strconv"
	"strings"
	"testing"
)

func getInput(filename string) []operation {
	f, err := os.Open(filename)
	if err != nil {
		return nil
	}
	defer f.Close()

	var lines []string
	scanner := bufio.NewScanner(f)
	maxWidth := 0
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) > maxWidth {
			maxWidth = len(line)
		}
		lines = append(lines, line)
	}

	if len(lines) == 0 {
		return nil
	}

	// pad lines to same width to simplify column slicing
	for i := range lines {
		if len(lines[i]) < maxWidth {
			lines[i] += strings.Repeat(" ", maxWidth-len(lines[i]))
		}
	}

	lastRow := len(lines) - 1
	var operations []operation

	col := 0
	for col < maxWidth {
		if columnBlank(lines, col) {
			col++
			continue
		}

		start := col
		for col < maxWidth && !columnBlank(lines, col) {
			col++
		}
		end := col

		var numbers []int
		for row := 0; row < lastRow; row++ {
			value := strings.TrimSpace(lines[row][start:end])
			if value == "" {
				continue
			}

			num, _ := strconv.Atoi(value)

			numbers = append(numbers, num)
		}

		if len(numbers) == 0 {
			continue
		}

		opVal := strings.TrimSpace(lines[lastRow][start:end])
		if opVal == "" {
			continue
		}

		op := operator(opVal)
		var ops []operator
		if len(numbers) > 1 {
			ops = make([]operator, len(numbers)-1)
			for i := range ops {
				ops[i] = op
			}
		}

		operations = append(operations, operation{
			numbers:   numbers,
			operators: ops,
		})
	}

	return operations
}

func TestDay6(t *testing.T) {
	operations := getInput("../inputs/day_6.txt")

	var tests = []struct {
		operations     []operation
		expectedResult int
	}{
		{[]operation{{[]int{123, 45, 6}, []operator{"*"}}, {[]int{328, 64, 98}, []operator{"+"}}, {[]int{51, 387, 215}, []operator{"*"}}, {[]int{64, 23, 314}, []operator{"+"}}}, 4277556},
		{operations, 6171290547579},
	}

	for _, tt := range tests {
		t.Run("Day 6", func(t *testing.T) {
			result := day6(tt.operations)

			if result != tt.expectedResult {
				t.Errorf("Wrong - got %d, expected %d", result, tt.expectedResult)
			}
		})
	}
}

func TestDay6Extra(t *testing.T) {
	//operations := getInput("../inputs/day_6.txt")

	var tests = []struct {
		operations     []operation
		expectedResult int
	}{
		{[]operation{{[]int{123, 45, 6}, []operator{"*"}}, {[]int{328, 64, 98}, []operator{"+"}}, {[]int{51, 387, 215}, []operator{"*"}}, {[]int{64, 23, 314}, []operator{"+"}}}, 3263827},
		//{operations, 0},
	}

	for _, tt := range tests {
		t.Run("Day 6 (Extra)", func(t *testing.T) {
			result := day6Extra(tt.operations)

			if result != tt.expectedResult {
				t.Errorf("Wrong - got %d, expected %d", result, tt.expectedResult)
			}
		})
	}
}
