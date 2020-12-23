package main

import "fmt"

func main() {


	name := "John"
	lastName := "Doe"

	sayHello(&name, &lastName)


}

func sayHello(name, lastName *string)  {

	*name = "robert"

	fmt.Println(*name, *lastName)
}
