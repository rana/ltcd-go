package main

import "testing"

// Time Complexity:
// Space Complexity:
func longestPalindrome2131b(wrds []string) int {
	return 0
}

func TestLongestPalindrome2131b(t *testing.T) {
	tests := []struct {
		name string
		wrds []string
		want int
	}{
		{
			name: "Example 1",
			wrds: []string{"lc", "cl", "gg"},
			want: 6,
		},
		{
			name: "Example 2",
			wrds: []string{"ab", "ty", "yt", "lc", "cl", "ab"},
			want: 8,
		},
		{
			name: "Example 3",
			wrds: []string{"cc", "ll", "xx"},
			want: 2,
		},
		{
			name: "Empty input",
			wrds: []string{},
			want: 0,
		},
		{
			name: "Single word",
			wrds: []string{"ab"},
			want: 0,
		},
		{
			name: "All same words",
			wrds: []string{"aa", "aa", "aa"},
			want: 6,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestPalindrome2131b(tt.wrds); got != tt.want {
				t.Errorf("longestPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
