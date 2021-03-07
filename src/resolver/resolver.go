package resolver

type SudokuValues struct {
	Values [9][9]uint8
}

func Resolve(initValues SudokuValues) SudokuValues {
	return initValues
}