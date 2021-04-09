package resolver

import (
	"github.com/varminas/sudokizer/model"
)

var SIZE int = 9
var BOX int = 3
var EMPTY int = 0

type Resolver interface {
	Solve(board *[9][9]int) bool
}

func Resolve(initSudokuValues model.SudokuValues, algorithm model.Algorithm) model.SudokuValues {
	tmpResult := initSudokuValues.Values

	var resolver Resolver

	// Apply BackTracking algorithm
	if algorithm == model.BackTracking {
		resolver = BackTrackingAlgorithm("BackTracking")
	}

	if algorithm == model.DancingLinks {
		resolver = DancingLinksAlgorithm("BackTracking")
	}

	resolver.Solve(&tmpResult)

	return model.SudokuValues{Values: tmpResult}
}
