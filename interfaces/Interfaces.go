package main

import "fmt"

type Writer interface {
	Write([]byte) (int, error)
}

type ConsoleWriter struct {

}

func (cw ConsoleWriter)Write(data []byte) (int, error) {

	s, err := fmt.Println(string(data))
	return s, err
}

func main() {

	var w Writer = ConsoleWriter{}

	w.Write( []byte("Hello Go!") )

}