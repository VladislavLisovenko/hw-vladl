package main

import (
	"fmt"
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
		{
			descr:    "5th iteration",
			expected: 45.5,
		},
	}
}

func DataGenerator() chan float64 {
	res := make(chan float64)
	go func() {
		defer close(res)
		for i := 1; i <= 50; i++ {
			res <- float64(i)
		}
	}()
	return res
}

func TestSensor(t *testing.T) {
	sensorDataCh := Sensor(time.Second * 5)
	var float64Value float64
	for v := range sensorDataCh {
		fmt.Println(v)
		require.IsType(t, float64Value, v)
	}
}

func TestSensorReader(t *testing.T) {
	tests := sensorTestData()
	sensorCh := DataGenerator()
	sensorReaderCh := SensorReader(sensorCh)
	i := 0
	for v := range sensorReaderCh {
		fmt.Println(v)
		require.Equal(t, tests[i].expected, v)
		i++
	}
}
