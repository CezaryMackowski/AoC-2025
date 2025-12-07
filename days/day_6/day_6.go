package day_6

type operator string

const (
	multiplication operator = "*"
	addition       operator = "+"
)

type operation struct {
	numbers   []int
	operators []operator
}

func columnBlank(lines []string, col int) bool {
	for _, line := range lines {
		if col < len(line) && line[col] != ' ' {
			return false
		}
	}

	return true
}

func applyOperation(numbers []int, numbersOperator operator) int {
	switch numbersOperator {
	case multiplication:
		tmp := 1
		for _, number := range numbers {
			tmp *= number
		}
		return tmp
	case addition:
		tmp := 0
		for _, number := range numbers {
			tmp += number
		}
		return tmp
	default:
		return 0
	}
}

func day6(operations []operation) int {
	grandTotal := 0

	for _, o := range operations {
		grandTotal += applyOperation(o.numbers, o.operators[0])
	}

	return grandTotal
}

func day6Extra(operations []operation) int {
	// TODO
	return 0
}
