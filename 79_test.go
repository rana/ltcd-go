package main

import "testing"

// Time complexity: O(m * n * 4^l), m and n are the matrix board dimensions. l is the length of the word.
// Space complexity: O(l), for recursive call stack.
// https://claude.ai/chat/c9e3517c-2eba-4377-bc2b-d83078db1d64
func exist79(brd [][]byte, wrd string) bool {
	// Word Search
	// Given a matrix brd with dimensions m x n.
	// Given a string wrd.
	// Determine whether the word exists in the board.
	// Conditions:
	//	* Horizontal movement
	//	* Vertical movement
	// Use backtracking, recursion, and early pruning.
	// Use in-place board updating for visited cells.

	// Check input min edge cases.
	if len(brd) == 0 || len(brd[0]) == 0 || len(wrd) == 0 {
		return false
	}

	// Initialize variables.
	hgt, wdt := len(brd), len(brd[0])

	// Start from each cell of the board looking for the word.
	for row := range hgt {
		for col := range wdt {
			if dfsFndPth79(brd, row, col, wrd, 0) {
				return true
			}
		}
	}

	return false
}

var dirs79 = [][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

// dfsFndPth79 performs a DFS searching for wrd.
func dfsFndPth79(brd [][]byte, row, col int, wrd string, idx int) bool {
	// Check early pruning: move beyond bounds or char mis-match.
	if row < 0 || col < 0 || row >= len(brd) || col >= len(brd[0]) ||
		brd[row][col] != wrd[idx] {
		return false
	}

	// Check base case: found matching word.
	if idx == len(wrd)-1 {
		return true
	}

	// Mark current cell visited, in-place.
	tmp := brd[row][col]
	brd[row][col] = '#'

	// Search all four directions with recursive backtracking.

	for _, dir := range dirs79 {
		nxtRow, nxtCol := row+dir[0], col+dir[1]
		if dfsFndPth79(brd, nxtRow, nxtCol, wrd, idx+1) {
			// Backtrack visited marker.
			brd[row][col] = tmp
			return true
		}
	}

	// Backtrack visited marker.
	brd[row][col] = tmp
	return false
}

func TestExist79(t *testing.T) {
	tests := []struct {
		name     string
		board    [][]byte
		word     string
		expected bool
	}{
		{
			name: "Example 1",
			board: [][]byte{
				{'A', 'B', 'C', 'E'},
				{'S', 'F', 'C', 'S'},
				{'A', 'D', 'E', 'E'},
			},
			word:     "ABCCED",
			expected: true,
		},
		{
			name: "Example 2",
			board: [][]byte{
				{'A', 'B', 'C', 'E'},
				{'S', 'F', 'C', 'S'},
				{'A', 'D', 'E', 'E'},
			},
			word:     "SEE",
			expected: true,
		},
		{
			name: "Example 3",
			board: [][]byte{
				{'A', 'B', 'C', 'E'},
				{'S', 'F', 'C', 'S'},
				{'A', 'D', 'E', 'E'},
			},
			word:     "ABCB",
			expected: false,
		},
		{
			name:     "Empty board",
			board:    [][]byte{},
			word:     "A",
			expected: false,
		},
		{
			name: "Single cell board",
			board: [][]byte{
				{'A'},
			},
			word:     "A",
			expected: true,
		},
		{
			name: "Word longer than board",
			board: [][]byte{
				{'A', 'B'},
				{'C', 'D'},
			},
			word:     "ABCDEF",
			expected: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			board := make([][]byte, len(tc.board))
			for i := range tc.board {
				board[i] = make([]byte, len(tc.board[i]))
				copy(board[i], tc.board[i])
			}

			got := exist79(board, tc.word)
			if got != tc.expected {
				t.Errorf("exist() = %v, want %v", got, tc.expected)
			}
		})
	}
}
