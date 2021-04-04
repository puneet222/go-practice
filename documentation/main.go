package main

import (
	"fmt"
	"github.com/puneet222/go-practice/documentation/dog"
)

type canine struct {
	name string
	age int
}

func main() {
	f := canine{
		name: "Fido",
		age: dog.Years(10),
	}
	fmt.Println("-> ", f)
}
