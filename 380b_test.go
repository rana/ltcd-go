package main

import (
	"math/rand"
	"testing"
)

// Insert Delete GetRandom O(1)
// Each operation in O(1) time complexity.
// Use an array and map to acheive.
type RandomizedSet struct {
	arr []int
	mp  map[int]int
}

func Constructor380b() RandomizedSet {
	return RandomizedSet{
		arr: []int{},
		mp:  make(map[int]int),
	}
}

// Insert if not present. Return true if inserted.
func (rs *RandomizedSet) Insert(val int) bool {
	if _, exists := rs.mp[val]; exists {
		return false
	}

	rs.mp[val] = len(rs.arr)
	rs.arr = append(rs.arr, val)
	return true
}

// Remove val if present. Return true if removed.
func (rs *RandomizedSet) Remove(val int) bool {
	idx, exists := rs.mp[val]
	if !exists {
		return false
	}

	// Swap last value with removed value.
	lstVal := rs.arr[len(rs.arr)-1]
	rs.mp[lstVal] = idx
	rs.arr[idx] = lstVal
	// Remove target value.
	rs.arr = rs.arr[:len(rs.arr)-1]
	delete(rs.mp, val)
	return true
}

func (rs *RandomizedSet) GetRandom() int {
	rndIdx := rand.Intn(len(rs.arr))
	return rs.arr[rndIdx]
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
