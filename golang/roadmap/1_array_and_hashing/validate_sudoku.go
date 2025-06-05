package array_and_hashing

/*
Valid Sudoku
You are given a a 9 x 9 Sudoku board board. A Sudoku board is valid if the following rules are followed:

Each row must contain the digits 1-9 without duplicates.
Each column must contain the digits 1-9 without duplicates.
Each of the nine 3 x 3 sub-boxes of the grid must contain the digits 1-9 without duplicates.
Return true if the Sudoku board is valid, otherwise return false

Note: A board does not need to be full or be solvable to be valid.

Example 1:



Input: board =
[["1","2",".",".","3",".",".",".","."],
 ["4",".",".","5",".",".",".",".","."],
 [".","9","8",".",".",".",".",".","3"],
 ["5",".",".",".","6",".",".",".","4"],
 [".",".",".","8",".","3",".",".","5"],
 ["7",".",".",".","2",".",".",".","6"],
 [".",".",".",".",".",".","2",".","."],
 [".",".",".","4","1","9",".",".","8"],
 [".",".",".",".","8",".",".","7","9"]]

Output: true
Example 2:

Input: board =
[["1","2",".",".","3",".",".",".","."],
 ["4",".",".","5",".",".",".",".","."],
 [".","9","1",".",".",".",".",".","3"],
 ["5",".",".",".","6",".",".",".","4"],
 [".",".",".","8",".","3",".",".","5"],
 ["7",".",".",".","2",".",".",".","6"],
 [".",".",".",".",".",".","2",".","."],
 [".",".",".","4","1","9",".",".","8"],
 [".",".",".",".","8",".",".","7","9"]]

Output: false
Explanation: There are two 1's in the top-left 3x3 sub-box.

Constraints:

board.length == 9
board[i].length == 9
board[i][j] is a digit 1-9 or '.'.
*/

func ValidateSudoku(board [][]string) bool {
	for i := 0; i < 9; i++ {
		line := make(map[string]bool)
		for j := 0; j < 9; j++ {
			if board[i][j] == "." {
				continue
			}

			if _, found := line[board[i][j]]; found {
				return false
			}

			line[board[i][j]] = true
		}
	}

	for i := 0; i < 9; i++ {
		column := make(map[string]bool)
		for j := 0; j < 9; j++ {
			if board[j][i] == "." {
				continue
			}

			if _, found := column[board[j][i]]; found {
				return false
			}

			column[board[j][i]] = true
		}
	}

	line := 0
	column := 0

	for line < 9 {
		for column < 9 {
			if isValid := validateBox(board, line, column); !isValid {
				return false
			}

			column = column + 3
		}

		column = 0
		line = line + 3
	}

	return true
}

func validateBox(board [][]string, line, column int) bool {
	box := make(map[string]bool)
	for i := line; i < line+3; i++ {
		for j := column; j < column+3; j++ {
			if board[i][j] == "." {
				continue
			}
			if _, ok := box[board[i][j]]; ok {
				return false
			}
			box[board[i][j]] = true
		}
	}

	return true
}
