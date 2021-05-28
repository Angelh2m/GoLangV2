package main

type Incrementer interface {
	Increment() int
}

type IntCounter int

func (ic *IntCounter) Increment()int  {
	*ic++
	return int(*ic)
}

func main() {
	
	//myInt := IntCounter(0)
	//myInt
	//
	//var ic Incrementer = &myInt
	//
	//for i := 0 {
	//
	//}
	
}