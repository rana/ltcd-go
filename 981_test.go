package main

import (
	"reflect"
	"sort"
	"testing"
)

// Time Based Key-Value Store
// key value and integer timestamp with value.
// get: return prv_tme <= srch_tme; or empty string
// Use Binary search for value.
// Use a struct for value, make array of it
// https://console.anthropic.com/workbench/67d6ed1b-5730-416f-bcc3-a82ac0a91101
type TimeEntry981 struct {
	val string
	tme int
}

type TimeMap981 struct {
	dat map[string][]TimeEntry981
}

func Constructor981() TimeMap981 {
	return TimeMap981{
		dat: make(map[string][]TimeEntry981),
	}
}
func (tm *TimeMap981) Set(key string, val string, tme int) {
	tm.dat[key] = append(tm.dat[key], TimeEntry981{val: val, tme: tme})
}

// Get retrieves the value associated with key at or before timestamp
func (tm *TimeMap981) Get(key string, tme int) string {
	// If key doesn't exist, return empty string
	ent, ok := tm.dat[key]
	if !ok || len(ent) == 0 {
		return ""
	}

	// Find the rightmost position where ent time <= timestamp
	idx := sort.Search(len(ent), func(i int) bool {
		return ent[i].tme > tme
	})

	// If idx is 0, all elements are > timestamp
	if idx == 0 {
		return ""
	}

	// Return the element just before the first element > timestamp
	return ent[idx-1].val
}

func TestTimeMap981(t *testing.T) {
	tests := []struct {
		name     string
		commands []string
		args     [][]any
		expected []any
	}{
		{
			name:     "Example 1",
			commands: []string{"TimeMap", "set", "get", "get", "set", "get", "get"},
			args: [][]any{
				{},
				{"foo", "bar", 1},
				{"foo", 1},
				{"foo", 3},
				{"foo", "bar2", 4},
				{"foo", 4},
				{"foo", 5},
			},
			expected: []any{nil, nil, "bar", "bar", nil, "bar2", "bar2"},
		},
		{
			name:     "Key doesn't exist",
			commands: []string{"TimeMap", "get"},
			args: [][]any{
				{},
				{"nonexistent", 1},
			},
			expected: []any{nil, ""},
		},
		{
			name:     "Multiple keys",
			commands: []string{"TimeMap", "set", "set", "get", "get"},
			args: [][]any{
				{},
				{"key1", "value1", 5},
				{"key2", "value2", 10},
				{"key1", 5},
				{"key2", 10},
			},
			expected: []any{nil, nil, nil, "value1", "value2"},
		},
		{
			name:     "Timestamp before any entry",
			commands: []string{"TimeMap", "set", "get"},
			args: [][]any{
				{},
				{"key", "value", 100},
				{"key", 50},
			},
			expected: []any{nil, nil, ""},
		},
		{
			name:     "Multiple timestamps for same key",
			commands: []string{"TimeMap", "set", "set", "set", "get", "get", "get", "get"},
			args: [][]any{
				{},
				{"key", "value1", 10},
				{"key", "value2", 20},
				{"key", "value3", 30},
				{"key", 5},
				{"key", 15},
				{"key", 25},
				{"key", 35},
			},
			expected: []any{nil, nil, nil, nil, "", "value1", "value2", "value3"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var tm TimeMap981
			var result any

			for i, cmd := range tt.commands {
				switch cmd {
				case "TimeMap":
					tm = Constructor981()
					result = nil
				case "set":
					key := tt.args[i][0].(string)
					value := tt.args[i][1].(string)
					timestamp := tt.args[i][2].(int)
					tm.Set(key, value, timestamp)
					result = nil
				case "get":
					key := tt.args[i][0].(string)
					timestamp := tt.args[i][1].(int)
					result = tm.Get(key, timestamp)
				}

				if !reflect.DeepEqual(result, tt.expected[i]) {
					t.Errorf("Operation %d (%s): got %v, want %v", i, cmd, result, tt.expected[i])
				}
			}
		})
	}
}
