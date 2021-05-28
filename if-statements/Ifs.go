package main

import "fmt"

func main() {

	
	states := map[string]int{
		"California": 50000,
		"Texas": 30000,
	}

	if pop,ok := states["California"]; ok {
		fmt.Println(pop)
	}
	
}