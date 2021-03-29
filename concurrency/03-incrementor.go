package main

import (
	"fmt"
	"runtime"
	"sync"
)



func main() {
	counter := 0

	fmt.Println("CPUs", runtime.NumCPU())
	fmt.Println("Goroutines:", runtime.NumGoroutine())

	var wg sync.WaitGroup
	limit := 100
	wg.Add(limit)

	for i := 0; i < limit; i++ {
		go func() {
			v := counter
			v++
			runtime.Gosched()
			counter = v
			wg.Done()
		}()
		fmt.Println("Goroutines:", runtime.NumGoroutine())
	}
	fmt.Println("Goroutines:", runtime.NumGoroutine())
	wg.Wait()
	fmt.Println("Counter: ", counter)
}
