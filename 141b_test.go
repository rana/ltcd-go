package main

import "testing"

// ListNode141b represents a node in a linked list.
type ListNode141b struct {
	Val  int
	Next *ListNode141b
}

// Time complexity:
// Space complexity:
func hasCycle141b(hed *ListNode141b) bool {
	return false
}

// Test cases
func TestHasCycle141b(t *testing.T) {
	// Helper function to create a linked list with a cycle
	createCyclicList := func(vals []int, pos int) *ListNode141b {
		if len(vals) == 0 {
			return nil
		}

		// Create nodes
		nodes := make([]*ListNode141b, len(vals))
		for i := range vals {
			nodes[i] = &ListNode141b{Val: vals[i]}
		}

		// Link nodes
		for i := 0; i < len(nodes)-1; i++ {
			nodes[i].Next = nodes[i+1]
		}

		// Create cycle if pos is valid
		if pos >= 0 && pos < len(nodes) {
			nodes[len(nodes)-1].Next = nodes[pos]
		}

		return nodes[0]
	}

	tests := []struct {
		name     string
		vals     []int
		pos      int
		expected bool
	}{
		{"Example1", []int{3, 2, 0, -4}, 1, true},
		{"Example2", []int{1, 2}, 0, true},
		{"Example3", []int{1}, -1, false},
		{"EmptyList", []int{}, -1, false},
		{"TwoNodesNoCycle", []int{1, 2}, -1, false},
		{"LongListNoCycle", []int{1, 2, 3, 4, 5}, -1, false},
		{"CycleAtEnd", []int{1, 2, 3, 4}, 3, true},
		{"CycleToStart", []int{1, 2, 3, 4}, 0, true},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			hed := createCyclicList(tc.vals, tc.pos)
			res := hasCycle141b(hed)
			if res != tc.expected {
				t.Errorf("hasCycle() = %v, want %v", res, tc.expected)
			}
		})
	}
}
