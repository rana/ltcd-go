package main

import "testing"

// Time Complexity: O(W * H), W = width of field, H = height of field
// Space Complexity: O(1) https://console.anthropic.com/workbench/de5872d4-9e69-42be-a82e-060820c681b2
func findDropColumnGCA3(fld [][]int, fig [][]int) int {
	// Implmentation Efficiency
	// Find Drop Column
	// Figure drop in field matrix.
	// Given matrcies fld and fig.
	// Drop a figure at a top column.
	// Determine the top column which drops
	// to fill a complete lower row.
	// Return the column index; or, -1.
	// Use closures hits, fullRow
	// Iterate through field columns which meets condition.
	fldRows, fldCols := len(fld), len(fld[0])
	const figRows = 3
	const figCols = 3

	hits := func(fld, fig [][]int, fldRow, fldCol int) bool {
		for figRow := range figRows {
			for figCol := range figCols {
				if fig[figRow][figCol] == 1 && fld[fldRow+figRow][fldCol+figCol] == 1 {
					return true
				}
			}
		}
		return false
	}

	fullRow := func(fld, fig [][]int, figRow, fldRow, fldCol int) bool {
		// Also check field before and after figure
		for curCol := range fldCols {
			if curCol < fldCol || curCol >= fldCol+figCols {
				if fld[fldRow][curCol] == 0 {
					return false
				}
			}
		}
		// Check figure overlay positions
		for figCol := range figCols {
			if fig[figRow][figCol] == 0 && fld[fldRow][fldCol+figCol] == 0 {
				return false
			}
		}
		return true
	}

	for fldCol := 0; fldCol <= fldCols-figCols; fldCol++ {
		landRow := 0
		for landRow < fldRows-figRows && !hits(fld, fig, landRow+1, fldCol) {
			landRow++
		}
		for figRow := range figRows {
			if fullRow(fld, fig, figRow, figRow+landRow, fldCol) {
				return fldCol
			}
		}
	}

	return -1
}

func TestSolutionGCA3(t *testing.T) {
	tests := []struct {
		name   string
		field  [][]int
		figure [][]int
		want   int
	}{
		{
			name: "Example 1",
			field: [][]int{
				{0, 0, 0},
				{0, 0, 0},
				{0, 0, 0},
				{1, 0, 0},
				{1, 1, 0},
			},
			figure: [][]int{
				{0, 0, 1},
				{0, 1, 1},
				{0, 0, 1},
			},
			want: 0,
		},
		{
			name: "Example 2",
			field: [][]int{
				{0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0},
				{1, 1, 0, 1, 0},
				{1, 0, 1, 0, 1},
			},
			figure: [][]int{
				{1, 1, 1},
				{1, 0, 1},
				{1, 0, 1},
			},
			want: 2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := findDropColumnGCA3(tt.field, tt.figure)
			if got != tt.want {
				t.Errorf("solution() = %v, want %v", got, tt.want)
			}
		})
	}
}
