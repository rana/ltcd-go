package main

import "testing"

// WordDictionary211 implements a trie with wildcard search support
// Space complexity: O(n * m), n is the number of words added. m is the maximum word length.
type WordDictionary211 struct {
	rot *TrieNode211
}

type TrieNode211 struct {
	chd [26]*TrieNode211
	fin bool
}

// Constructor211 initializes the WordDictionary
func Constructor211() WordDictionary211 {
	return WordDictionary211{
		rot: &TrieNode211{},
	}
}

// AddWord adds a word to the dictionary
// Time complexity: O(m), m is the length of the word.
func (wdc *WordDictionary211) AddWord(wrd string) {
	cur := wdc.rot

	// Iterate through characters in word.
	for _, chr := range wrd {
		idx := chr - 'a'
		if cur.chd[idx] == nil {
			cur.chd[idx] = &TrieNode211{}
		}
		cur = cur.chd[idx]
	}

	cur.fin = true // set end of word
}

// Search checks if word exists in dictionary, supporting wildcards
// Time complexity: O(m), for no wildcard.
// Time complexity: O(26^k * m), k is the number of wild cards.
func (wdc *WordDictionary211) Search(wrd string) bool {
	return wdc.dfsFndWrd(wdc.rot, wrd, 0)
}

func (wdc *WordDictionary211) dfsFndWrd(cur *TrieNode211, wrd string, pos int) bool {
	// Check base case.
	if len(wrd) == pos {
		return cur.fin
	}

	// Get current character.
	chr := wrd[pos]

	// Handle wild card.
	if chr == '.' {
		// Iterate through character set.
		for idx := 0; idx < 26; idx++ {
			if cur.chd[idx] != nil &&
				wdc.dfsFndWrd(cur.chd[idx], wrd, pos+1) {
				return true
			}
		}
		return false
	}

	// Regular character.
	idx := chr - 'a'
	if cur.chd[idx] == nil {
		return false
	}

	// Recurse
	return wdc.dfsFndWrd(cur.chd[idx], wrd, pos+1)
}

// Unit tests
func TestWordDictionary211(t *testing.T) {
	tests := []struct {
		name     string
		actions  []string
		inputs   []string
		expected []interface{}
	}{
		{
			name:     "Example from problem",
			actions:  []string{"Constructor", "addWord", "addWord", "addWord", "search", "search", "search", "search"},
			inputs:   []string{"", "bad", "dad", "mad", "pad", "bad", ".ad", "b.."},
			expected: []interface{}{nil, nil, nil, nil, false, true, true, true},
		},
		{
			name:     "Empty word edge case",
			actions:  []string{"Constructor", "addWord", "search", "search"},
			inputs:   []string{"", "", "", "a"},
			expected: []interface{}{nil, nil, true, false},
		},
		{
			name:     "Multiple wildcards",
			actions:  []string{"Constructor", "addWord", "search", "search", "search"},
			inputs:   []string{"", "hello", "h.ll.", "..ll.", "....."},
			expected: []interface{}{nil, nil, true, true, true},
		},
		{
			name:     "No matches with wildcards",
			actions:  []string{"Constructor", "addWord", "search", "search"},
			inputs:   []string{"", "abc", ".d.", "..d"},
			expected: []interface{}{nil, nil, false, false},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			var wdc WordDictionary211

			for idx, action := range tc.actions {
				switch action {
				case "Constructor":
					wdc = Constructor211()
				case "addWord":
					wdc.AddWord(tc.inputs[idx])
				case "search":
					got := wdc.Search(tc.inputs[idx])
					want := tc.expected[idx].(bool)
					if got != want {
						t.Errorf("Search(%s) = %v, want %v", tc.inputs[idx], got, want)
					}
				}
			}
		})
	}
}
