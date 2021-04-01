package main

import "fmt"

func main() {
	q := make(chan int)
	c := gen2(q)

	receive2(c, q)
	fmt.Println("EXIT")
}

func gen2(q chan<- int) <-chan int {
	c := make(chan int)

	go func() {
		for i := 0; i < 100; i++ {
			c <- i
		}
		q <- 1
	}()

	return c
}

func receive2(c <- chan int, q <- chan int) {
	for {
		select {
			case v := <-c :
				fmt.Println("value received: ", v)

			case <- q:
				fmt.Println("QUIT: ")
				return
		}
	}
}
