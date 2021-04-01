package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	var xc []<-chan int
	rt := 100
	for i := 0; i < rt; i++ {
		xc = append(xc, put2(i*100))
	}
	f := fanIn(xc)
	fmt.Println("Go routines:", runtime.NumGoroutine())
	for v := range f {
		fmt.Println("Value got: ", v)
	}
	fmt.Println("EXIT")
}

func fanIn(ch []<-chan int) <- chan int{
	var wg sync.WaitGroup
	f := make(chan int)
	wg.Add(len(ch))
	for _, c := range ch {
		go func(t <-chan int) {
			for v := range t {
				f <- v
			}
			wg.Done()
		}(c)
	}
	go func() {
		wg.Wait()
		close(f)
	}()
	return f
}

func put2(v int) <-chan int {
	c := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			c <- i + v
		}
		close(c)
	}()
	return c
}
