package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/sudokizer/src/resolver"
)

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func checkInt(value int64) {
	if value < 1 || value > 9 {
		log.Fatal("Value must be between 1 and 9")
	}
}

func mainHandler(writer http.ResponseWriter, request *http.Request) {
	html, err := template.ParseFiles("templates/input.html")
	check(err)

	sudokuValues := resolver.SudokuValues{}
	for i := 0; i < len(sudokuValues.Values); i++ {
		for j := 0; j < len(sudokuValues.Values[0]); j++ {
			sudokuValues.Values[i][j] = 1
		}
	}
	err = html.Execute(writer, sudokuValues)
	check(err)
}

func resolve(initValues resolver.SudokuValues) resolver.SudokuValues {
	fmt.Printf("BEGIN of resolving: \t %s \n", time.Now())
	result := resolver.Resolve(initValues)
	fmt.Printf("END of resolving: \t %s \n", time.Now())
	return result
}

func resolveHandler(writer http.ResponseWriter, request *http.Request) {
	initSudokuValues := resolver.SudokuValues{}

	for i := 0; i < 9; i++ {
		for j  := 0; j < 9; j++ {
			formValueName := fmt.Sprintf("value-%d-%d", i, j)
			formValueStr := request.FormValue(formValueName)
			formValueInt, err := strconv.ParseInt(formValueStr, 10, 8)

			check(err)
			checkInt(formValueInt)

			initSudokuValues.Values[i][j] = uint8(formValueInt)
		}
	}

	sudokuValues := resolve(initSudokuValues)

	html, err := template.ParseFiles("templates/solution.html")
	check(err)
	err = html.Execute(writer, sudokuValues)
	check(err)
}

func main() {
	fmt.Println("Starting Sudokizer")

	http.HandleFunc("/main", mainHandler)
	http.HandleFunc("/resolve", resolveHandler)

	err := http.ListenAndServe("localhost:8080", nil)
	log.Fatal(err)
}
