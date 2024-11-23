package main

import "testing"

func exist79b(brd [][]byte, wrd string) bool {
	// Word Search
	// Given a byte matrix brd with dimensions m x n.
	// Given a string wrd.
	// Determine whether wrd exists in brd.
	// Return true if conditions met.
	// Conditions:
	// * Horizontal movement
	// * Vertical movement
	// Use in-place visited marking with backtracking.
	// Use backtracking and recursion.
	// Use a separate function for recursion.

	// Check input min edge cases.
	if len(brd) == 0 || len(brd[0]) == 0 || len(wrd) == 0 {
		return false
	}

	// Initialize variables.
	hgt, wdt := len(brd), len(brd[0])

	// Start search at each cell of board.
	for row := range hgt {
		for col := range wdt {
			if dfsFndPth(brd, row, col, wrd, 0) {
				return true
			}
		}
	}

	return false
}

var dirs79b = [][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

func dfsFndPth(brd [][]byte, row, col int, wrd string, idx int) bool {
	// Check early pruning: move out of bounds, char mis-match.
	if row < 0 || col < 0 || row >= len(brd) || col >= len(brd[0]) ||
		brd[row][col] != wrd[idx] {
		return false
	}

	// Check base case.
	// Check whether word found on board.
	if idx == len(wrd)-1 {
		return true
	}

	// Mark board visited in-place.
	tmp := brd[row][col]
	brd[row][col] = '#'

	// Search board in all four directions.
	for _, dir := range dirs79b {
		nxtRow, nxtCol := row+dir[0], col+dir[1]
		// Recurse next index.
		if dfsFndPth(brd, nxtRow, nxtCol, wrd, idx+1) {
			// Backtrack in-place visit marker.
			brd[row][col] = tmp
			return true
		}
	}

	// Backtrack in-place visit marker.
	brd[row][col] = tmp
	return false
}

func TestExist79b(t *testing.T) {
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

			got := exist79b(board, tc.word)
			if got != tc.expected {
				t.Errorf("exist() = %v, want %v", got, tc.expected)
			}
		})
	}
}
