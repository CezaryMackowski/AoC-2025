package day_3

func selectMaxDigits(line string, digits int) string {
	if digits <= 0 || digits > len(line) {
		return ""
	}

	result := make([]byte, 0, digits)
	start := 0
	remaining := digits

	for remaining > 0 {
		lastStart := len(line) - remaining
		maxIdx := start

		for i := start; i <= lastStart; i++ {
			if line[i] > line[maxIdx] {
				maxIdx = i
			}
		}

		result = append(result, line[maxIdx])
		start = maxIdx + 1
		remaining--
	}

	return string(result)
}

func buildMaxJoltageString(line string, digits int) int {
	selected := selectMaxDigits(line, digits)
	if selected == "" {
		return -1
	}

	value := 0
	for i := 0; i < len(selected); i++ {
		digit := selected[i]
		value = value*10 + int(digit-'0')
	}

	return value
}

func day3(lines []string, digits int) int {
	sumOfJoltage := 0

	for _, line := range lines {
		joltage := buildMaxJoltageString(line, digits)

		if joltage == -1 {
			// invalid row
			continue
		}

		sumOfJoltage += joltage
	}

	return sumOfJoltage
}

func day3Extra(lines []string, digits int) int {
	sumOfJoltage := 0

	for _, line := range lines {
		joltage := buildMaxJoltageString(line, digits)

		if joltage == -1 {
			// invalid row
			continue
		}

		sumOfJoltage += joltage
	}

	return sumOfJoltage
}
