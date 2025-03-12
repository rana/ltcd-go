package main

import (
	"testing"
)

// Time complexity:
// Space complexity:
func groupAnagrams49b(strs []string) [][]string {
	return nil
}

func TestGroupAnagrams49b(t *testing.T) {
	tests := []struct {
		name string
		strs []string
		want int // Number of groups expected
	}{
		{
			name: "Example 1",
			strs: []string{"eat", "tea", "tan", "ate", "nat", "bat"},
			want: 3,
		},
		{
			name: "Example 2",
			strs: []string{""},
			want: 1,
		},
		{
			name: "Example 3",
			strs: []string{"a"},
			want: 1,
		},
		{
			name: "All Different",
			strs: []string{"cat", "dog", "pig"},
			want: 3,
		},
		{
			name: "All Same",
			strs: []string{"abc", "bac", "cab", "acb"},
			want: 1,
		},
	}

	// Test both implementations
	implementations := []struct {
		name string
		fn   func([]string) [][]string
	}{
		{"Count Method", groupAnagrams49b},
	}

	for _, impl := range implementations {
		for _, tt := range tests {
			t.Run(impl.name+": "+tt.name, func(t *testing.T) {
				got := impl.fn(tt.strs)
				if len(got) != tt.want {
					t.Errorf("got %v groups, want %v groups", len(got), tt.want)
				}

				// Verify each group contains anagrams
				for _, grp := range got {
					if !areAnagrams49b(grp) {
						t.Errorf("group %v contains non-anagrams", grp)
					}
				}
			})
		}
	}
}

// Helper function to verify if strings are anagrams
func areAnagrams49b(strs []string) bool {
	if len(strs) <= 1 {
		return true
	}

	// Compare character counts of first string with others
	var cnt [26]int
	for _, chr := range strs[0] {
		cnt[chr-'a']++
	}

	for _, str := range strs[1:] {
		var curCnt [26]int
		for _, chr := range str {
			curCnt[chr-'a']++
		}
		if curCnt != cnt {
			return false
		}
	}
	return true
}
