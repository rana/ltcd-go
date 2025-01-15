package main

import (
	"strings"
	"testing"
)

func convert6b(s string, numRows int) string {
	// Zigzag Conversion
	// Given a string s.
	// Given an integer numRows.
	// Transform the string into a zigzag pattern.
	// Return the transformed string.
	// Pattern recognition.
	// Use an array of string builders.

	if len(s) <= numRows {
		return s
	}

	rows := make([]strings.Builder, numRows)
	dir, curRow := 1, 0
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
