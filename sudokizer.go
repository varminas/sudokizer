package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"time"

	"sudokizer/model"
	"sudokizer/resolver"
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
	sudokuValues.Values[0][0] = 4
	sudokuValues.Values[0][1] = 0
	sudokuValues.Values[0][2] = 2
	sudokuValues.Values[0][3] = 0
	sudokuValues.Values[0][4] = 0
	sudokuValues.Values[0][5] = 0
	sudokuValues.Values[0][6] = 0
	sudokuValues.Values[0][7] = 0
	sudokuValues.Values[0][8] = 6

	sudokuValues.Values[1][1] = 0
	sudokuValues.Values[1][2] = 0
	sudokuValues.Values[1][4] = 0
	sudokuValues.Values[1][6] = 2
	sudokuValues.Values[1][7] = 8

	sudokuValues.Values[2][0] = 8
	sudokuValues.Values[2][1] = 7
	sudokuValues.Values[2][4] = 6
	sudokuValues.Values[2][5] = 1

	sudokuValues.Values[3][0] = 1
	sudokuValues.Values[3][1] = 2
	sudokuValues.Values[3][6] = 3
	sudokuValues.Values[3][8] = 8

	sudokuValues.Values[4][0] = 0
	sudokuValues.Values[4][1] = 0
	sudokuValues.Values[4][2] = 0
	sudokuValues.Values[4][4] = 4

	sudokuValues.Values[5][0] = 3
	sudokuValues.Values[5][1] = 0
	sudokuValues.Values[5][2] = 4
	sudokuValues.Values[5][7] = 2
	sudokuValues.Values[5][8] = 5

	sudokuValues.Values[6][3] = 9
	sudokuValues.Values[6][4] = 7
	sudokuValues.Values[6][7] = 6
	sudokuValues.Values[6][8] = 3

	sudokuValues.Values[7][1] = 4
	sudokuValues.Values[7][2] = 7
	sudokuValues.Values[7][3] = 0
	sudokuValues.Values[7][4] = 0

	sudokuValues.Values[8][0] = 5
	sudokuValues.Values[8][6] = 5
	sudokuValues.Values[8][8] = 4

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
