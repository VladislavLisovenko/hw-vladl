package main

import (
	"fmt"
)

func Search(data []int, number int) (int, bool) {
	dataLength := len(data)
	switch dataLength {
	case 0:
		return -1, false
	case 1:
		if number == data[0] {
			return 0, true
		}
		return -1, false
	case 2:
		if number != data[0] && number != data[1] {
			return -1, false
		}
		if number == data[0] {
			return 0, true
		}
		return 1, true
	}

	if number < data[0] || number > data[dataLength-1] {
		return -1, false
	}

	startIndex := 0
	lastIndex := dataLength - 1

	for startIndex <= lastIndex {
		i := startIndex + (lastIndex-startIndex)/2
		curNumber := data[i]

		if curNumber == number {
			return i, true
		}

		if curNumber > number {
			lastIndex = i
		} else {
			startIndex = i
		}

		if lastIndex-startIndex == 1 {
			switch number {
			case data[startIndex]:
				return startIndex, true
			case data[lastIndex]:
				return lastIndex, true
			default:
				return -1, false
			}
		}
	}

	return -1, false
}

func main() {
	data := []int{4, 5}
	number := 4

	ind, res := Search(data, number)
	fmt.Println(ind, res)
}
