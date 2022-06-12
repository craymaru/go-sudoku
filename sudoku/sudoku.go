package sudoku

import (
	"bytes"
	"fmt"
	"strconv"
)

type Board [9][9]int

func Pretty(board Board) string {
	var buf bytes.Buffer
	for i := 0; i < 9; i++ {
		if i%3 == 0 {
			buf.WriteString(" +-------+-------+-------+\n")
		}
		for j := 0; j < 9; j++ {
			if j%3 == 0 {
				buf.WriteString(" | ")
			}
			buf.WriteString(strconv.Itoa(board[i][j]))
			if j != 2 && j != 5 && j != 8 {
				buf.WriteString(" ")
			}
		}
		buf.WriteString(" |\n")
	}
	buf.WriteString(" +-------+-------+-------+\n")
	return buf.String()
}
func Duplicated(counters [10]int) bool {
	for num, count := range counters {
		if num == 0 {
			continue
		}
		if 1 < count {
			return true
		}

	}
	return false
}

func Verify(board Board) bool {
	// Row
	for i := 0; i < 9; i++ {
		var counts [10]int
		for j := 0; j < 9; j++ {
			counts[board[i][j]]++
		}
		if Duplicated(counts) {
			return false
		}
	}

	// Columns
	for i := 0; i < 9; i++ {
		var counts [10]int
		for j := 0; j < 9; j++ {
			counts[board[j][i]]++
		}
		if Duplicated(counts) {
			return false
		}
	}

	// Square
	for i := 0; i < 9; i += 3 {
		for j := 0; j < 9; j += 3 {
			var counts [10]int
			for row := i; row < i+3; row++ {
				for col := j; col < j+3; col++ {
					counts[board[row][col]]++
				}
			}
			if Duplicated(counts) {
				return false
			}
		}
	}
	return true
}

func Solved(board Board) bool {
	if !Verify(board) {
		fmt.Printf("Failed...")
		return false
	}

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] == 0 {
				return false
			}
		}
	}

	return true
}

func Backtrack(board *Board) bool {
	// time.Sleep(time.Microsecond * 500000)
	// fmt.Printf("%+v\n", Pretty(*board))

	if Solved(*board) {
		return true
	}

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] == 0 {
				for candidate := 9; candidate >= 1; candidate-- {
					board[i][j] = candidate
					if Verify(*board) {
						// Recursion
						if Backtrack(board) {
							return true
						}
					}
					board[i][j] = 0
				}
				return false
			}
		}
	}

	return false
}

func Sudoku() {
	b := Board{
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
	}
	fmt.Println(Pretty(b))
}
