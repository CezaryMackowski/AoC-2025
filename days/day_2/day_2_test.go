package day_2

import (
	"bufio"
	"os"
	"strconv"
	"strings"
	"testing"
)

func getInput(filename string) []identifierRange {
	f, err := os.Open(filename)
	if err != nil {
		return nil
	}
	defer f.Close()

	var idRanges []identifierRange

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()

		for _, ir := range strings.Split(line, ",") {
			singleRange := strings.Split(ir, "-")
			start, _ := strconv.Atoi(singleRange[0])
			end, _ := strconv.Atoi(singleRange[1])
			idRanges = append(idRanges, identifierRange{start: start, end: end})
		}
	}

	return idRanges
}

func TestDay2(t *testing.T) {
	inputRotations := getInput("../inputs/day_2.txt")

	var tests = []struct {
		idRanges       []identifierRange
		expectedResult int
	}{
		{[]identifierRange{{start: 11, end: 22}, {start: 95, end: 115}, {start: 998, end: 1012}, {start: 1188511880, end: 1188511890}, {start: 222220, end: 222224}, {start: 1698522, end: 1698528}, {start: 446443, end: 446449}, {start: 38593856, end: 38593862}, {start: 565653, end: 565659}, {start: 824824821, end: 824824827}, {start: 2121212118, end: 2121212124}}, 1227775554},
		{inputRotations, 18952700150},
	}

	for _, tt := range tests {
		t.Run("Day 2", func(t *testing.T) {
			result := day2(tt.idRanges)

			if result != tt.expectedResult {
				t.Errorf("Wrong - got %d, expected %d", result, tt.expectedResult)
			}
		})
	}
}

func TestDay2Extra(t *testing.T) {
	inputRotations := getInput("../inputs/day_2.txt")

	var tests = []struct {
		idRanges       []identifierRange
		expectedResult int
	}{
		{[]identifierRange{{start: 11, end: 22}, {start: 95, end: 115}, {start: 998, end: 1012}, {start: 1188511880, end: 1188511890}, {start: 222220, end: 222224}, {start: 1698522, end: 1698528}, {start: 446443, end: 446449}, {start: 38593856, end: 38593862}, {start: 565653, end: 565659}, {start: 824824821, end: 824824827}, {start: 2121212118, end: 2121212124}}, 4174379265},
		{inputRotations, 28858486244},
	}

	for _, tt := range tests {
		t.Run("Day 2 (Extra)", func(t *testing.T) {
			result := day2Extra(tt.idRanges)

			if result != tt.expectedResult {
				t.Errorf("Wrong - got %d, expected %d", result, tt.expectedResult)
			}
		})
	}
}
