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
	sudokuValues.Values[0][0] = 9
	sudokuValues.Values[0][1] = 1
	sudokuValues.Values[0][2] = 5
	sudokuValues.Values[0][3] = 3
	sudokuValues.Values[0][4] = 4
	sudokuValues.Values[0][8] = 7

	sudokuValues.Values[1][1] = 8
	sudokuValues.Values[1][2] = 3
	sudokuValues.Values[1][4] = 9
	sudokuValues.Values[1][5] = 7
	sudokuValues.Values[1][7] = 5

	sudokuValues.Values[2][0] = 4
	sudokuValues.Values[2][1] = 2
	sudokuValues.Values[2][2] = 7
	sudokuValues.Values[2][7] = 1

	sudokuValues.Values[3][2] = 2
	sudokuValues.Values[3][3] = 6
	sudokuValues.Values[3][4] = 8
	sudokuValues.Values[3][6] = 4

	sudokuValues.Values[4][0] = 7
	sudokuValues.Values[4][2] = 4
	sudokuValues.Values[4][3] = 2
	sudokuValues.Values[4][5] = 9

	sudokuValues.Values[5][2] = 8
	sudokuValues.Values[5][4] = 3
	sudokuValues.Values[5][5] = 4
	sudokuValues.Values[5][6] = 1
	sudokuValues.Values[5][7] = 6

	sudokuValues.Values[6][0] = 8
	sudokuValues.Values[6][7] = 4

	sudokuValues.Values[7][2] = 9
	sudokuValues.Values[7][6] = 7
	sudokuValues.Values[7][7] = 2
	sudokuValues.Values[7][8] = 6

	sudokuValues.Values[8][1] = 5
	sudokuValues.Values[8][2] = 6
	sudokuValues.Values[8][5] = 3
	sudokuValues.Values[8][6] = 8
	sudokuValues.Values[8][8] = 1


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
			// checkInt(formValueInt)

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
	fmt.Printf("BEGIN of resolving: \t %s \n", timeStart)

	values := resolver.Resolve(initValues)

	timeEnd := time.Now()
	fmt.Printf("END of resolving: \t %s \n", timeEnd)

	processingTime := timeEnd.Sub(timeStart)
	
	return model.SudokuSolution{
		Values: values, 
		TimeStart: timeStart.Format(time.RFC3339), 
		TimeEnd: timeEnd.Format(time.RFC3339),
		ProcessingTime: processingTime,
	}
}
