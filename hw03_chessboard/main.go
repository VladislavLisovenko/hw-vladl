package main

import "fmt"

func main() {
	size := 0
	fmt.Print("Enter chess board size: ")
	fmt.Scanf("%d", &size)
	Draw(size)
}

func Draw(size int) {
	s := ""
	for i := 1; i <= size; i++ {
		for j := 1; j <= size; j++ {
			if (i+j)%2 == 0 {
				s = " "
			} else {
				s = "#"
			}
			print(s)
		}
		println()
	}
}
