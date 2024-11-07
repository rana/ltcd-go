package main

import (
	"math/rand"
	"testing"
)

// Insert Delete GetRandom O(1)
// Implement RandomizedSet operations
// insert, delete, getRandom with
// O(1) time complexity.
// Use an array and a map.
// For delete, swap the last element with the removal element.
// https://chatgpt.com/c/671549b4-6a30-8002-8548-a1fb6596dbfe
type RandomizedSet struct {
	arr []int
	mp  map[int]int
}

func Constructor() RandomizedSet {
	return RandomizedSet{
		arr: []int{},
		mp:  make(map[int]int),
	}
}

func (rs *RandomizedSet) Insert(val int) bool {
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

func (rs *RandomizedSet) Remove(val int) bool {
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
	// Update the last element to index mapping.
	rs.mp[lstVal] = idx
	// Remove the duplicate last element.
	rs.arr = rs.arr[:len(rs.arr)-1]
	// Remove the removal element from the map.
	delete(rs.mp, val)

	return true
}

func (rs *RandomizedSet) GetRandom() int {
	// Create a random index.
	rndIdx := rand.Intn(len(rs.arr))
	// Return a random element.
	return rs.arr[rndIdx]
}

func TestInsert380(t *testing.T) {
	rs := Constructor()
	if !rs.Insert(1) {
		t.Errorf("Insert(1) should return true")
	}
	if rs.Insert(1) {
		t.Errorf("Insert(1) should return false when inserting duplicate")
	}
}

func TestRemove380(t *testing.T) {
	rs := Constructor()
	rs.Insert(1)
	if !rs.Remove(1) {
		t.Errorf("Remove(1) should return true")
	}
	if rs.Remove(1) {
		t.Errorf("Remove(1) should return false when removing non-existent element")
	}
}

func TestGetRandom(t *testing.T) {
	rs := Constructor()
	rs.Insert(1)
	rs.Insert(2)
	rs.Insert(3)
	rnd := rs.GetRandom()
	if rnd != 1 && rnd != 2 && rnd != 3 {
		t.Errorf("GetRandom() returned invalid value")
	}
}

func TestInsertRemoveGetRandom(t *testing.T) {
	rs := Constructor()
	if !rs.Insert(1) {
		t.Errorf("Insert(1) should return true")
	}
	if !rs.Remove(1) {
		t.Errorf("Remove(1) should return true")
	}
	if rs.Insert(1) != true {
		t.Errorf("Insert(1) should return true after removal")
	}
}
