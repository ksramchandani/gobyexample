package main

import "fmt"

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

func (a *Admin) Notify() {
	fmt.Printf("Name: %q, Email: %q, Level: %q\n", a.User.Name, a.User.Email, a.Level)
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
			Name:  "Kamal",
			Email: "ksr",
		},
		Level: "admin",
	}

	SendNotification(&a)
	SendNotification(&a.User)

}
