package main

import (
	"reflect"
	"testing"
)

// Time Based Key-Value Store
// set: key, val
// get: key, val, tme. prv_tme <= srch_tme
// Use a struct entry value with val,tme

type TimeMap981b struct {
}

func Constructor981b() TimeMap981b {
	return TimeMap981b{}
}
func (tm *TimeMap981b) Set(key string, val string, tme int) {

}
func (tm *TimeMap981b) Get(key string, tme int) string {
	return ""
}
func TestTimeMap981b(t *testing.T) {
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
			var tm TimeMap981b
			var result any

			for i, cmd := range tt.commands {
				switch cmd {
				case "TimeMap":
					tm = Constructor981b()
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
