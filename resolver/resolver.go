package resolver

import (
	"fmt"
	"math/rand"

	// "sort"
	"sudokizer/model"
)

var SIZE int = 9
var BOX int = 3
var EMPTY int = 0

type MissingValues struct {
	Rows    [][]int
	Cols    [][]int
	Squares [][]int
}

func Resolve(initSudokuValues model.SudokuValues) model.SudokuValues {
	tmpResult := initSudokuValues.Values

	// missingValues := resolveSingleMissingValues(&tmpResult)
	// fmt.Printf("missRows %#v \n\n", missingValues.Rows)
	// fmt.Printf("missColu %#v \n\n", missingValues.Cols)
	// fmt.Printf("missSqua %#v \n\n", missingValues.Squares)

	// Apply BackTracking algorithm
	solveWithBackTracking(&tmpResult)

	return model.SudokuValues{Values: tmpResult}
}

// Apply Dancing Links method

func resolveByDancingLinks(tmpResult *[9][9]int) {

}

func resolveSingleMissingValues(tmpResult *[9][9]int) MissingValues {
	missingValuesInRows := make([][]int, SIZE)
	missingValuesInColumns := make([][]int, SIZE)
	missingValuesInSquares := make([][]int, SIZE)

	foundSingleMissingValue := true
	for foundSingleMissingValue == true {
		for i, row := range tmpResult {
			foundSingleMissingValue = true
			// Missing row values
			missingValuesInRow, missingIndexesInRow := findMissingValuesAndIndexes(row)
			// Means found single missing value
			if len(missingValuesInRow) == 1 {
				fmt.Println("===Found value in row===")
				fmt.Printf("[%v][%v]=%v\n", i, missingIndexesInRow, missingValuesInRow)
				tmpResult[i][missingIndexesInRow[0]] = missingValuesInRow[0]
				fmt.Printf("\n\n")
				break
			}
			missingValuesInRows[i] = missingValuesInRow

			// Missing column values
			initColumnValues := [9]int{}
			for rowNum := 0; rowNum < 9; rowNum++ {
				initColumnValues[rowNum] = tmpResult[rowNum][i]
			}
			missingValuesInColumn, missingIndexesInColumn := findMissingValuesAndIndexes(initColumnValues)
			// Means found single missing value
			if len(missingValuesInColumn) == 1 {
				fmt.Println("===Found value in colum===")
				fmt.Printf("[%v][%v]=%v\n", missingIndexesInColumn[0], i, missingValuesInColumn)
				tmpResult[missingIndexesInColumn[0]][i] = missingValuesInColumn[0]
				fmt.Printf("\n\n")
				break
			}
			missingValuesInColumns[i] = missingValuesInColumn

			// Missing 9x9 square values
			initSquareValues := [9]int{}
			for idx := 0; idx < SIZE; idx++ {
				squareRowIdx := (idx / 3) + (i / 3 * 3)
				squareColIdx := (idx % 3) + (i % 3 * 3)
				initSquareValues[idx] = tmpResult[squareRowIdx][squareColIdx]
			}
			missingValuesInSquare, missingIndexesInSquare := findMissingValuesAndIndexes(initSquareValues)
			// Means found single missing value
			if len(missingValuesInSquare) == 1 {
				fmt.Println("===Found value in square 9x9===")
				idx := missingIndexesInSquare[0]
				squareRowIdx := (idx / 3) + (i / 3 * 3)
				squareColIdx := (idx % 3) + (i % 3 * 3)
				fmt.Printf("[%v][%v]=%v\n", squareRowIdx, squareColIdx, missingValuesInSquare)
				tmpResult[squareRowIdx][squareColIdx] = missingValuesInSquare[0]
				fmt.Printf("\n\n")
				break
			}
			missingValuesInSquares[i] = missingValuesInSquare
			foundSingleMissingValue = false
		}
	}

	return MissingValues{
		Rows:    missingValuesInRows,
		Cols:    missingValuesInColumns,
		Squares: missingValuesInSquares,
	}
}

func findMissingValuesAndIndexes(row [9]int) ([]int, []int) {
	valuesRegister := make([]int, SIZE + 1) // 10 because highest possible input value is 9
	indexesRegister := []int{}
	for i, v := range row {
		if v == 0 {
			indexesRegister = append(indexesRegister, i)
			continue
		}
		valuesRegister[v] = 1
	}
	var result []int
	for i, v := range valuesRegister {
		if v == 0 && i != 0 {
			result = append(result, i)
		}
	}
	return result, indexesRegister
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
	r := row - row % BOX
	c := col - col % BOX

	for i := r; i < r + BOX; i++ {
		for j := c; j < c + BOX; j++ {
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
		   !isInBox(board, row, col, number);
}

func solveWithBackTracking(board *[9][9]int) bool {
	for row := 0; row < SIZE; row++ {
		for col := 0; col < SIZE; col++ {
		  if board[row][col] == EMPTY {
			for number := 1; number <= SIZE; number++ {
			  if isOk(board, row, col, number) {
				board[row][col] = number;
			  
				if solveWithBackTracking(board) {
				  return true;
				} else {
				  board[row][col] = EMPTY;
				}
			  }
			}
		
			return false;
		  }
		}
	  }
	return true
}