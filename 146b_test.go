package main

import "testing"

type LRUCache146b struct {
}

func Constructor146b(cap int) LRUCache146b {
	return LRUCache146b{}
}

func (c *LRUCache146b) Get(key int) int {
	return 0
}

func (c *LRUCache146b) Put(key int, val int) {

}

func TestLRUCache146b(t *testing.T) {
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
			var cache LRUCache146b
			for i, op := range tc.ops {
				var result int
				switch op {
				case "LRUCache":
					cache = Constructor146b(tc.args[i][0])
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
