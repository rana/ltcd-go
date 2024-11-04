package main

import "testing"

// Time complexity: O(n * k), n is the length of the strs array. k is the maximum string length. We traverse the array once.
// Space complexity: O(n * k), we store each word in a map.
func groupAnagrams(strs []string) [][]string {
	// Group Anagrams
	// Given an array of strings strs.
	// The strings are anagrams.
	// Group string by anagrams.
	// Return an array of groups.
	// Condition:
	//	* strs is lowercase English letters.
	// Use a frquency counter array as a map key to groups.

	grpMap := make(map[[26]int][]string)

	// Iterate through each string.
	for _, str := range strs {
		// Create a character frequency counter for the current string.
		var frq [26]int
		for _, chr := range str {
			frq[chr-'a']++
		}

		// Add the current string to the group.
		grpMap[frq] = append(grpMap[frq], str)
	}

	// Transform groups to an array for return.
	var res [][]string
	for _, grp := range grpMap {
		res = append(res, grp)
	}

	return res
}

func TestGroupAnagrams(t *testing.T) {
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
		{"Count Method", groupAnagrams},
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
					if !areAnagrams(grp) {
						t.Errorf("group %v contains non-anagrams", grp)
					}
				}
			})
		}
	}
}

// Helper function to verify if strings are anagrams
func areAnagrams(strs []string) bool {
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
