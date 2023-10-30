package main

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func sensorTestData() []struct {
	descr    string
	expected float64
} {
	return []struct {
		descr    string
		expected float64
	}{
		{
			descr:    "1st iteration",
			expected: 5.5,
		},
		{
			descr:    "2nd iteration",
			expected: 15.5,
		},
		{
			descr:    "3rd iteration",
			expected: 25.5,
		},
		{
			descr:    "4th iteration",
			expected: 35.5,
		},
	}
}

func TestSensorReader(t *testing.T) {
	sensorData := make(chan float64)
	processedData := make(chan float64)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	go Sensor(ctx, sensorData)
	go SensorReader(ctx, sensorData, processedData)

	testData := sensorTestData()
	for _, td := range testData {
		td := td
		t.Run(td.descr, func(t *testing.T) {
			select {
			case got := <-processedData:
				require.Equal(t, td.expected, got)
			case <-ctx.Done():
				return
			}
		})
	}
}
