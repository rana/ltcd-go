package main

import "testing"

// Node138b represents a node in the linked list
type Node138b struct {
	Val    int
	Next   *Node138b
	Random *Node138b
}

// Time complexity:
// Space complexity:
func copyRandomList138b(hed *Node138b) *Node138b {
	return nil
}

func TestCopyRandomList138b(t *testing.T) {
	createList := func(vals [][]int) *Node138b {
		if len(vals) == 0 {
			return nil
		}
		nods := make([]*Node138b, len(vals))
		for idx := range vals {
			nods[idx] = &Node138b{Val: vals[idx][0]}
		}
		for idx := 0; idx < len(nods)-1; idx++ {
			nods[idx].Next = nods[idx+1]
		}
		for idx, val := range vals {
			if val[1] >= 0 {
				nods[idx].Random = nods[val[1]]
			}
		}
		return nods[0]
	}
	compareNodes := func(org, cpy *Node138b) bool {
		nodMap := make(map[*Node138b]*Node138b)
		cur1, cur2 := org, cpy
		for cur1 != nil {
			if cur1.Val != cur2.Val {
				return false
			}
			nodMap[cur1] = cur2
			cur1 = cur1.Next
			cur2 = cur2.Next
		}
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
			org := createList(tc.vals)
			cpy := copyRandomList138b(org)

			// Check if lists are identical but separate
			if !compareNodes(org, cpy) {
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
