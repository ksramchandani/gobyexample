package main

import (
	"fmt"
)

// Create a struct type
type User struct {
	Name  string
	Email string
}

// Create a named type
type Mystring string

func (u *User) Notify() {
	fmt.Printf("Name: %q, Email: %q\n", u.Name, u.Email)
}

func (m Mystring) Notify() {
	fmt.Printf("String: %q\n", m)
}

type Notifier interface {
	Notify()
}

func SendNotification(n Notifier) {
	n.Notify()
}

func main() {
	u := User{
		Name:  "Kamal",
		Email: "ksr",
	}
	SendNotification(&u)

	m := Mystring("abc")
	SendNotification(m)

}
