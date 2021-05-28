package main

import (
	"fmt"
	"sync"

	"../scopes"
)

func main() {

	fmt.Println("RUN")
	fmt.Println(scopes.Three)

	wg := &sync.WaitGroup{} // Wait group to make go rutines execute.
	ch := make(chan int)
	wg.Add(2) // ADD TWO GO RUNTINES TO WAIT GROUP

	go func(ch <-chan int, wg *sync.WaitGroup) { // SUBSCRIBE THE RECIVER
		for msg := range ch {
			fmt.Println(msg)
		}
		wg.Done()
	}(ch, wg)

	go func(ch chan<- int, wg *sync.WaitGroup) { // SUBSCRIBE THE SENDER
		for i := 0; i <= 10; i++ {
			ch <- i
		}
		for {
		}
		// ch <- 40
		close(ch)
		wg.Done()
	}(ch, wg)

	wg.Wait()
}
