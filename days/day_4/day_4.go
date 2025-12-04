package day_4

var kernel = [8][2]int{
	{-1, -1}, {-1, 0}, {-1, 1},
	{0, -1} /*self*/, {0, 1},
	{1, -1}, {1, 0}, {1, 1},
}

type position struct {
	row int
	col int
}

func countNeighbors(rows []string, row, col int) int {
	neighbors := 0

	for _, neighbor := range kernel {
		neighborRow := row + neighbor[0]
		neighborCol := col + neighbor[1]

		if neighborRow < 0 || neighborRow >= len(rows) {
			continue
		}

		if neighborCol < 0 || neighborCol >= len(rows[neighborRow]) {
			continue
		}

		if rows[neighborRow][neighborCol] == '@' {
			neighbors++
		}
	}

	return neighbors
}

func countAvailableRolls(rows []string) int {
	countOfRolls := 0

	for i := 0; i < len(rows); i++ {
		for j := 0; j < len(rows[i]); j++ {
			if rows[i][j] != '@' {
				continue
			}

			neighborCount := countNeighbors(rows, i, j)
			if neighborCount < 4 {
				countOfRolls++
			}
		}
	}

	return countOfRolls
}

func replaceChar(row string, col int, ch byte) string {
	b := []byte(row)
	b[col] = ch
	return string(b)
}

func day4(rows []string) int {
	return countAvailableRolls(rows)
}

func day4Extra(rows []string) int {
	grid := make([]string, len(rows))
	copy(grid, rows)
	totalRemoved := 0

	for {
		var toRemove []position

		for i := 0; i < len(grid); i++ {
			for j := 0; j < len(grid[i]); j++ {
				if grid[i][j] != '@' {
					continue
				}

				if countNeighbors(grid, i, j) < 4 {
					toRemove = append(toRemove, position{row: i, col: j})
				}
			}
		}

		if len(toRemove) == 0 {
			break
		}

		for _, pos := range toRemove {
			if grid[pos.row][pos.col] == '@' {
				grid[pos.row] = replaceChar(grid[pos.row], pos.col, '.')
				totalRemoved++
			}
		}
	}

	return totalRemoved
}
