package main

import (
	"testing"
)

func convert6b(s string, numRows int) string {
	return ""
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
