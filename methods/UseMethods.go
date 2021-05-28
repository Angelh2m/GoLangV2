package main

import "fmt"

type User struct {
	ID int
	FirstName string
	LastName string
}

var(
	users []*User
	nextID = 1
)

func getUsers() []*User {
	return users
}

func AddUser(u User) (User, error)  {
	u.ID = nextID
	nextID ++
	users = append(users, &u )
	return u, nil
}


func main() {

	u := User{1,"One", "Two"}
	AddUser(u)

	all := getUsers()

	for _, user := range all  {
		fmt.Println(user)

	}

	//fmt.Println(a)

}
