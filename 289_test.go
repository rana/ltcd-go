package main

import "testing"

// https://claude.ai/chat/767089ff-3f10-4d62-9272-9f8dc6cb630a
func gameOfLife(brd [][]int) {
	// Game of Life
	// Given an MxN integer matrix `brd`.
	// Update the matrix to a next state based on rules.
	// Rules:
	//	* Integer values are 0 or 1.
	// 	* Each element interacts with eight neighbors (horizontal, vertical, diagonal)
	//	1. Live cell <2 live heighbors dies
	//	2. Live cell with 2-3 live neighbors survives
	//	3. Live cell with >3 live neighbors dies
	//	4. Dead cell with exactly 3 live neighbors become live
	// Solve in-place with additional element states.
	// Cell States:
	// 0: Dead cell
	// 1: Live cell
	// 2: Live cell (1) that will become dead (0)
	// 3: Dead cell (0) that will become live (1)
	// Use a two-pass approach.
	//	* First pass determines new state.
	//	* Second pass sets new state.
	// Counting live neighbors is important.

	// Check input minium edge cases.
	if len(brd) == 0 || len(brd[0]) == 0 {
		return
	}

	// Get board dimensions.
	rows := len(brd)
	cols := len(brd[0])

	// Define neighbor directions.
	dirs := []Point{
		{-1, -1}, {-1, 0}, {-1, 1}, // Top row
		{0, -1}, {0, 1}, // Middle row
		{1, -1}, {1, 0}, {1, 1}, // Bottom row
	}

	// First pass: Mark cell that will change state.
	// Counting live neighbors is focus.
	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			cur := Point{row, col}
			lvn := 0

			// Count live neighbors in all 8 directions.
			for _, dir := range dirs {
				nbr := Point{cur.row + dir.row, cur.col + dir.col}
				if inBounds(nbr, brd) && isLive(brd[nbr.row][nbr.col]) {
					lvn++
				}
			}

			// Set intermediate states.
			if isLive(brd[cur.row][cur.col]) {
				// Rules 1 & 3: Live cell dies
				if lvn < 2 || lvn > 3 {
					brd[cur.row][cur.col] = 2 // Mark for death
				}
			} else if lvn == 3 {
				// Rule 4: Dead cell becomes alive
				brd[cur.row][cur.col] = 3 // Mark for life
			}
		}
	}

	// Second pass: Update intermediate state to final values.
	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			switch brd[row][col] {
			case 2:
				brd[row][col] = 0 // Cell dies
			case 3:
				brd[row][col] = 1 // Cell born
			}
		}
	}
}

// Write some helper function for readability.

// isLive return true if cel represents a live state (1 or 2)
func isLive(cel int) bool {
	return cel == 1 || cel == 2
}

// inBounds checks if pnt is within mtx bounds.
func inBounds(pnt Point, mtx [][]int) bool {
	return pnt.row >= 0 && pnt.row < len(mtx) &&
		pnt.col >= 0 && pnt.col < len(mtx[0])
}

type Point struct {
	row int
	col int
}

// Test cases for the Game of Life implementation
func TestGameOfLife(t *testing.T) {
	tests := []struct {
		name string
		brd  [][]int
		want [][]int
	}{
		{
			name: "Example 1 - Multiple patterns",
			brd:  [][]int{{0, 1, 0}, {0, 0, 1}, {1, 1, 1}, {0, 0, 0}},
			want: [][]int{{0, 0, 0}, {1, 0, 1}, {0, 1, 1}, {0, 1, 0}},
		},
		{
			name: "Example 2 - Stable pattern",
			brd:  [][]int{{1, 1}, {1, 0}},
			want: [][]int{{1, 1}, {1, 1}},
		},
		{
			name: "Single live cell dies",
			brd:  [][]int{{1}},
			want: [][]int{{0}},
		},
		{
			name: "All dead cells remain dead",
			brd:  [][]int{{0, 0}, {0, 0}},
			want: [][]int{{0, 0}, {0, 0}},
		},
		{
			name: "Empty board edge case",
			brd:  [][]int{},
			want: [][]int{},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// Create deep copy of input board
			inp := make([][]int, len(tc.brd))
			for i := range tc.brd {
				inp[i] = make([]int, len(tc.brd[i]))
				copy(inp[i], tc.brd[i])
			}

			// Run the game simulation
			gameOfLife(inp)

			// Verify results
			for i := range tc.want {
				for j := range tc.want[i] {
					if inp[i][j] != tc.want[i][j] {
						t.Errorf("%s: board[%d][%d] = %d, want %d",
							tc.name, i, j, inp[i][j], tc.want[i][j])
					}
				}
			}
		})
	}
}
