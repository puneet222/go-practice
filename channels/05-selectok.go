package main

import "fmt"

func main() {
	c := gen3()

	receive3(c)
	fmt.Println("EXIT")
}

func gen3() <-chan int {
	c := make(chan int)

	go func() {
		for i := 0; i < 100; i++ {
			c <- i
		}
		close(c)
	}()

	return c
}

func receive3(c <- chan int) {
	for {
		select {
			case v, ok := <-c :
				if !ok {
					fmt.Println("Closed: ", ok)
					return
				}
				fmt.Println("value received: ", v)
		}
	}
}
