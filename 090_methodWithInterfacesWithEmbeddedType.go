package main

import (
	"fmt"
)

// Create a struct type
type User struct {
	Name  string
	Email string
}

type Admin struct {
	User
	Level string
}

func (u *User) Notify() {
	fmt.Printf("Name: %q, Email: %q\n", u.Name, u.Email)
}

type Notifier interface {
	Notify()
}

func SendNotification(n Notifier) {
	n.Notify()
}

func main() {
	a := Admin{
		User: User{
			Name:  "kamal",
			Email: "ksr",
		},
		Level: "admin",
	}
	SendNotification(&a)

}
