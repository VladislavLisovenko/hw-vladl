package hw07wordcounter

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func tests() []struct {
	descr    string
	text     string
	expected map[string]int
} {
	return []struct {
		descr    string
		text     string
		expected map[string]int
	}{
		{
			descr:    "empty string",
			text:     "",
			expected: map[string]int{},
		},
		{
			descr: "two different words",
			text:  "one?, \n two!\n",
			expected: map[string]int{
				"one": 1,
				"two": 1,
			},
		},
		{
			descr: "two repeated words",
			text:  "one, 	Two?.\n    one-\rtwo. One",
			expected: map[string]int{
				"one": 3,
				"two": 2,
			},
		},
	}
}

func TestCountWords(t *testing.T) {
	tests := tests()

	for _, td := range tests {
		td := td
		t.Run(td.descr, func(t *testing.T) {
			got := countWords(td.text)
			require.Equal(t, td.expected, got)
		})
	}
}
