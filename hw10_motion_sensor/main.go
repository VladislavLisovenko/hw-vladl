package main

import (
	"context"
	"fmt"
	"time"
)

func Sensor(ctx context.Context, ch chan<- float64) {
	i := 0.0
	stop := false
	for !stop {
		i += 1.0
		select {
		case ch <- i:
		case <-ctx.Done():
			stop = true
		}
		time.Sleep(100 * time.Millisecond)
	}
}

func SensorReader(ctx context.Context, ch <-chan float64, processedData chan<- float64) {
	sum := 0.0
	v := 0.0
	i := 0
	for {
		i++
		select {
		case v = <-ch:
			sum += v
			if i%10 == 0 {
				processedData <- (sum / 10)
				sum = 0.0
			}
		case <-ctx.Done():
			return
		}
	}
}

func main() {
	sensorData := make(chan float64)
	processedData := make(chan float64)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	go Sensor(ctx, sensorData)
	go SensorReader(ctx, sensorData, processedData)

	for {
		select {
		case pd := <-processedData:
			fmt.Printf("%f\n", pd)
		case <-ctx.Done():
			return
		}
	}
}
