package main

import "fmt"

type Doctor struct {
	number int
	actorName string
	companions []string
}

func main() {

	aDoctor := Doctor{
		number: 9,
		actorName: "John",
		companions: []string{
			"One",
			"two",
		},
	}

	fmt.Println(aDoctor)

	// Anonymous struct
	aDoctor2 := struct {
		name string
	}{
		name: "John3",
	}

	fmt.Println(aDoctor2)
}