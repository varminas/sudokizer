package resolver

import (
	"fmt"
	"testing"
	"sudokizer/model"
)

func TestResolve(t *testing.T) {
	grid := model.SudokuValues{}
	grid.Values = [9][9]int{
		{0, 1, 0, 0, 0, 0, 7, 0, 3},
		{3, 0, 0, 0, 0, 8, 0, 6, 0},
		{9, 0, 0, 0, 1, 3, 0, 0, 0},
		{0, 0, 4, 0, 0, 0, 0, 0, 5},
		{0, 0, 0, 2, 9, 4, 0, 0, 0},
		{7, 0, 0, 0, 0, 0, 1, 0, 0},
		{0, 0, 0, 3, 8, 0, 0, 0, 9},
		{0, 5, 0, 6, 0, 0, 0, 0, 4},
		{2, 0, 6, 0, 0, 0, 0, 1, 0},
	}

	want := [9][9]int{
		{4, 1, 8, 5, 2, 6, 7, 9, 3},
		{3, 7, 2, 9, 4, 8, 5, 6, 1},
		{9, 6, 5, 7, 1, 3, 4, 2, 8},
		{6, 2, 4, 1, 3, 7, 9, 8, 5},
		{5, 8, 1, 2, 9, 4, 3, 7, 6},
		{7, 9, 3, 8, 6, 5, 1, 4, 2},
		{1, 4, 7, 3, 8, 2, 6, 5, 9},
		{8, 5, 9, 6, 7, 1, 2, 3, 4},
		{2, 3, 6, 4, 5, 9, 8, 1, 7},
	}

	got := Resolve(grid).Values
	if got != want {
		t.Errorf(errorString(grid, got, want))
	}
}

func errorString(grid model.SudokuValues, got [9][9]int, want [9][9]int) string {
	return fmt.Sprintf("Resolve(%#v) \n\ngot= %#v \n\nwant=%#v", grid, got, want)
}