package main

import "fmt"

type Animal struct {
	Name string
	Origin string
}

type Bird struct {
	Animal
	CanFly bool
}

func main() {

	bird := Bird{}
	bird.Name = "Hello"
	bird.CanFly = true


	fmt.Println(bird.Name)

}