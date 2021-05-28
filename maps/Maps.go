package main

import "fmt"

func main() {

	// Pre define the map
	statePopulations := make(map[string]int)

	// Maps are passed by reference
	statePopulations = map[string]int{
		"California": 3988888,
		"Texas": 3000000,
		"Florida": 250000,
	}

	statePopulations["Georgia"] = 100000

	fmt.Println(statePopulations)

	// This will return zero
	delete(statePopulations, "Florida")

	// This will check for existance
	_, ok := statePopulations["Florida"]
	fmt.Println(ok)



	fmt.Println(statePopulations)
	fmt.Println(statePopulations["Florida"])
	fmt.Println(statePopulations["Texas"])
}
