package main

import "fmt"

func main() {
	c := make(chan int)
	go put(c)
	get(c)
	fmt.Println("Exit")
}

func put(c chan<- int) {
	for i := 0; i < 100; i++ {
		c <- i
	}
	close(c)
}

func get(c <-chan int) {
	for v := range c {
		fmt.Println("Value received: ", v)
	}
}