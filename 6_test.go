package main

import (
	"strings"
	"testing"
)

// Time complexity: O(n), n is the length of string `str`. We traverse the string twice. Once to build the zigzag pattern. And a second time to concatenate the return string.
// Space complexity: O(n), n characters are stored in rows for the zigzag transform operation.
// https://claude.ai/chat/e67a775f-d714-463a-9420-fd20d31a9f60
func convert6(str string, numRows int) string {
	// Zigzag Converstion
	// Given a string str with English letters.
	// Given an integer numRows.
	// Transform the string into a zigzag pattern.
	// Return the transformed string.
	// Zigzag transform conditions:
	//	* Descending vertical column.
	//	* Ascend by incrementing one column, one row.
	//	* Cycle until complete.
	// Technique is pattern recognition and string manipulation.

	// Check input minimum edge cases.
	if numRows == 1 || numRows >= len(str) {
		return str
	}

	// Initialize variables.
	rows := make([]strings.Builder, numRows)
	curRow := 0 // Current row index
	dir := 1    // 1 for descending rows. -1 for ascending rows.

	// Iterate through each character in str.
	for _, chr := range str {
		// Store the current character in the current row.
		rows[curRow].WriteRune(chr)

		// Update the movement direction as needed.
		if curRow == 0 {
			// Descending direction.
			dir = 1
		} else if curRow == numRows-1 {
			// Ascending direction.
			dir = -1
		}

		// Move to the next row.
		curRow += dir
	}

	// Concatenate the result string.
	var res strings.Builder
	for _, row := range rows {
		res.WriteString(row.String())
	}

	return res.String()
}

func TestConvert6(t *testing.T) {
	tests := []struct {
		name    string
		str     string
		numRows int
		want    string
	}{
		{
			name:    "Example 1",
			str:     "PAYPALISHIRING",
			numRows: 3,
			want:    "PAHNAPLSIIGYIR",
		},
		{
			name:    "Example 2",
			str:     "PAYPALISHIRING",
			numRows: 4,
			want:    "PINALSIGYAHRPI",
		},
		{
			name:    "Single Row",
			str:     "A",
			numRows: 1,
			want:    "A",
		},
		{
			name:    "Two Rows",
			str:     "AB",
			numRows: 2,
			want:    "AB",
		},
		{
			name:    "More Rows Than Characters",
			str:     "ABC",
			numRows: 4,
			want:    "ABC",
		},
		{
			name:    "Empty String",
			str:     "",
			numRows: 3,
			want:    "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := convert6(tt.str, tt.numRows); got != tt.want {
				t.Errorf("convert() = %v, want %v", got, tt.want)
			}
		})
	}
}
