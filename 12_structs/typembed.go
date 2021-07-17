package main

import "fmt"

type user struct {
	name  string
	email string
}

func (u user) notify() {
	fmt.Printf("Sending user email to %s<%s>\n", u.name, u.email)

}

type admin struct {
	user
	level string
}

func main() {
	ad := admin{
		user: user{
			name:  "foo bar",
			email: "foo@bar.com",
		},
		level: "admin",
	}

	ad.user.notify()

	ad.notify()
}
