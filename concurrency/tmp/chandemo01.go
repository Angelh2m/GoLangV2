package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := &sync.WaitGroup{}
	ch := make(chan int)
	wg.Add(2)
	go func(ch <-chan int, wg *sync.WaitGroup) {
		if v, ok := <-ch; ok {
			fmt.Println(v)
		}
		wg.Done()
	}(ch, wg)
	go func(ch chan<- int, wg *sync.WaitGroup) {
		ch <- 42
		close(ch)
		wg.Done()
	}(ch, wg)
	wg.Wait()
}
