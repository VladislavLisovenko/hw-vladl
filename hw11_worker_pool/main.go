package main

import (
	"fmt"
	"sync"
)

func Increment(n *int, wg *sync.WaitGroup, mu *sync.Mutex) {
	defer wg.Done()
	mu.Lock()
	*n++
	fmt.Println(*n)
	mu.Unlock()
}

func main() {
	wg := &sync.WaitGroup{}
	mu := &sync.Mutex{}

	v := 0
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go Increment(&v, wg, mu)
	}

	wg.Wait()
	fmt.Println(v)
}
