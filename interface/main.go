package main

import (
	"fmt"
	"log"
	"reflect"
)

type notifier interface {
	notify()
}

type user struct {
	name  string
	email string
}

type test struct {
	A interface{}
}

func (u *user) notify() {
	fmt.Printf("Sending user email to %s<%s>\n",
		u.name,
		u.email)
}

func main() {
	//u := user{}
	// this will cause compiled error
	//var _ notifier = u
	//u.notify()

	//var a any = test{A: nil}
	//
	//if v, ok := a.(test); ok {
	//	fmt.Printf("v is test : %v\n", v)
	//} else {
	//	fmt.Printf("v is not test : %v\n", v)
	//}

	var a *user = nil
	var b notifier = nil

	accept(a)
	accept(b)
	accept(nil)
}

func accept(b notifier) {
	if !reflect.ValueOf(b).IsValid() || reflect.ValueOf(b).IsNil() {
		log.Printf("b is nil")
	}
}
