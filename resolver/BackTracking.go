package resolver

import "fmt"

type BackTrackingAlgorithm string

func (r BackTrackingAlgorithm) Solve(board *[9][9]int) bool {
	fmt.Println("Using BackTracking algorithm")
	return solveIntern(board)
}

func solveIntern(board *[9][9]int) bool {
	for row := 0; row < SIZE; row++ {
		for col := 0; col < SIZE; col++ {
			if board[row][col] == EMPTY {
				for number := 1; number <= SIZE; number++ {
					if isOk(board, row, col, number) {
						board[row][col] = number

						if solveIntern(board) {
							return true
						} else {
							board[row][col] = EMPTY
						}
					}
				}

				return false
			}
		}
	}
	return true
}

func isInRow(board *[9][9]int, row int, number int) bool {
	for i := 0; i < SIZE; i++ {
		if board[row][i] == number {
			return true
		}
	}
	return false
}

func isInCol(board *[9][9]int, col int, number int) bool {
	for i := 0; i < SIZE; i++ {
		if board[i][col] == number {
			return true
		}
	}
	return false
}

func isInBox(board *[9][9]int, row int, col int, number int) bool {
	r := row - row%BOX
	c := col - col%BOX

	for i := r; i < r+BOX; i++ {
		for j := c; j < c+BOX; j++ {
			if board[i][j] == number {
				return true
			}
		}
	}
	return false
}

func isOk(board *[9][9]int, row int, col int, number int) bool {
	return !isInRow(board, row, number) &&
		!isInCol(board, col, number) &&
		!isInBox(board, row, col, number)
}