package main

import (
	"fmt"
)

func Search(data []int, number int) bool {
	dataLength := len(data)
	switch dataLength {
	case 0:
		return false
	case 1:
		return number == data[0]
	case 2:
		return number == data[0] || number == data[1]
	}

	if number < data[0] || number > data[dataLength-1] {
		return false
	}

	startIndex := 0
	lastIndex := dataLength - 1

	for startIndex <= lastIndex {
		i := startIndex + (lastIndex-startIndex)/2
		curNumber := data[i]

		if curNumber == number {
			return true
		}

		if curNumber > number {
			lastIndex = i
		} else {
			startIndex = i
		}

		if lastIndex-startIndex == 1 {
			return data[startIndex] == number || data[lastIndex] == number
		}
	}

	return false
}

func main() {
	data := []int{4, 5}
	number := 3

	res := Search(data, number)
	fmt.Println(res)
}
