package main

import "testing"

// Time complexity:
// Space complexity:
func isIsomorphic205b(s string, t string) bool {
	return false
}

func TestIsIsomorphic205b(t *testing.T) {
	tests := []struct {
		name string
		str  string
		txt  string
		want bool
	}{
		{
			name: "Example 1: Valid mapping",
			str:  "egg",
			txt:  "add",
			want: true,
		},
		{
			name: "Example 2: Invalid mapping",
			str:  "foo",
			txt:  "bar",
			want: false,
		},
		{
			name: "Example 3: Paper-Title",
			str:  "paper",
			txt:  "title",
			want: true,
		},
		{
			name: "Empty strings",
			str:  "",
			txt:  "",
			want: true,
		},
		{
			name: "Single character",
			str:  "a",
			txt:  "b",
			want: true,
		},
		{
			name: "Same characters different pattern",
			str:  "badc",
			txt:  "baba",
			want: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if got := isIsomorphic205b(tc.str, tc.txt); got != tc.want {
				t.Errorf("isIsomorphic(%q, %q) = %v, want %v",
					tc.str, tc.txt, got, tc.want)
			}
		})
	}
}
