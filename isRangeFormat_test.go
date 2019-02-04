package main

import "testing"

func TestRangeFormat(t *testing.T) {
	tables := []struct {
    value string
    res bool

	}{
		{"100&200", true},
		{"a&b", true},
		{"a&", true},
		{"a", false},
	}

	for _, table := range tables {
		result := isRangeFormat(table.value)
		if result != table.res {
			t.Errorf("Result was incorrect, got: %s, want: %s.", result, table.res)
		}
	}
}
