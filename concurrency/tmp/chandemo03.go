package main

import (
	"fmt"
	"sync"
)

func main() {

	wg := &sync.WaitGroup{}
	ch := make(chan int)
	wg.Add(2) // ADD TWO GO RUNTINES

	go func(ch <-chan int, wg *sync.WaitGroup) { // SUBSCRIBE THE RECIVER
		for msg := range ch {
			fmt.Println(msg)
		}
		wg.Done()
	}(ch, wg)

	go func(ch chan<- int, wg *sync.WaitGroup) { // SUBSCRIBE THE SENDER
		for i := 0; i < 10; i++ {
			ch <- i
		}
		ch <- 40
		close(ch)
		wg.Done()
	}(ch, wg)

	wg.Wait()
}
