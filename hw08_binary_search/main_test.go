package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func tests() []struct {
	descr      string
	dataSource []int
	number     int
	expected   bool
} {
	return []struct {
		descr      string
		dataSource []int
		number     int
		expected   bool
	}{
		{
			descr:      "empty slice",
			dataSource: []int{},
			number:     4,
			expected:   false,
		},
		{
			descr:      "1-element slice contains number",
			dataSource: []int{3},
			number:     3,
			expected:   true,
		},
		{
			descr:      "1-element slice doesn't contains number",
			dataSource: []int{4},
			number:     3,
			expected:   false,
		},
		{
			descr:      "2-element slice contains number",
			dataSource: []int{3, 4},
			number:     3,
			expected:   true,
		},
		{
			descr:      "2-element slice doesn't contains number",
			dataSource: []int{4, 5},
			number:     3,
			expected:   false,
		},
		{
			descr:      "3-element slice contains number",
			dataSource: []int{3, 4, 5},
			number:     5,
			expected:   true,
		},
		{
			descr:      "3-element slice doesn't contains number",
			dataSource: []int{4, 5, 6},
			number:     3,
			expected:   false,
		},
		{
			descr:      "slice contains number",
			dataSource: []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
			number:     4,
			expected:   true,
		},
		{
			descr:      "slice doesn't contains number",
			dataSource: []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
			number:     99,
			expected:   false,
		},
	}
}

func TestSearch(t *testing.T) {
	tests := tests()

	for _, td := range tests {
		td := td
		t.Run(td.descr, func(t *testing.T) {
			got := Search(td.dataSource, td.number)
			require.Equal(t, td.expected, got)
		})
	}
}
