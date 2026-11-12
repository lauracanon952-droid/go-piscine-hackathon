package main

import (
	"fmt"
	"os"
)

type Row []int
type Grid []Row

func main() {
	if len(os.Args) != 10 {
		fmt.Println("Error")
		return
	}

	board := make(Grid, 9)
	for i := range board {
		board[i] = make(Row, 9)
	}

	for r := 0; r < 9; r++ {
		row := os.Args[r+1]
		if len(row) != 9 {
			fmt.Println("Error")
			return
		}
		for c := 0; c < 9; c++ {
			ch := row[c]
			if ch == '.' {
				board[r][c] = 0
			} else if ch >= '1' && ch <= '9' {
				num := int(ch - '0')
				if !isValid(board, r, c, num) {
					fmt.Println("Error")
					return
				}
				board[r][c] = num
			} else {
				fmt.Println("Error")
				return
			}
		}
	}

	if solve(board) {
		printBoard(board)
	} else {
		fmt.Println("Error")
	}
}

func isValid(board Grid, row, col, num int) bool {
	for i := 0; i < 9; i++ {
		if board[row][i] == num || board[i][col] == num {
			return false
		}
		subRow := 3*(row/3) + i/3
		subCol := 3*(col/3) + i%3
		if board[subRow][subCol] == num {
			return false
		}
	}
	return true
}

func solve(board Grid) bool {
	for r := 0; r < 9; r++ {
		for c := 0; c < 9; c++ {
			if board[r][c] == 0 {
				for num := 1; num <= 9; num++ {
					if isValid(board, r, c, num) {
						board[r][c] = num
						if solve(board) {
							return true
						}
						board[r][c] = 0
					}
				}
				return false
			}
		}
	}
	return true
}

func printBoard(board Grid) {
	for r := 0; r < 9; r++ {
		for c := 0; c < 9; c++ {
			fmt.Print(board[r][c])
			if c < 8 {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}
