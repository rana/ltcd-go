package main

import "testing"

// Time Complexity Analysis:
// - Let k1 and k2 be the number of 1's in img1 and img2 respectively
// - Finding all 1's positions: O(n²) where n is the matrix dimension
// - Computing all possible translations: O(k1 * k2)
// - Overall complexity: O(n² + k1*k2)
// - In worst case where all elements are 1's: O(n⁴)
//
// Space Complexity:
// - O(k1 + k2) for storing positions of 1's
// - O(k1 * k2) for storing translation frequencies
// - In worst case: O(n²)
//
// https://console.anthropic.com/workbench/cd318ef3-2f71-42a6-9bba-aab3c3d052b3
func largestOverlap835(img1 [][]int, img2 [][]int) int {
	// Image Overlap
	// Given two n x n matrixes img1, img2.
	// Translate aribtrary x-y.
	// Determine max overlap between img1 img2.
	// Return max overlap.
	// Conditions:
	// * Matrix element is 0 or 1
	// * Translate 1-bits -x, +x, -y, +y
	// * Translation does not rotate values.
	// * Translation edges are replaced with zeros.
	// Track 1-bit positions in each image.
	// Ensure 1s exist in both image.
	// Translate positions.
	// Frequency count positions.
	// Track max freq count.

	n := len(img1)
	var pos1 [][2]int
	var pos2 [][2]int
	for row := range n {
		for col := range n {
			if img1[row][col] == 1 {
				pos1 = append(pos1, [2]int{row, col})
			}
			if img2[row][col] == 1 {
				pos2 = append(pos2, [2]int{row, col})
			}
		}
	}

	if len(pos1) == 0 || len(pos2) == 0 {
		return 0
	}

	maxTrnCnt := 0
	trns := make(map[[2]int]int)
	for _, p1 := range pos1 {
		for _, p2 := range pos2 {
			trn := [2]int{p2[0] - p1[0], p2[1] - p1[1]}
			trns[trn]++
			if trns[trn] > maxTrnCnt {
				maxTrnCnt = trns[trn]
			}
		}
	}
	return maxTrnCnt
}

func TestLargestOverlap835(t *testing.T) {
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
			got := largestOverlap835(tc.img1, tc.img2)
			if got != tc.expected {
				t.Errorf("got %v, want %v", got, tc.expected)
			}
		})
	}
}
