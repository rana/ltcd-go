package main

import "testing"

// Node133 represents a graph node with a value and list of neighbors
type Node133 struct {
	Val       int
	Neighbors []*Node133
}

// Time complexity: O(v + e), v is the number of vertexes. e is the number of edges. We process each vertex. We visit each edge.
// Space complexity: O(v), each vertex is stored in a `seen` map.
// https://claude.ai/chat/eba816b7-0d7a-47c6-ac9c-91bd5bf94a95
func cloneGraph133(nod *Node133) *Node133 {
	// Clone Graph
	// Given a node in a connected, undirected graph.
	// Create a deep copy.
	// Return a node of the deep copy.
	// Conditions:
	//	* Node connections are maintained
	//	* Node connections are to clones
	// 	* Cycles may be present
	// Use a map to handle cycles.
	// Map int-node copy.
	// Use depth-first search with recursion.
	// Use a helper function to support recursion.

	// Check input min edge case.
	if nod == nil {
		return nil
	}

	// Create a seen map, int-node copy.
	seen := make(map[int]*Node133)

	// Clone graph recursively.
	return cloneDFS133(nod, seen)
}

// cloneDFS133 recursively clones.
// Receives an original node and return a cloned node.
func cloneDFS133(org *Node133, seen map[int]*Node133) *Node133 {
	// Check for existing copy.
	if cpy, exists := seen[org.Val]; exists {
		return cpy
	}

	// Copy current node.
	cpy := &Node133{
		Val:       org.Val,
		Neighbors: make([]*Node133, 0, len(org.Neighbors)),
	}

	// Mark current node as seen.
	seen[cpy.Val] = cpy

	// Copy and connect neighbor nodes.
	for _, nbrOrg := range org.Neighbors {
		nbrCpy := cloneDFS133(nbrOrg, seen)
		cpy.Neighbors = append(cpy.Neighbors, nbrCpy)
	}

	// Return cpy.
	return cpy
}

// Test cases
func TestCloneGraph133(t *testing.T) {
	cases := []struct {
		name  string
		input []*Node133
		size  int
	}{
		{
			name:  "four node cycle",
			input: buildCycleGraph133(4),
			size:  4,
		},
		{
			name:  "single node",
			input: buildCycleGraph133(1),
			size:  1,
		},
		{
			name:  "empty graph",
			input: nil,
			size:  0,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			var org *Node133
			if len(tc.input) > 0 {
				org = tc.input[0]
			}

			cpy := cloneGraph133(org)

			// Verify nil case
			if tc.size == 0 {
				if cpy != nil {
					t.Error("expected nil for empty graph")
				}
				return
			}

			// Verify graph structure
			seen := make(map[int]bool)
			verifyClone133(t, org, cpy, seen)
		})
	}
}

// Helper to build test graphs
func buildCycleGraph133(size int) []*Node133 {
	if size == 0 {
		return nil
	}

	// Create nodes
	nodes := make([]*Node133, size)
	for i := 0; i < size; i++ {
		nodes[i] = &Node133{Val: i + 1}
	}

	// Connect nodes in a cycle
	for i := 0; i < size; i++ {
		prv := (i - 1 + size) % size
		nxt := (i + 1) % size
		nodes[i].Neighbors = []*Node133{nodes[prv], nodes[nxt]}
	}

	return nodes
}

// Helper to verify cloned graph structure
func verifyClone133(t *testing.T, org, cpy *Node133, seen map[int]bool) {
	if org == nil || cpy == nil {
		t.Error("unexpected nil node")
		return
	}

	// Check for cycles
	if seen[org.Val] {
		return
	}
	seen[org.Val] = true

	// Verify node properties
	if org.Val != cpy.Val {
		t.Errorf("value mismatch: got %d, want %d", cpy.Val, org.Val)
	}
	if len(org.Neighbors) != len(cpy.Neighbors) {
		t.Errorf("neighbor count mismatch: got %d, want %d",
			len(cpy.Neighbors), len(org.Neighbors))
	}

	// Verify neighbors recursively
	for i, nbrOrg := range org.Neighbors {
		nbrCpy := cpy.Neighbors[i]
		if nbrOrg.Val != nbrCpy.Val {
			t.Errorf("neighbor value mismatch: got %d, want %d",
				nbrCpy.Val, nbrOrg.Val)
		}
		verifyClone133(t, nbrOrg, nbrCpy, seen)
	}
}
