package hw05shapes

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func tests() []struct {
	descr    string
	s        any
	expected float64
} {
	return []struct {
		descr    string
		s        any
		expected float64
	}{
		{
			descr: "Circle",
			s: &Circle{
				radius: 5,
			},
			expected: 78.53981633974483,
		},
		{
			descr: "Resctangle",
			s: &Rectangle{
				width:  10,
				height: 5,
			},
			expected: 50,
		},
		{
			descr: "Triangle",
			s: &Triangle{
				base:   8,
				height: 6,
			},
			expected: 24,
		},
		{
			descr: "Square (must be error)",
			s: &Square{
				side: 5,
			},
			expected: 25,
		},
	}
}

func TestCalculateArea(t *testing.T) {
	tests := tests()

	for _, td := range tests {
		td := td
		t.Run(td.descr, func(t *testing.T) {
			got, err := CalculateArea(td.s)
			switch td.s.(type) {
			case Shape:
				require.Equal(t, td.expected, got)
			default:
				require.Error(t, err)
			}
		})
	}
}
