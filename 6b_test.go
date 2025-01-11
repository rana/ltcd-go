package main

import (
	"strings"
	"testing"
)

func convert6b(s string, numRows int) string {
	// Zigzag Conversion
	// Given a string s.
	// Given an integer numRows.
	// Transform string to zigzag pattern.
	// Return sigzag string.
	// Conditions:
	// * Zigzag descends until last row, then ascends til first row.
	// * Repeat descend ascend until characters complete.
	// Use an array of string builders.

	if numRows == 1 || len(s) <= numRows {
		return s
	}

	rows := make([]strings.Builder, numRows)
	curRow, dir := 0, 1
	for _, chr := range s {
		rows[curRow].WriteRune(chr)
		if curRow == 0 {
			dir = 1
		} else if curRow == numRows-1 {
			dir = -1
		}
		curRow += dir
	}

	var res strings.Builder
	for _, row := range rows {
		res.WriteString(row.String())
	}
	return res.String()
}

func TestConvert6b(t *testing.T) {
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
			if got := convert6b(tt.str, tt.numRows); got != tt.want {
				t.Errorf("convert() = %v, want %v", got, tt.want)
			}
		})
	}
}
