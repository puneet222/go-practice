package main

import "fmt"

type person struct {
	name string
	age int
}

func (p *person) speak()  {
	fmt.Println("Bla bla bla")
}

type human interface {
	speak()
}

func saySomething(h human) {
	h.speak()
}

func main() {
	p := person{
		name: "Puneet",
		age: 23,
	}

	p.speak()

	saySomething(&p)
}
