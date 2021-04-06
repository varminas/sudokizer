package resolver

import (
	"github.com/varminas/sudokizer/model"
)

var SIZE int = 9
var BOX int = 3
var EMPTY int = 0

type MissingValues struct {
	Rows    [][]int
	Cols    [][]int
	Squares [][]int
}

func Resolve(initSudokuValues model.SudokuValues, algorithm model.Algorithm) model.SudokuValues {
	tmpResult := initSudokuValues.Values

	// Apply BackTracking algorithm
	if algorithm == model.BackTracking {
		solveWithBackTracking(&tmpResult)
	}

	if algorithm == model.DancingLinks {
		resolveByDancingLinks(&tmpResult)
	}

	return model.SudokuValues{Values: tmpResult}
}

// Apply Dancing Links method
func resolveByDancingLinks(tmpResult *[9][9]int) {

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


func solveWithBackTracking(board *[9][9]int) bool {
	for row := 0; row < SIZE; row++ {
		for col := 0; col < SIZE; col++ {
			if board[row][col] == EMPTY {
				for number := 1; number <= SIZE; number++ {
					if isOk(board, row, col, number) {
						board[row][col] = number

						if solveWithBackTracking(board) {
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
