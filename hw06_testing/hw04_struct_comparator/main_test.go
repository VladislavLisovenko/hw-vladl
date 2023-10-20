package hw04structcomparator

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func tests() []struct {
	descr    string
	b1       *Book
	b2       *Book
	method   CompareMethod
	expected bool
} {
	return []struct {
		descr    string
		b1       *Book
		b2       *Book
		method   CompareMethod
		expected bool
	}{
		{
			descr: "By rate",
			b1: &Book{
				id:     1,
				title:  "Book 1",
				author: "Author 1",
				year:   2020,
				size:   30,
				rate:   3.1,
			},
			b2: &Book{
				id:     2,
				title:  "Book 2",
				author: "Author 2",
				year:   2023,
				size:   40,
				rate:   2.1,
			},
			method:   ByRate,
			expected: true,
		},
		{
			descr: "By size",
			b1: &Book{
				id:     1,
				title:  "Book 1",
				author: "Author 1",
				year:   2020,
				size:   30,
				rate:   3.1,
			},
			b2: &Book{
				id:     2,
				title:  "Book 2",
				author: "Author 2",
				year:   2023,
				size:   40,
				rate:   2.1,
			},
			method:   BySize,
			expected: false,
		},
		{
			descr: "By year",
			b1: &Book{
				id:     1,
				title:  "Book 1",
				author: "Author 1",
				year:   2020,
				size:   30,
				rate:   3.1,
			},
			b2: &Book{
				id:     2,
				title:  "Book 2",
				author: "Author 2",
				year:   2023,
				size:   40,
				rate:   2.1,
			},
			method:   ByYear,
			expected: false,
		},
	}
}

func TestCompare(t *testing.T) {
	tests := tests()

	for _, td := range tests {
		td := td
		t.Run(td.descr, func(t *testing.T) {
			comparator := NewBookComparator(td.method)
			got, err := comparator.Compare(td.b1, td.b2)
			if err != nil {
				require.NoError(t, err)
				return
			}
			require.Equal(t, td.expected, got)
		})
	}
}
