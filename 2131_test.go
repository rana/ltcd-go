package main

import "testing"

// Time Complexity: O(n), Building frequency map
// Space Complexity: O(n),
func longestPalindrome2131(wrds []string) int {
	// Longest Paindrome by Concatenating Two Letter Words
	// Determine the max length creatable palindrome.
	// Return the max length.
	// Conditions:
	// * Words are two letters
	// * Select each word once
	// * Word may have duplicate char "gg"
	// * May have duplicate words
	// Use a frequency counter map.
	frq := make(map[string]int)
	for _, wrd := range wrds {
		frq[wrd]++
	}
	chrCnt := 0
	centrCnt := 0
	for wrd, wrdCnt := range frq {
		rev := string([]byte{wrd[1], wrd[0]})
		if wrd == rev {
			chrCnt += (wrdCnt / 2) * 4
			centrCnt += wrdCnt % 2
		} else if revCnt := frq[rev]; revCnt > 0 {
			chrCnt += min(wrdCnt, revCnt) * 2
		}
	}
	if centrCnt > 0 {
		chrCnt += 2
	}
	return chrCnt
}

func TestLongestPalindrome2131(t *testing.T) {
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
			if got := longestPalindrome2131(tt.wrds); got != tt.want {
				t.Errorf("longestPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
