package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func tests() []struct {
	descr         string
	dataSource    []int
	number        int
	expected      bool
	expectedIndex int
} {
	return []struct {
		descr         string
		dataSource    []int
		number        int
		expected      bool
		expectedIndex int
	}{
		{
			descr:         "empty slice",
			dataSource:    []int{},
			number:        4,
			expected:      false,
			expectedIndex: -1,
		},
		{
			descr:         "1-element slice contains number",
			dataSource:    []int{3},
			number:        3,
			expected:      true,
			expectedIndex: 0,
		},
		{
			descr:         "1-element slice doesn't contains number",
			dataSource:    []int{4},
			number:        3,
			expected:      false,
			expectedIndex: -1,
		},
		{
			descr:         "2-element slice contains number",
			dataSource:    []int{3, 4},
			number:        3,
			expected:      true,
			expectedIndex: 0,
		},
		{
			descr:         "2-element slice doesn't contains number",
			dataSource:    []int{4, 5},
			number:        3,
			expected:      false,
			expectedIndex: -1,
		},
		{
			descr:         "3-element slice contains number",
			dataSource:    []int{3, 4, 5},
			number:        5,
			expected:      true,
			expectedIndex: 2,
		},
		{
			descr:         "3-element slice doesn't contains number",
			dataSource:    []int{4, 5, 6},
			number:        3,
			expected:      false,
			expectedIndex: -1,
		},
		{
			descr:         "slice contains number",
			dataSource:    []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
			number:        4,
			expected:      true,
			expectedIndex: 3,
		},
		{
			descr:         "slice doesn't contains number",
			dataSource:    []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
			number:        99,
			expected:      false,
			expectedIndex: -1,
		},
		{
			descr:         "number less first number",
			dataSource:    []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
			number:        0,
			expected:      false,
			expectedIndex: -1,
		},
		{
			descr:         "number more last number",
			dataSource:    []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
			number:        15,
			expected:      false,
			expectedIndex: -1,
		},
	}
}

func TestSearch(t *testing.T) {
	tests := tests()

	for _, td := range tests {
		td := td
		t.Run(td.descr, func(t *testing.T) {
			gotIndex, got := Search(td.dataSource, td.number)
			require.Equal(t, td.expected, got)
			require.Equal(t, td.expectedIndex, gotIndex)
		})
	}
}
