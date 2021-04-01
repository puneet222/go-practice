package main

import "fmt"

func main() {
	cs := make(chan<- int) // send only channel
	// cr := make(<-chan int) // receive only channel

	go func() {
		cs <- 42
	}()

	//fmt.Println(<- cs) // send only channel
	//fmt.Println(<-cr)  // this will cause a deadlock as we are waiting on receive channel and
						 // no one is sending on this channel (which is not possible as it is receive only)

	fmt.Println("---------")
	fmt.Printf("cs\t%T\n", cs)
}
