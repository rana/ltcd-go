package main

import (
	"math/rand"
	"testing"
)

// Insert Delete GetRandom O(1)
// Design RanomizedSet with all operations
// with average time complexity of O(1).
// Conditions:
// * insert(val int) bool
// * remove(val int) bool
// * getRandom() int
// Use map[int]int val-idx for insert, remove.
// Use []int of vals for getRandom().
// Use package "rand" for array index generation.

type RandomizedSet380b struct {
	mp  map[int]int
	arr []int
}

func Constructor380b() RandomizedSet380b {
	return RandomizedSet380b{
		mp: make(map[int]int),
	}
}

// Inserts an item val into the set if not present.
// Returns true if the item was not present.
func (rs *RandomizedSet380b) Insert(val int) bool {
	_, exists := rs.mp[val]
	if exists {
		return false
	}
	rs.mp[val] = len(rs.arr)
	rs.arr = append(rs.arr, val)
	return true
}

// Removes an item val from the set if present.
// Returns true if the item was present.
func (rs *RandomizedSet380b) Remove(val int) bool {
	idx, exists := rs.mp[val]
	if !exists {
		return false
	}
	lstVal := rs.arr[len(rs.arr)-1]
	rs.arr[idx] = lstVal
	rs.arr = rs.arr[:len(rs.arr)-1]
	delete(rs.mp, val)
	rs.mp[lstVal] = idx
	return true
}

// Returns a random element from the
// current set of elements
func (rs *RandomizedSet380b) GetRandom() int {
	idx := rand.Intn(len(rs.arr))
	return rs.arr[idx]
}

func TestRandomizedSet380b(t *testing.T) {
	tests := []struct {
		name     string
		ops      []string
		args     []int
		expected []bool
	}{
		{
			name:     "Test Insert",
			ops:      []string{"Insert", "Insert"},
			args:     []int{1, 1},
			expected: []bool{true, false},
		},
		{
			name:     "Test Remove",
			ops:      []string{"Insert", "Remove", "Remove"},
			args:     []int{1, 1, 1},
			expected: []bool{true, true, false},
		},
		{
			name:     "Test Insert Remove Sequence",
			ops:      []string{"Insert", "Remove", "Insert"},
			args:     []int{1, 1, 1},
			expected: []bool{true, true, true},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rs := Constructor380b()

			for i, op := range tt.ops {
				var result bool

				switch op {
				case "Insert":
					result = rs.Insert(tt.args[i])
				case "Remove":
					result = rs.Remove(tt.args[i])
				}

				if result != tt.expected[i] {
					t.Errorf("%s: operation %d: got %v, want %v",
						tt.name, i, result, tt.expected[i])
				}
			}
		})
	}
}

func TestRandomizedSet_Random380b(t *testing.T) {
	tests := []struct {
		name     string
		inserts  []int
		validSet map[int]bool
	}{
		{
			name:     "Test GetRandom with 3 numbers",
			inserts:  []int{1, 2, 3},
			validSet: map[int]bool{1: true, 2: true, 3: true},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rs := Constructor380b()
			for _, num := range tt.inserts {
				rs.Insert(num)
			}

			rnd := rs.GetRandom()
			if !tt.validSet[rnd] {
				t.Errorf("%s: GetRandom() = %d, want one of %v",
					tt.name, rnd, tt.validSet)
			}
		})
	}
}
