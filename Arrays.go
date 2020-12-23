package main

import (
	"fmt"
	"log"
)

func main() {
	grades := [3]int{1,2,3}
	grades2 := [...]int{1,2,3,4}

	log.Println(grades)
	log.Println(grades2)


	b := &grades

	b[0] = 9

	fmt.Println(b)
	log.Println(grades)

}
