package main

import "testing"

// ListNode141 represents a node in a linked list.
type ListNode141 struct {
	Val  int
	Next *ListNode141
}

// Time complexity: O(n), n is the number of nodes in the linked list. Slow and fast pointers move through up to n nodes if there is a cycle. Possibly less if there is no cycle.
// Space complexity: O(1), constant additional space used.
// https://claude.ai/chat/66252c28-5c22-4051-83ca-5068017ef405
func hasCycle(hed *ListNode141) bool {
	// Linked List Cycle
	// Given the head of a linked list.
	// Determine whether the linked list has a cycle.
	// Return false if the linked list does not have a cycle.
	// Use Floyd's cycle-finding algorithm.

	// Check input min edge cases.
	if hed == nil || hed.Next == nil {
		return false
	}

	// Initialize slow and fast pointers.
	slw := hed
	fst := hed.Next

	// Move pointers forward until end; or, they meet.
	for fst != nil && fst.Next != nil && slw != fst {
		slw = slw.Next
		fst = fst.Next.Next
	}

	// Check success conditions.
	if fst == nil || fst.Next == nil {
		return false
	}

	// Cycle detected.
	return true
}

// Test cases
func TestHasCycle(t *testing.T) {
	// Helper function to create a linked list with a cycle
	createCyclicList := func(vals []int, pos int) *ListNode141 {
		if len(vals) == 0 {
			return nil
		}

		// Create nodes
		nodes := make([]*ListNode141, len(vals))
		for i := range vals {
			nodes[i] = &ListNode141{Val: vals[i]}
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
			res := hasCycle(hed)
			if res != tc.expected {
				t.Errorf("hasCycle() = %v, want %v", res, tc.expected)
			}
		})
	}
}
