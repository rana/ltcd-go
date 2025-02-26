package main

import "testing"

// https://claude.ai/chat/8ec09d6f-cc23-4f9d-9837-f8341f760caf
type LRUCache146 struct {
	// LRU Cache
	// Design a data structure with least recently used conditions.
	// Conditions:
	//	* O(1) time complexity for all operations.
	// Use helper function remove, addToFront for LRU logic.
	// Use a doubly-linked list for `remove` O(1) average time complexity.
	// Use a map for `get` O(1) average time complexity.
	itms map[int]*Node146
	hed  *Node146
	tal  *Node146
	cap  int
}

// Node146 is a doubly-linked list node.
type Node146 struct {
	key int
	val int
	prv *Node146
	nxt *Node146
}

func Constructor146(cap int) LRUCache146 {
	return LRUCache146{
		cap:  cap,
		itms: make(map[int]*Node146),
	}
}

// Use helper function remove, addToFront for LRU logic.

// remove removes a node from a doubly-linked list.
// No map operations.
func (c *LRUCache146) remove(nod *Node146) {
	// Handle next or head.
	if nod.prv != nil {
		// Node is non-head.
		nod.prv.nxt = nod.nxt
	} else {
		// Node is head.
		c.hed = nod.nxt
	}

	// Handle previous or tail.
	if nod.nxt != nil {
		// Node is non-tail.
		nod.nxt.prv = nod.prv
	} else {
		// Node is tail.
		c.tal = nod.prv
	}
}

// addToFront adds a node to the front of a doubly-linked list.
// No map operations.
func (c *LRUCache146) addToFront(nod *Node146) {
	// Update node pointers.
	nod.prv = nil
	nod.nxt = c.hed

	// Update head.
	if c.hed != nil {
		c.hed.prv = nod
	}
	c.hed = nod

	// Update tail.
	if c.tal == nil {
		c.tal = nod
	}
}

// Conditions:
//   - O(1) average time complexity.
//   - Return value for key; or, -1.
func (c *LRUCache146) Get(key int) int {
	// Check whether the item exists.
	if nod, exists := c.itms[key]; exists {
		// Remove and add for least recently used logic.
		c.remove(nod)
		c.addToFront(nod)
		return nod.val
	}
	return -1
}

// Conditions:
//   - O(1) average time complexity.
//   - If item exists, update existing value.
//   - If item doesn't exist, add a new item.
//   - If number of items exceeds capacity, remove tail.
func (c *LRUCache146) Put(key int, val int) {
	// Check whether item exists.
	if nod, exists := c.itms[key]; exists {
		// Update value.
		nod.val = val
		// Remove and add for least recently used logic.
		c.remove(nod)
		c.addToFront(nod)
		return
	}

	// Create a new item.
	nod := &Node146{key: key, val: val}
	c.itms[key] = nod
	c.addToFront(nod)

	// Check over capacity.
	if len(c.itms) > c.cap {
		// Remove tail item.
		delete(c.itms, c.tal.key)
		c.remove(c.tal)
	}
}

func TestLRUCache146(t *testing.T) {
	tests := []struct {
		name     string
		ops      []string
		args     [][]int
		expected []int
	}{
		{
			name:     "Example from problem",
			ops:      []string{"LRUCache", "put", "put", "get", "put", "get", "put", "get", "get", "get"},
			args:     [][]int{{2}, {1, 1}, {2, 2}, {1}, {3, 3}, {2}, {4, 4}, {1}, {3}, {4}},
			expected: []int{0, 0, 0, 1, 0, -1, 0, -1, 3, 4},
		},
		{
			name:     "Single capacity cache",
			ops:      []string{"LRUCache", "put", "get", "put", "get", "get"},
			args:     [][]int{{1}, {1, 1}, {1}, {2, 2}, {1}, {2}},
			expected: []int{0, 0, 1, 0, -1, 2},
		},
		{
			name:     "Update existing key",
			ops:      []string{"LRUCache", "put", "put", "get", "put", "get"},
			args:     [][]int{{2}, {1, 1}, {1, 2}, {1}, {3, 3}, {1}},
			expected: []int{0, 0, 0, 2, 0, 2},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			var cache LRUCache146
			for i, op := range tc.ops {
				var result int
				switch op {
				case "LRUCache":
					cache = Constructor146(tc.args[i][0])
				case "put":
					cache.Put(tc.args[i][0], tc.args[i][1])
				case "get":
					result = cache.Get(tc.args[i][0])
					if result != tc.expected[i] {
						t.Errorf("Operation %d: expected %d, got %d", i, tc.expected[i], result)
					}
				}
			}
		})
	}
}
