package main

import "testing"

// Time complexity: O(1)
// Space complexity: O(1)
// https://claude.ai/chat/85478412-a046-44b5-bedb-30ffab80d68e
func isValidSudoku36(brd [][]byte) bool {
	// Valid Sudoku
	// Given a matrix `brd`.
	// Determine whether the matrix is a valid Sudoku board.
	// Return true if conditions met.
	// Condtions:
	//	* Each row contains unique digits 1-9.
	//	* Each column contains unique digits 1-9.
	//	* Each box contains unique digits 1-9.
	//	* Characters are 1-9, '.'
	// Use hashsets for each row, column, and box to solve.

	brdLen := len(brd)

	// Initialize array of maps.
	rows := make([]map[byte]bool, brdLen)
	cols := make([]map[byte]bool, brdLen)
	boxs := make([]map[byte]bool, brdLen)

	// Initialize each map.
	for idx := range brdLen {
		rows[idx] = make(map[byte]bool)
		cols[idx] = make(map[byte]bool)
		boxs[idx] = make(map[byte]bool)
	}

	// Iterate through each board cell.
	for row := range brdLen {
		for col := range brdLen {
			// Get the current board character.
			val := brd[row][col]

			// Check for dot.
			if val == '.' {
				// Skip dot
				continue
			}

			// Calculate box index (0-8) from column and row.
			box := (row/3)*3 + col/3

			// Check whether digit has already been seen.
			if rows[row][val] || cols[col][val] || boxs[box][val] {
				// Invalid Sudoku board.
				return false
			}

			// Mark digit as seen.
			rows[row][val] = true
			cols[col][val] = true
			boxs[box][val] = true
		}
	}

	return true
}

func TestIsValidSudoku36(t *testing.T) {
	tests := []struct {
		name  string
		board [][]byte
		want  bool
	}{
		{
			name: "valid board",
			board: [][]byte{
				{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
				{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
				{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
				{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
				{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
				{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
				{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
				{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
				{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
			},
			want: true,
		},
		{
			name: "invalid board",
			board: [][]byte{
				{'8', '3', '.', '.', '7', '.', '.', '.', '.'},
				{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
				{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
				{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
				{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
				{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
				{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
				{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
				{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
			},
			want: false,
		},
		{
			name: "empty board",
			board: [][]byte{
				{'.', '.', '.', '.', '.', '.', '.', '.', '.'},
				{'.', '.', '.', '.', '.', '.', '.', '.', '.'},
				{'.', '.', '.', '.', '.', '.', '.', '.', '.'},
				{'.', '.', '.', '.', '.', '.', '.', '.', '.'},
				{'.', '.', '.', '.', '.', '.', '.', '.', '.'},
				{'.', '.', '.', '.', '.', '.', '.', '.', '.'},
				{'.', '.', '.', '.', '.', '.', '.', '.', '.'},
				{'.', '.', '.', '.', '.', '.', '.', '.', '.'},
				{'.', '.', '.', '.', '.', '.', '.', '.', '.'},
			},
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isValidSudoku36(tt.board); got != tt.want {
				t.Errorf("isValidSudoku() = %v, want %v", got, tt.want)
			}
		})
	}
}
