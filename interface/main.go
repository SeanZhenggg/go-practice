package main

import "fmt"

type notifier interface {
	notify()
}

type user struct {
	name  string
	email string
}

func (u *user) notify() {
	fmt.Printf("Sending user email to %s<%s>\n",
		u.name,
		u.email)
}

func main() {
	u := user{}
	// this will cause compiled error
	//var _ notifier = u
	u.notify()
}
