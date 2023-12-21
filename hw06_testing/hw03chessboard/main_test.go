package hw03chessboard

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func tests() []struct {
	descr    string
	size     int
	expected string
} {
	return []struct {
		descr    string
		size     int
		expected string
	}{
		{
			descr:    "2x2 board",
			size:     2,
			expected: " #\n# ",
		},
		{
			descr:    "3x3 board",
			size:     3,
			expected: " # \n# #\n # ",
		},
		{
			descr:    "4x4 board",
			size:     4,
			expected: " # #\n# # \n # #\n# # ",
		},
		{
			descr:    "8x8 board",
			size:     8,
			expected: " # # # #\n# # # # \n # # # #\n# # # # \n # # # #\n# # # # \n # # # #\n# # # # ",
		},
	}
}

func TestDraw(t *testing.T) {
	tests := tests()

	for _, td := range tests {
		td := td
		t.Run(td.descr, func(t *testing.T) {
			got := Draw(td.size)
			require.Equal(t, td.expected, got)
		})
	}
}
