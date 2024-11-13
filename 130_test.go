package main

import (
	"reflect"
	"testing"
)

// Time complexity: O(m * n), m and n are the dimensions of the board.
// Space complexity: O(m * n), for recursion call stack.
// https://claude.ai/chat/3b70d4a4-73cd-4711-b678-28c0914437f2
func solve130(brd [][]byte) {
	// Surrounded Regions
	// Given an m x n matrix `brd`.
	// Each cell is region 'O' or surround 'X'.
	// Capture regions by surrounding.
	// Conditions:
	//	* Connect is horizontal or vertical
	//	* Region is one or more 'O' connected
	//	* Surround is one or more 'O' surrounded by 'X'
	//	* Cannot surround a 'O' on a board edge.
	//	* Surround is performed by switching 'O' to 'X'
	// Key insights:
	//	* Easier to identify edge 'O' first.
	//	* Mark edge 'O' as safe 'S'.
	//	* Perform surround.
	//	* Restore safe 'S' to region 'O'
	// Use depth-first search (DFS) with recursion for safe 'S'.
	// Use helper function to support recursion.

	// Check input min edge cases.
	if len(brd) == 0 || len(brd[0]) == 0 {
		return
	}

	// Initialize variables.
	rows, cols := len(brd), len(brd[0])

	// Mark edge column 'O' as 'S'
	for row := 0; row < rows; row++ {
		markSafe130(brd, row, 0)
		markSafe130(brd, row, cols-1)
	}
	// Mark edge row 'O' as 'S'
	for col := 0; col < cols; col++ {
		markSafe130(brd, 0, col)
		markSafe130(brd, rows-1, col)
	}

	// Update board.
	// Transform 'O' to 'X'.
	// Transform 'S' to 'O'.
	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			if brd[row][col] == 'O' {
				brd[row][col] = 'X'
			} else if brd[row][col] == 'S' {
				brd[row][col] = 'O'
			}
		}
	}
}

func markSafe130(brd [][]byte, row, col int) {
	// Check bounds and region 'O'
	if row < 0 || row >= len(brd) ||
		col < 0 || col >= len(brd[0]) ||
		brd[row][col] != 'O' {
		return
	}

	// Mark region 'O' as safe 'S'
	brd[row][col] = 'S'

	// Recursively mark connected region.
	markSafe130(brd, row-1, col) // up
	markSafe130(brd, row+1, col) // down
	markSafe130(brd, row, col-1) // left
	markSafe130(brd, row, col+1) //right
}

// For testing
func TestSolve130(t *testing.T) {
	cases := []struct {
		name     string
		input    [][]byte
		expected [][]byte
	}{
		{
			name: "Example 1",
			input: [][]byte{
				{'X', 'X', 'X', 'X'},
				{'X', 'O', 'O', 'X'},
				{'X', 'X', 'O', 'X'},
				{'X', 'O', 'X', 'X'},
			},
			expected: [][]byte{
				{'X', 'X', 'X', 'X'},
				{'X', 'X', 'X', 'X'},
				{'X', 'X', 'X', 'X'},
				{'X', 'O', 'X', 'X'},
			},
		},
		{
			name:     "Example 2",
			input:    [][]byte{{'X'}},
			expected: [][]byte{{'X'}},
		},
		{
			name:     "Empty board",
			input:    [][]byte{},
			expected: [][]byte{},
		},
		{
			name: "All edge O's",
			input: [][]byte{
				{'O', 'O', 'O'},
				{'O', 'O', 'O'},
				{'O', 'O', 'O'},
			},
			expected: [][]byte{
				{'O', 'O', 'O'},
				{'O', 'O', 'O'},
				{'O', 'O', 'O'},
			},
		},
		{
			name: "Center O surrounded",
			input: [][]byte{
				{'X', 'X', 'X'},
				{'X', 'O', 'X'},
				{'X', 'X', 'X'},
			},
			expected: [][]byte{
				{'X', 'X', 'X'},
				{'X', 'X', 'X'},
				{'X', 'X', 'X'},
			},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			// Make a deep copy since solve modifies input
			input := make([][]byte, len(tc.input))
			for i := range tc.input {
				input[i] = make([]byte, len(tc.input[i]))
				copy(input[i], tc.input[i])
			}

			solve130(input)

			// Compare results
			if !reflect.DeepEqual(input, tc.expected) {
				t.Errorf("got %v, want %v", input, tc.expected)
			}
		})
	}
}
