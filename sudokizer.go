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

type SudokuSolution struct {
	Values model.SudokuValues
	TimeStart string
	TimeEnd string
	ProcessingTime time.Duration
}

type AppState struct {
	Inputs model.SudokuValues
	Solution SudokuSolution
}

var appState = AppState{}

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
	for i := 0; i < len(sudokuValues.Values); i++ {
		for j := 0; j < len(sudokuValues.Values[0]); j++ {
			sudokuValues.Values[i][j] = 1
		}
	}
	appState.Inputs = sudokuValues
	// fmt.Println("state", appState)

	var err = t.Execute(writer, appState)
	check(err)
}

func solutionHandler(writer http.ResponseWriter, request *http.Request) {
	initSudokuValues := &model.SudokuValues{}

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

	appState.Solution = resolve(initSudokuValues)

	var err = tSolution.Execute(writer, appState)
	check(err)
}

func resolve(initValues *model.SudokuValues) SudokuSolution {
	timeStart := time.Now()
	fmt.Printf("BEGIN of resolving: \t %s \n", timeStart)

	values := resolver.Resolve(initValues)

	timeEnd := time.Now()
	fmt.Printf("END of resolving: \t %s \n", timeEnd)
	
	processingTime := timeEnd.Sub(timeStart)
	
	return SudokuSolution{
		Values: values, 
		TimeStart: timeStart.Format(time.RFC3339), 
		TimeEnd: timeEnd.Format(time.RFC3339),
		ProcessingTime: processingTime,
	}
}
