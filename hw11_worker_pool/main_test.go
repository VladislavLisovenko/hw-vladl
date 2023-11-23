package main

import (
	"sync"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestIncrement(t *testing.T) {
	tests := []struct {
		name             string
		startValue       int
		iterationsNumber int
		expected         int
	}{
		{
			name:             "10 iterations from 0, expected 10",
			startValue:       0,
			iterationsNumber: 10,
			expected:         10,
		},
		{
			name:             "10 iterations from -5, expected 5",
			startValue:       -5,
			iterationsNumber: 10,
			expected:         5,
		},
		{
			name:             "-10 iterations from 0, expected 0",
			startValue:       0,
			iterationsNumber: -10,
			expected:         0,
		},
	}
	wg := &sync.WaitGroup{}
	mu := &sync.Mutex{}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := tt.startValue
			for i := 0; i < tt.iterationsNumber; i++ {
				wg.Add(1)
				go Increment(&n, wg, mu)
			}
			wg.Wait()
			require.Equal(t, tt.expected, n)
		})
	}
}
