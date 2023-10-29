package main

import (
	"fmt"
	"time"
)

func Sensor(ch chan<- float64) {
	i := 0.0
	for {
		i += 1.0
		ch <- i
		time.Sleep(100 * time.Millisecond)
	}
}

func SensorReader(ch <-chan float64, processedData chan<- float64) {
	sum := 0.0
	i := 0
	for {
		i++
		sum += <-ch
		if i%10 == 0 {
			processedData <- (sum / 10)
			sum = 0.0
		}
	}
}

func main() {
	sensorData := make(chan float64)
	processedData := make(chan float64)
	timer := time.NewTimer(1 * time.Minute)

	go Sensor(sensorData)
	go SensorReader(sensorData, processedData)

exitFor:
	for {
		select {
		case <-timer.C:
			close(sensorData)
			close(processedData)
			break exitFor
		case pd := <-processedData:
			fmt.Printf("%f\n", pd)
		}
	}
}
