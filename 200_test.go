package main

import "testing"

// Time complexity: O(m * n), m and n are dimensions of the grid.
// Space complexity: O(m * n), for recursive stack. Worst case, all cells are land and are visited.
// https://claude.ai/chat/c161780e-32e6-4bea-876c-3dcb17ac296d
func numIslands200(grid [][]byte) int {
	// Number of Islands
	// Given an MxN matrix grid.
	// Each cell is water '0', or land '1'.
	// Count the number of islands.
	// Return the number of islands.
	// Use depth-first search and recursion.
	// Toggle visited land cells to water.
	// Use a helper function to visit a single cell recursively.

	// Check input min edge cases.
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}

	// Initialize variables.
	rows, cols := len(grid), len(grid[0])
	cnt := 0 // island count

	// Iterate through each cell in the grid.
	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			// Explore unvisited land cell.
			if grid[row][col] == '1' {
				exploreIsland200(grid, row, col)
				cnt++
			}
		}
	}

	return cnt
}

func exploreIsland200(grid [][]byte, row, col int) {
	// Check bounds and land.
	if row < 0 || row >= len(grid) ||
		col < 0 || col >= len(grid[0]) ||
		grid[row][col] != '1' {
		return
	}

	// Toggle current cell from land to water to mark visiting.
	grid[row][col] = '0'

	// Explore adjacent cells horizontally, vertically.
	exploreIsland200(grid, row-1, col) // up
	exploreIsland200(grid, row+1, col) // down
	exploreIsland200(grid, row, col-1) // left
	exploreIsland200(grid, row, col+1) // right
}

func TestNumIslands200(t *testing.T) {
	cases := []struct {
		name string
		grid [][]byte
		want int
	}{
		{
			name: "single island",
			grid: [][]byte{
				{'1', '1', '1', '1', '0'},
				{'1', '1', '0', '1', '0'},
				{'1', '1', '0', '0', '0'},
				{'0', '0', '0', '0', '0'},
			},
			want: 1,
		},
		{
			name: "multiple islands",
			grid: [][]byte{
				{'1', '1', '0', '0', '0'},
				{'1', '1', '0', '0', '0'},
				{'0', '0', '1', '0', '0'},
				{'0', '0', '0', '1', '1'},
			},
			want: 3,
		},
		{
			name: "empty grid",
			grid: [][]byte{},
			want: 0,
		},
		{
			name: "all water",
			grid: [][]byte{
				{'0', '0', '0'},
				{'0', '0', '0'},
			},
			want: 0,
		},
		{
			name: "all land",
			grid: [][]byte{
				{'1', '1', '1'},
				{'1', '1', '1'},
			},
			want: 1,
		},
		{
			name: "single cell",
			grid: [][]byte{{'1'}},
			want: 1,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			// Create deep copy of input grid since we modify it
			gridCopy := make([][]byte, len(tc.grid))
			for i := range tc.grid {
				gridCopy[i] = make([]byte, len(tc.grid[i]))
				copy(gridCopy[i], tc.grid[i])
			}

			got := numIslands200(gridCopy)
			if got != tc.want {
				t.Errorf("got %d, want %d", got, tc.want)
			}
		})
	}
}
