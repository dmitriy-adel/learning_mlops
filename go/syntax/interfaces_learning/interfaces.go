package main

import (
	"fmt"
)

type Speaker interface {
	Speak() string
}

type Dog struct{}

func (d Dog) Speak() string {
	return "Woof! ept"
}

type Person struct{}

func (p Person) Speak() string {
	return "Say hello to my ..."
}

func SaySomething(s Speaker) {
	fmt.Println(s.Speak())
}

func main() {

	SaySomething(Dog{})
	SaySomething(Person{})
}
