package resolver

import (
	"sudokizer/model"
)

func Resolve(initValues *model.SudokuValues) model.SudokuValues {
	return *initValues
}