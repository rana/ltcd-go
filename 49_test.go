package main

import (
	"testing"
)

// Time complexity: O(n * k), n is the length of the strs array. k is the max length of a string.
// Space complexity: O(n * k), each word is stored in a map.
// https://claude.ai/chat/14c52221-8e91-4ef1-969b-15950dfcd5f4
func groupAnagrams49(strs []string) [][]string {
	// Group Anagrams
	// Given and array of string strs.
	// Each string is an anagram.
	// Group the anagrams.
	// Return an array of the groups.
	// Condition:
	//	* Lowercase English letters.
	// Use a frequency counter array as a key to a map group.

	// Initialize a map for grouping.
	// Use the frequency counter array as a key.
	grpMap := make(map[[26]int][]string)

	// Iterate through each string.
	for _, str := range strs {
		// Create a frequency counter array for the current string.
		var frq [26]int
		for _, chr := range str {
			frq[chr-'a']++
		}

		// Add the original string to the group using the counter.
		grpMap[frq] = append(grpMap[frq], str)
	}

	// Transform the groups to an aray for return.
	var res [][]string
	for _, grp := range grpMap {
		res = append(res, grp)
	}

	return res
}

// Test cases
func TestGroupAnagrams49(t *testing.T) {
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
		{"Count Method", groupAnagrams49},
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
					if !areAnagrams49(grp) {
						t.Errorf("group %v contains non-anagrams", grp)
					}
				}
			})
		}
	}
}

// Helper function to verify if strings are anagrams
func areAnagrams49(strs []string) bool {
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
