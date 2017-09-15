package main

import "fmt"

type Hello struct {
	Name string
}

func (hello *Hello) hello() {
	fmt.Println("Hello " + hello.Name + "!")
}