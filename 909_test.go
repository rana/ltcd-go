package main

import "testing"

// Time complexity: O(n^2), n is the length of the board.
// Space complexity: O(n^2), we store distance moved for each cell.
// https://claude.ai/chat/d4f1d733-1a14-41ae-9637-aba56fc4a3e0
func snakesAndLadders909(brd [][]int) int {
	// Snakes and Ladders
	// Given an n x n matrix `brd`.
	// Cells labeled 1 to n^2.
	// 1 is lower left, move left to right.
	// Second row up is right to left.
	// Third ro lft to right.
	// Conditions:
	//	* Start square 1
	//	* Roll

	// Initialize dimensions.
	brdLen := len(brd)
	brdSqr := brdLen * brdLen

	// Create closure for coordinate transformation.
	getBrdPos := func(pos int) (int, int) {
		// Adjust to 0-based indexing.
		pos--
		// Calculate row from bottom.
		row := brdLen - 1 - pos/brdLen
		// Calculate column based on row.
		col := pos % brdLen
		if (brdLen-1-row)%2 == 1 {
			col = brdLen - 1 - col
		}
		return row, col
	}

	// Create distance array to track minimum moves.
	dst := make([]int, brdSqr+1)
	for idx := range dst {
		dst[idx] = -1
	}
	dst[1] = 0

	// BFS with initial position
	que := []int{1}
	for len(que) > 0 {
		// Get and remove first position from que.
		pos := que[0]
		que = que[1:]

		// Try all possible dice rolls (1-6)
		for nxt := pos + 1; nxt <= min(pos+6, brdSqr); nxt++ {
			row, col := getBrdPos(nxt)

			// Final position after nake/ladder teleport.
			fnl := nxt
			if brd[row][col] != -1 {
				fnl = brd[row][col]
			}

			// Skip is already found.
			if dst[fnl] != -1 {
				continue
			}

			// Record distance and add to que.
			dst[fnl] = dst[pos] + 1
			que = append(que, fnl)
		}
	}

	return dst[brdSqr]
}

// // Helper function for min
// func min(a, b int) int {
// 	if a < b {
// 		return a
// 	}
// 	return b
// }

// TestSnakesAndLadders909 implements comprehensive tests
func TestSnakesAndLadders909(t *testing.T) {
	tests := []struct {
		name string
		brd  [][]int
		want int
	}{
		{
			name: "Example 1 - Complex board with snakes and ladders",
			brd: [][]int{
				{-1, -1, -1, -1, -1, -1},
				{-1, -1, -1, -1, -1, -1},
				{-1, -1, -1, -1, -1, -1},
				{-1, 35, -1, -1, 13, -1},
				{-1, -1, -1, -1, -1, -1},
				{-1, 15, -1, -1, -1, -1},
			},
			want: 4,
		},
		{
			name: "Example 2 - Small board with one ladder",
			brd:  [][]int{{-1, -1}, {-1, 3}},
			want: 1,
		},
		{
			name: "Minimum board size",
			brd:  [][]int{{-1, -1}, {-1, -1}},
			want: 1,
		},
		{
			name: "Direct path possible",
			brd: [][]int{
				{-1, -1, -1, -1},
				{-1, -1, -1, -1},
				{-1, -1, -1, -1},
				{-1, -1, -1, -1},
			},
			want: 3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := snakesAndLadders909(tt.brd); got != tt.want {
				t.Errorf("snakesAndLadders() = %v, want %v", got, tt.want)
			}
		})
	}
}
