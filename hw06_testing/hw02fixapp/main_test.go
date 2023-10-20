package hw02fixapp

import (
	"hw02fixapp/reader"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func tests() []struct {
	descr    string
	path     string
	expected int
} {
	return []struct {
		descr    string
		path     string
		expected int
	}{
		{
			descr:    "File exists",
			path:     "data.json",
			expected: 2,
		},
		{
			descr:    "File doesn't exist, must be error",
			path:     "data1.json",
			expected: 0,
		},
	}
}

func TestReadJSON(t *testing.T) {
	tests := tests()

	for _, td := range tests {
		td := td
		t.Run(td.descr, func(t *testing.T) {

			if _, er := os.Stat(td.path); er != nil {
				require.Error(t, er)
			} else {
				got, err := reader.ReadJSON(td.path)
				if err != nil {
					require.NoError(t, err)
					return
				}
				require.Equal(t, td.expected, len(got))
			}

		})
	}
}
