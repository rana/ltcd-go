package main

import "testing"

// Implement trie/prefix tree.
// Trie208 represents a prefix tree data structure.
// Time complexity: O(m), m is the length of the string.
// Space complexity: O(n * m), n is the number of words.
type Trie208 struct {
	rot *TrieNode208
}

type TrieNode208 struct {
	chd map[byte]*TrieNode208
	isw bool
}

// Constructor208 initializes a new Trie
func Constructor208() Trie208 {
	return Trie208{
		rot: &TrieNode208{
			chd: make(map[byte]*TrieNode208),
			isw: false,
		},
	}
}

// Insert adds a word to the trie
func (tri *Trie208) Insert(wrd string) {
	cur := tri.rot

	// Traverse/create each character in word.
	for idx := range len(wrd) {
		chr := wrd[idx]
		if cur.chd[chr] == nil {
			cur.chd[chr] = &TrieNode208{
				chd: make(map[byte]*TrieNode208),
				isw: false,
			}
		}
		cur = cur.chd[chr]
	}
	cur.isw = true
}

// Search checks if the word exists in the trie
func (tri *Trie208) Search(wrd string) bool {
	cur := tri.rot

	for idx := range len(wrd) {
		chr := wrd[idx]
		if cur.chd[chr] == nil {
			return false
		}
		cur = cur.chd[chr]
	}

	// Check if character is end of word.
	return cur.isw
}

// StartsWith checks if any word in trie starts with given prefix
func (tri *Trie208) StartsWith(pfx string) bool {
	cur := tri.rot

	for idx := range len(pfx) {
		chr := pfx[idx]
		if cur.chd[chr] == nil {
			return false
		}
		cur = cur.chd[chr]
	}

	return true
}

// Tests
func TestTrie208(t *testing.T) {
	tests := []struct {
		name     string
		ops      []string
		args     []string
		expected []interface{}
	}{
		{
			name:     "Example from description",
			ops:      []string{"Trie", "insert", "search", "search", "startsWith", "insert", "search"},
			args:     []string{"", "apple", "apple", "app", "app", "app", "app"},
			expected: []interface{}{nil, nil, true, false, true, nil, true},
		},
		{
			name:     "Empty string operations",
			ops:      []string{"Trie", "insert", "search", "startsWith"},
			args:     []string{"", "", "", ""},
			expected: []interface{}{nil, nil, true, true},
		},
		{
			name:     "Single character operations",
			ops:      []string{"Trie", "insert", "search", "startsWith", "search"},
			args:     []string{"", "a", "a", "a", "b"},
			expected: []interface{}{nil, nil, true, true, false},
		},
		{
			name:     "Prefix operations",
			ops:      []string{"Trie", "insert", "startsWith", "startsWith", "startsWith"},
			args:     []string{"", "hello", "hel", "hello", "world"},
			expected: []interface{}{nil, nil, true, true, false},
		},
		{
			name:     "Multiple words with common prefix",
			ops:      []string{"Trie", "insert", "insert", "search", "search", "startsWith"},
			args:     []string{"", "cat", "cats", "cat", "cats", "ca"},
			expected: []interface{}{nil, nil, nil, true, true, true},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			var tri Trie208

			for idx := range tc.ops {
				switch tc.ops[idx] {
				case "Trie":
					tri = Constructor208()
				case "insert":
					tri.Insert(tc.args[idx])
				case "search":
					got := tri.Search(tc.args[idx])
					want := tc.expected[idx].(bool)
					if got != want {
						t.Errorf("Search(%s) = %v, want %v", tc.args[idx], got, want)
					}
				case "startsWith":
					got := tri.StartsWith(tc.args[idx])
					want := tc.expected[idx].(bool)
					if got != want {
						t.Errorf("StartsWith(%s) = %v, want %v", tc.args[idx], got, want)
					}
				}
			}
		})
	}
}
