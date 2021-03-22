package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/varminas/sudokizer/model"
	"github.com/varminas/sudokizer/resolver"
)

var t = template.Must(template.ParseFiles("templates/index.html"))
var tSolution = template.Must(template.ParseFiles("templates/index.html", "templates/solution.html"))

type AppState struct {
	Inputs model.SudokuValues
	Solution model.SudokuSolution
}

func newAppState() AppState {
	return AppState{}
}

var appState = newAppState()

func main() {
	fmt.Println("Starting Sudokizer at port 8080")

	http.HandleFunc("/home", homeHandler)
	http.HandleFunc("/solution", solutionHandler)
	http.HandleFunc("/", mainHandler)

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	err := http.ListenAndServe("localhost:8080", nil)
	log.Fatal(err)
}

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
	http.Redirect(writer, request, "/home", http.StatusSeeOther)
}

func homeHandler(writer http.ResponseWriter, request *http.Request) {
	sudokuValues := model.SudokuValues{}
	// TODO: delete Init values 
	sudokuValues.Values = [9][9]int{
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

	appState.Inputs = sudokuValues

	var err = t.Execute(writer, appState)
	check(err)
}

func solutionHandler(writer http.ResponseWriter, request *http.Request) {
	initSudokuValues := model.SudokuValues{}

	for i := 0; i < 9; i++ {
		for j  := 0; j < 9; j++ {
			formValueName := fmt.Sprintf("value-%d-%d", i, j)
			formValueStr := request.FormValue(formValueName)
			formValueInt, err := strconv.ParseInt(formValueStr, 10, 8)
			if err != nil {
				formValueInt = 0
			}

			initSudokuValues.Values[i][j] = int(formValueInt)
		}
	}

	appState.Inputs = initSudokuValues
	appState.Solution = resolve(initSudokuValues)

	var err = tSolution.Execute(writer, appState)
	check(err)
}

func resolve(initValues model.SudokuValues) model.SudokuSolution {
	timeStart := time.Now()
	fmt.Println("-----------------------------------")
	fmt.Printf("BEGIN of resolving: \t %s \n", timeStart)

	values := resolver.Resolve(initValues)

	timeEnd := time.Now()
	fmt.Printf("END of resolving: \t %s \n", timeEnd)
	fmt.Println("-----------------------------------\n\n")

	processingTime := timeEnd.Sub(timeStart)
	
	return model.SudokuSolution{
		Values: values, 
		TimeStart: timeStart.Format(time.RFC3339), 
		TimeEnd: timeEnd.Format(time.RFC3339),
		ProcessingTime: processingTime,
	}
}
