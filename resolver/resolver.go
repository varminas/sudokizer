package resolver

import (
	"fmt"
	"math/rand"
	// "sort"
	"sudokizer/model"
)

func Resolve(initSudokuValues model.SudokuValues) model.SudokuValues {
	initValues := initSudokuValues.Values
	// rand.Seed(time.Now().UnixNano())

	missingValuesInRows := make([][]int, 10)
	missingValuesInColumns := make([][]int, 10)
	missingValuesInSquares := make([][]int, 10)

	for i, row := range initValues {
		// sort.Ints(row[:])
		// Missing row values
		missingValuesInRow := findMissingValues(row)
		missingValuesInRows[i] = missingValuesInRow

		// Missing column values
		initColumnValues := [9]int{}
		for rowNum := 0; rowNum < 9; rowNum++ {
			initColumnValues[rowNum] = initValues[rowNum][i]
		}
		missingValuesInColumn := findMissingValues(initColumnValues)
		missingValuesInColumns[i] = missingValuesInColumn

		// Missing 9x9 square values
		initSquareValues := [9]int{}
		for idx := 0; idx < 9; idx++ {
			squareRowIdx := (idx / 3) + (i / 3 * 3)
			squareColIdx := (idx % 3) + (i % 3 * 3)
			initSquareValues[idx] = initValues[squareRowIdx][squareColIdx]
		}
		missingValuesInSquare := findMissingValues(initSquareValues)
		missingValuesInSquares[i] = missingValuesInSquare
	}
	fmt.Printf("missRows %#v \n\n", missingValuesInRows)
	fmt.Printf("missColu %#v \n\n", missingValuesInColumns)
	fmt.Printf("missSqua %#v \n\n", missingValuesInSquares)

	return model.SudokuValues{Values: initValues}
}


func findMissingValues(row [9]int) []int {
	register := make([]int, 10) // 10 because highest possible input value is 9
	for _, v := range row {
		if v == 0 {
			continue
		}
		register[v] = 1
	}
	var result []int
	for i, v := range register {
		if v == 0 && i != 0 {
			result = append(result, i)
		}
	}
	return result
}

func randFromRange() int {
	return rand.Intn(9 - 1)
}
