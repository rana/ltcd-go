package main

import "testing"

// Time complexity:
// Space complexity:
func isValid20b(s string) bool {
	return false
}

func TestIsValid20b(t *testing.T) {
	tests := []struct {
		name string
		str  string
		want bool
	}{
		{
			name: "single pair",
			str:  "()",
			want: true,
		},
		{
			name: "multiple pairs",
			str:  "()[]{}",
			want: true,
		},
		{
			name: "nested pairs",
			str:  "([{}])",
			want: true,
		},
		{
			name: "mismatched brackets",
			str:  "(]",
			want: false,
		},
		{
			name: "unmatched opening",
			str:  "([",
			want: false,
		},
		{
			name: "unmatched closing",
			str:  ")]",
			want: false,
		},
		{
			name: "wrong order",
			str:  "([)]",
			want: false,
		},
		{
			name: "single character",
			str:  "[",
			want: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := isValid20b(tc.str)
			if got != tc.want {
				t.Errorf("isValid(%q) = %v; want %v", tc.str, got, tc.want)
			}
		})
	}
}
