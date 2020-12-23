package main

import (
	"fmt"
)

func main() {
	a := []int{1,2,3,4,5,6,7,8,9}
	b := a[:] // Slice of all elements
	c := a[3:] // Slice from 4th element to end
	d := a[:6] // Slice first 6 elements
	e := a[3:6] // Slice 4,5,6 elements

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(e)

	//Takes the type + len + capacity
	s := make([]int, 3, 100)

	s = append(s, a...)

	fmt.Println(s)


}