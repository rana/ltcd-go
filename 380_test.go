package main

import (
	"math/rand"
	"testing"
)

// Insert Delete GetRandom O(1)
// Implement RandomizedSet380 operations
// insert, delete, getRandom with
// O(1) time complexity.
// Use an array and a map.
// For delete, swap the last element with the removal element.
// https://chatgpt.com/c/671549b4-6a30-8002-8548-a1fb6596dbfe
type RandomizedSet380 struct {
	arr []int
	mp  map[int]int
}

func Constructor380() RandomizedSet380 {
	return RandomizedSet380{
		arr: []int{},
		mp:  make(map[int]int),
	}
}

func (rs *RandomizedSet380) Insert(val int) bool {
	// Check whether the element already exists.
	if _, exists := rs.mp[val]; exists {
		return false
	}

	// Add the new element to the end of the array.
	rs.arr = append(rs.arr, val)
	// Map the element to the index.
	rs.mp[val] = len(rs.arr) - 1

	return true
}

func (rs *RandomizedSet380) Remove(val int) bool {
	// Check whether the element already exists.
	idx, exists := rs.mp[val]
	if !exists {
		return false
	}

	// Swap the last element with the removal element.

	// Get the last element.
	lstVal := rs.arr[len(rs.arr)-1]
	// Overwrite the array removal element with the last element.
	rs.arr[idx] = lstVal
	// Remove the duplicate last element.
	rs.arr = rs.arr[:len(rs.arr)-1]
	// Remove the removal element from the map.
	delete(rs.mp, val)
	// Update the last element to index mapping.
	rs.mp[lstVal] = idx

	return true
}

func (rs *RandomizedSet380) GetRandom() int {
	// Create a random index.
	rndIdx := rand.Intn(len(rs.arr))
	// Return a random element.
	return rs.arr[rndIdx]
}

func TestRandomizedSet380(t *testing.T) {
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
			rs := Constructor380()

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

func TestRandomizedSet_Random380(t *testing.T) {
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
			rs := Constructor380()
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
