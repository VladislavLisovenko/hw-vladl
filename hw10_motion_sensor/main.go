package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"time"
)

func randFloat64() float64 {
	max := big.NewInt(1_000_000_000)
	randInt, err := rand.Int(rand.Reader, max)
	if err != nil {
		panic(err)
	}
	return float64(randInt.Int64())
}

func Sensor(d time.Duration) chan float64 {
	// создаем канал
	res := make(chan float64)
	// запускаем горутину
	go func() {
		// закроем канал перед выходом из функции
		defer close(res)
		// создаем таймер
		// After - помещает значение в канал по прошествии заданного временни
		timer := time.After(d)
		// бесконечный цикл
		for {
			select {
			// помещаем в канал очередное значение
			case res <- randFloat64():
			// как только значение таймера прочитано - завершаем выполнение функции
			case <-timer:
				return
			}
			// пауза 100 миллисекунд
			time.Sleep(100 * time.Millisecond)
		}
	}()
	// возвращаем канал
	return res
}

// в параметре sensorData ReadOnly-канал.
func SensorReader(sensorData <-chan float64) chan float64 {
	// создаем канал
	res := make(chan float64)

	// запускаем горутину
	go func() {
		// закроем канал перед выходом из функции
		defer close(res)
		// инициализация счетчика итераций цикла, по которому будем определять каждое 10-е значение
		i := 0
		// инициализация переменной для суммирования получаемых из сенсора данных
		var accum float64
		// получаем даные из сенсора
		for v := range sensorData {
			// инкремент счетчика
			i++
			// прибавляем очередное значение, полученное из сенсора
			accum += v
			// если получено 10-е значение, то вычисляем среднее значение,
			// помещаем его в результирующий канал и обнуляем переменную суммы
			if i%10 == 0 {
				res <- accum / 10.0
				accum = 0.0
			}
		}
	}()

	// возвращаем канал
	return res
}

func main() {
	// запускаем счетчик
	dataCh := Sensor(time.Minute)
	// передаем ссылку на канал сенсора в ридер и получаем ссылку на канал со средними значениями
	acumCh := SensorReader(dataCh)

	// выводим на экран значения из канала со средними значениями
	for d := range acumCh {
		fmt.Println(d)
	}
}
