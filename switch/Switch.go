package main

import "fmt"

func main() {

	switch 1 {
	case 1,4,5 :
		fmt.Println("1,4,5")
	case 6, 7, 8:
		fmt.Println("7,8,9")
	default:
		fmt.Println("Done!")
	}
}