package main

import "fmt"

func main() {

	r, err := devide(2.5, 3.0)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(r)
}

func devide(a, b float64)(float64, error)  {

	if b == 0{
		return 0.0, fmt.Errorf("Cannot devide by zero")
	}

	return a / b, nil
}
