package main

import (
	"container/heap"
	"reflect"
	"testing"
)

// Time complexity: O(k log k), We visit k elements due to sort-ascending. We push and pop each element into a min-heap. heap operations are each O(log k).
// Space complexity: O(k), k elements are stored in a heap.
// https://claude.ai/chat/611bbd15-fe3c-4b7c-9b1a-578b754692fe
func kSmallestPairs373(nums1 []int, nums2 []int, k int) [][]int {
	// Find K Pairs with Smallest Sums
	// Given two sort-ascending integer arrays nums1 and nums2.
	// Given an integer k.
	// Find k pairs which are min sums.
	// Use a struct Pair, idx1, idx2, sum.
	// Use a min-heap with sorting based on sum.
	// Use Go Heap interface.
	// Rely on sort-ascending for early pruning.

	// Check input min edge cases.
	if len(nums1) == 0 || len(nums2) == 0 || k <= 0 {
		return [][]int{}
	}

	// Initialize variables.
	maxPairCnt := min(k, len(nums1)*len(nums2))
	res := make([][]int, 0, maxPairCnt)
	h := &PairHeap373{}
	heap.Init(h)

	// Populate heap with nums1 and first column.
	for n := 0; n < min(k, len(nums1)); n++ {
		heap.Push(h, Pair373{n, 0, nums1[n] + nums2[0]})
	}

	// Process pairs.
	for h.Len() > 0 && len(res) < maxPairCnt {
		// Get current pair.
		cur := heap.Pop(h).(Pair373)

		// Record element pair in result.
		res = append(res, []int{nums1[cur.idx1], nums2[cur.idx2]})

		// Search next nums2.
		if cur.idx2+1 < len(nums2) {
			heap.Push(h, Pair373{
				cur.idx1, cur.idx2 + 1,
				nums1[cur.idx1] + nums2[cur.idx2+1],
			})
		}
	}

	return res
}

type Pair373 struct {
	idx1 int
	idx2 int
	sum  int
}

type PairHeap373 []Pair373

func (h PairHeap373) Len() int            { return len(h) }
func (h PairHeap373) Less(i, j int) bool  { return h[i].sum < h[j].sum }
func (h PairHeap373) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *PairHeap373) Push(x interface{}) { *h = append(*h, x.(Pair373)) }
func (h *PairHeap373) Pop() interface{} {
	prv := *h
	val := prv[len(prv)-1]
	*h = prv[:len(prv)-1]
	return val
}

func TestKSmallestPairs373(t *testing.T) {
	tests := []struct {
		name     string
		nums1    []int
		nums2    []int
		k        int
		expected [][]int
	}{
		{
			name:     "Example 1",
			nums1:    []int{1, 7, 11},
			nums2:    []int{2, 4, 6},
			k:        3,
			expected: [][]int{{1, 2}, {1, 4}, {1, 6}},
		},
		{
			name:     "Example 2",
			nums1:    []int{1, 1, 2},
			nums2:    []int{1, 2, 3},
			k:        2,
			expected: [][]int{{1, 1}, {1, 1}},
		},
		{
			name:     "Empty arrays",
			nums1:    []int{},
			nums2:    []int{1},
			k:        1,
			expected: [][]int{},
		},
		{
			name:     "K larger than possible pairs",
			nums1:    []int{1},
			nums2:    []int{2},
			k:        5,
			expected: [][]int{{1, 2}},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := kSmallestPairs373(tc.nums1, tc.nums2, tc.k)
			if !reflect.DeepEqual(got, tc.expected) {
				t.Errorf("got %v, want %v", got, tc.expected)
			}
		})
	}
}
