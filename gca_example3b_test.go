package main

import "testing"

// Time Complexity:
// Space Complexity:
func findDropColumnGCA3b(fld, fig [][]int) int {
	// Find Drop Column
	// Given matricies fld, fig
	// Determine drop column of fig in fld which completes a row.
	// Return the drop column.
	// Conditions:
	// * Matrix elements are 0 or 1
	// * Complete row is all 1s
	// Use closures hits, fullRow
	fldRows, fldCols := len(fld), len(fld[0])
	figRows, figCols := 3, 3

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
		// Check fld outside of fig.
		for curFldCol := range fldCols {
			if curFldCol < fldCol || curFldCol >= fldCol+figCols {
				if fld[fldRow][curFldCol] == 0 {
					return false
				}
			}
		}
		// Check inside fig and fld translated
		for figCol := range figCols {
			if fig[figRow][figCol] == 0 && fld[fldRow][fldCol+figCol] == 0 {
				return false
			}
		}
		return true
	}

	// Scan fld cols for drop
	for fldCol := 0; fldCol <= fldCols-figCols; fldCol++ {
		// Search landing row
		landRow := 0
		for landRow <= fldRows-figRows && !hits(fld, fig, landRow+1, fldCol) {
			landRow++
		}
		// Search fig complete row
		for figRow := range figRows {
			if fullRow(fld, fig, figRow, landRow+figRow, fldCol) {
				return fldCol
			}
		}
	}
	return -1
}

func TestSolutionGCA3b(t *testing.T) {
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
			got := findDropColumnGCA3b(tt.field, tt.figure)
			if got != tt.want {
				t.Errorf("solution() = %v, want %v", got, tt.want)
			}
		})
	}
}
