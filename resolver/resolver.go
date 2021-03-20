package resolver

import (
	"fmt"
	"math/rand"
	// "sort"
	"sudokizer/model"
)

type MissingValues struct {
	Rows [][]int
	Cols [][]int
	Squares [][]int
}

func Resolve(initSudokuValues model.SudokuValues) model.SudokuValues {
	tmpResult := initSudokuValues.Values
	// rand.Seed(time.Now().UnixNano())

	missingValues := resolveSingleMissingValues(&tmpResult)
	fmt.Printf("missRows %#v \n\n", missingValues.Rows)
	fmt.Printf("missColu %#v \n\n", missingValues.Cols)
	fmt.Printf("missSqua %#v \n\n", missingValues.Squares)

	resolveBySmallestArray(&tmpResult, missingValues)

	return model.SudokuValues{Values: tmpResult}
}

func resolveBySmallestArray(tmpResult* [9][9]int, missingValues MissingValues) {
	
}

func resolveSingleMissingValues(tmpResult* [9][9]int) MissingValues {
	missingValuesInRows := make([][]int, 9)
	missingValuesInColumns := make([][]int, 9)
	missingValuesInSquares := make([][]int, 9)

	foundSingleMissingValue := true
	for ; foundSingleMissingValue == true; {
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
			for idx := 0; idx < 9; idx++ {
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
		Rows: missingValuesInRows,
		Cols: missingValuesInColumns,
		Squares: missingValuesInSquares,
	}
}

func findMissingValuesAndIndexes(row [9]int) ([]int, []int) {
	valuesRegister := make([]int, 10) // 10 because highest possible input value is 9
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

func randFromRange() int {
	return rand.Intn(9 - 1)
}
