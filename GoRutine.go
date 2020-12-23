package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wgroup = sync.WaitGroup{}
var m =  sync.RWMutex{}
var counter = 0

func main() {

	//wgroup.Add(1)
	//go sayHi()
	//wgroup.Wait()
	//fmt.Println("====================")

	runtime.GOMAXPROCS(100)

	for i := 0; i < 10; i++ {

		wgroup.Add(1) // Specify the routine count
		m.RLock() // Go routine starts

		go sayHi()

		m.Lock() // Wait for the go routine to complete
		m.Unlock() // Once complete unlock and proceed

	}
	wgroup.Wait()

}

func sayHi()  {


	fmt.Println("Say Hello!", counter)
	counter++

	m.RUnlock() // Go routine ends writing
	wgroup.Done()
}