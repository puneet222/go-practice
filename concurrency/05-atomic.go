package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

func main() {
	var counter int64

	fmt.Println("CPUs", runtime.NumCPU())
	fmt.Println("Goroutines:", runtime.NumGoroutine())

	var wg sync.WaitGroup
	limit := 100
	wg.Add(limit)

	for i := 0; i < limit; i++ {
		go func() {
			atomic.AddInt64(&counter, 1)
			wg.Done()
		}()
		fmt.Println("count: ", atomic.LoadInt64(&counter))
		fmt.Println("Goroutines:", runtime.NumGoroutine())
	}
	fmt.Println("Goroutines:", runtime.NumGoroutine())
	wg.Wait()
	fmt.Println("Counter: ", counter)
}
