# sudokizer
## Sudoku resolver written in GO language

This is a simple web service running on port 8080. It serves sudoku input and solves puzzle by clicking "Resolve" button.

**Note**: Input puzzle must be valid, i.e. resolvale.

## How to use the program
In the console run:
```sh
go run ./main.go
```
Open your favorite browser and type __http://localhost:8080__. Enter input puzzle and click button __Resolve__.

__NOTE:__ Currenly only BackTracking algorithm is implemented!
