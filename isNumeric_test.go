package main

import "testing"

func TestNumeric(t *testing.T) {
	tables := []struct {
    value string
    res bool

	}{
		{"100", true},
		{"a", false},
		{"1", true},
		{"ab&c", false},
	}

	for _, table := range tables {
		result := isNumeric(table.value)
		if result != table.res {
			t.Errorf("Result was incorrect, got: %s, want: %s.", result, table.res)
		}
	}
}
