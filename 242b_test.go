package main

import "testing"

// Time complexity:
// Space complexity:
func isAnagram242b(s, t string) bool {
	return false
}

func TestIsAnagram242b(t *testing.T) {
	tests := []struct {
		name string
		s    string
		t    string
		want bool
	}{
		{
			name: "example 1 - simple anagram",
			s:    "anagram",
			t:    "nagaram",
			want: true,
		},
		{
			name: "example 2 - not anagram",
			s:    "rat",
			t:    "car",
			want: false,
		},
		{
			name: "same strings",
			s:    "hello",
			t:    "hello",
			want: true,
		},
		{
			name: "empty strings",
			s:    "",
			t:    "",
			want: true,
		},
		{
			name: "different lengths",
			s:    "ab",
			t:    "abc",
			want: false,
		},
		{
			name: "same letters different counts",
			s:    "aaab",
			t:    "aaba",
			want: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := isAnagram242b(tc.s, tc.t)
			if got != tc.want {
				t.Errorf("isAnagram(%q, %q) = %v, want %v",
					tc.s, tc.t, got, tc.want)
			}
		})
	}
}
