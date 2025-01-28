package main

import "testing"

// Time Complexity:
// Space Complexity:
func largestOverlap835b(img1 [][]int, img2 [][]int) int {
	return 0
}

func TestLargestOverlap835b(t *testing.T) {
	tests := []struct {
		name     string
		img1     [][]int
		img2     [][]int
		expected int
	}{
		{
			name:     "Example 1",
			img1:     [][]int{{1, 1, 0}, {0, 1, 0}, {0, 1, 0}},
			img2:     [][]int{{0, 0, 0}, {0, 1, 1}, {0, 0, 1}},
			expected: 3,
		},
		{
			name:     "Example 2",
			img1:     [][]int{{1}},
			img2:     [][]int{{1}},
			expected: 1,
		},
		{
			name:     "Example 3",
			img1:     [][]int{{0}},
			img2:     [][]int{{0}},
			expected: 0,
		},
		{
			name:     "No overlap possible",
			img1:     [][]int{{1, 0}, {0, 0}},
			img2:     [][]int{{0, 0}, {0, 1}},
			expected: 1,
		},
		{
			name:     "Full overlap possible",
			img1:     [][]int{{1, 1}, {1, 1}},
			img2:     [][]int{{1, 1}, {1, 1}},
			expected: 4,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := largestOverlap835b(tc.img1, tc.img2)
			if got != tc.expected {
				t.Errorf("got %v, want %v", got, tc.expected)
			}
		})
	}
}
