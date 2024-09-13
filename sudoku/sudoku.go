package main

import (
	"fmt"
	"os"
)

var board = [9][9]int{}

func main() {
	// Zero out board
	for row := 0; row < 9; row++ {
		for col := 0; col < 9; col++ {
			board[row][col] = 0
		}
	}
	Args := os.Args[1:]
	// Check to make sure there are 9 rows
	if len(Args) < 9 {
		fmt.Println("Error")
		return
	}
	// Check to make sure each row contains 9 columns
	for row, arg := range Args {
		if len(arg) < 9 {
			fmt.Println("Error")
			return
		}
		for col, r := range arg {
			if r == '.' {
				continue
			} else if r < 48 || r > 57 {
				fmt.Println("Error")
				return
			} else {
				board[row][col] = int(r) - 48
			}
		}
	}
	if Solve(0, 0) {
		Draw()
		return
	} else {
		fmt.Println("Error")
		return
	}
}

func Solve(row, col int) bool {
	// fmt.Printf("row: %v, col: %v\n", row, col)
	if row == 9 {
		return true
	}
	if board[row][col] != 0 {
		return Solve(Next(row, col))
	} else {
		for i := 0; i < 9; i++ {
			value := i + 1
			if CanPut(row, col, value) {
				board[row][col] = value
				if Solve(Next(row, col)) {
					return true
				}
				board[row][col] = 0
			}
		}
		return false
	}
}

func CanPut(row, col, value int) bool {
	// Checks if the value is already
	// in the horizontal line (column)
	for i := 0; i < 9; i++ {
		if board[row][i] == value {
			return false
		}
	}
	// Checks if the value is already
	// in the vertical line (row)
	for j := 0; j < 9; j++ {
		if board[j][col] == value {
			return false
		}
	}
	// Checks if the value is already
	// in the 3*3 square
	srow, scol := int(row/3)*3, int(col/3)*3
	for k := 0; k < 3; k++ {
		for l := 0; l < 3; l++ {
			if board[srow+k][scol+l] == value {
				return false
			}
		}
	}
	return true
}

func Next(row, col int) (int, int) {
	nextRow, nextCol := row, (col+1)%9
	if nextCol == 0 {
		nextRow = row + 1
	}
	return nextRow, nextCol
}

func Draw() {
	for row := 0; row < 9; row++ {
		for col := 0; col < 9; col++ {
			fmt.Printf("%d", board[row][col])
			if col != 8 {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}
