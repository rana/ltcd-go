package main

import "testing"

// Node138 represents a node in the linked list
type Node138 struct {
	Val    int
	Next   *Node138
	Random *Node138
}

// Time complexity: O(n), n is the number of nodes in the list. We traverse the list three times.
// Space complexity: O(1), constant additional space used.
// https://claude.ai/chat/12e691a1-bfcc-4c78-9fef-002ea1437c32
func copyRandomList(hed *Node138) *Node138 {
	// Copy List with Random Pointer
	// Given the head of a linked list.
	// Linked list node has a pointer to a random node.
	// Create a deep copy of the linked list.
	// Return the head of the copied list.
	// Use a three-step approach:
	//	Step 1. Interleave copies with original list.
	//	Step 2.	Update random pointers for copied nodes.
	//	Step 3. De-interleave and separate lists.

	// Check input min edge case.
	if hed == nil {
		return nil
	}

	// Step 1. Interleave copies with original list.
	// Original: A -> B -> C
	// After: A -> A' -> B -> B' -> C -> C'
	cur := hed
	for cur != nil {
		nxt := cur.Next
		cpy := &Node138{Val: cur.Val}
		cur.Next = cpy
		cpy.Next = nxt
		cur = nxt
	}

	// Step 2. Update pointers to random in copies.
	cur = hed
	for cur != nil {
		cpy := cur.Next
		if cur.Random != nil {
			cpy.Random = cur.Random.Next
		}
		cur = cpy.Next
	}

	// Step 3. De-interleave: separate lists.
	cur = hed
	cpyHed := hed.Next
	for cur != nil {
		cpy := cur.Next
		nxt := cpy.Next

		// Update original list.
		cur.Next = nxt

		// Update coped list.
		if nxt != nil {
			cpy.Next = nxt.Next
		}

		cur = nxt
	}

	return cpyHed
}

// For testing purposes
func createList138(vals [][]int) *Node138 {
	if len(vals) == 0 {
		return nil
	}

	// Create nodes
	nods := make([]*Node138, len(vals))
	for idx := range vals {
		nods[idx] = &Node138{Val: vals[idx][0]}
	}

	// Set next pointers
	for idx := 0; idx < len(nods)-1; idx++ {
		nods[idx].Next = nods[idx+1]
	}

	// Set random pointers
	for idx, val := range vals {
		if val[1] >= 0 {
			nods[idx].Random = nods[val[1]]
		}
	}

	return nods[0]
}

func compareNodes138(org, cpy *Node138) bool {
	// Map to store node mapping from original to copy
	nodMap := make(map[*Node138]*Node138)

	// First pass: build node mapping
	cur1, cur2 := org, cpy
	for cur1 != nil {
		if cur1.Val != cur2.Val {
			return false
		}
		nodMap[cur1] = cur2
		cur1 = cur1.Next
		cur2 = cur2.Next
	}

	// Second pass: verify random pointers
	cur1, cur2 = org, cpy
	for cur1 != nil {
		if (cur1.Random == nil && cur2.Random != nil) ||
			(cur1.Random != nil && cur2.Random != nodMap[cur1.Random]) {
			return false
		}
		cur1 = cur1.Next
		cur2 = cur2.Next
	}

	return true
}

func TestCopyRandomList(t *testing.T) {
	tests := []struct {
		name string
		vals [][]int
	}{
		{
			name: "Example 1",
			vals: [][]int{{7, -1}, {13, 0}, {11, 4}, {10, 2}, {1, 0}},
		},
		{
			name: "Example 2",
			vals: [][]int{{1, 1}, {2, 1}},
		},
		{
			name: "Example 3",
			vals: [][]int{{3, -1}, {3, 0}, {3, -1}},
		},
		{
			name: "Empty List",
			vals: [][]int{},
		},
		{
			name: "Single Node with Self Random",
			vals: [][]int{{1, 0}},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			org := createList138(tc.vals)
			cpy := copyRandomList(org)

			// Check if lists are identical but separate
			if !compareNodes138(org, cpy) {
				t.Errorf("Lists are not identical")
			}

			// Verify copied list is separate from original
			cur1, cur2 := org, cpy
			for cur1 != nil {
				if cur1 == cur2 {
					t.Errorf("Copied list contains nodes from original list")
				}
				cur1 = cur1.Next
				cur2 = cur2.Next
			}
		})
	}
}
