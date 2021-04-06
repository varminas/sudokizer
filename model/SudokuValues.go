package model

import "time"

type SudokuValues struct {
	Values [9][9]int
}

type SudokuSolution struct {
	Values SudokuValues
	TimeStart string
	TimeEnd string
	ProcessingTime time.Duration
}

type Algorithm string

const(
	BackTracking Algorithm = "BackTracking"
	DancingLinks Algorithm = "DancingLinks"
)