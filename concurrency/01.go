package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func main()  {
	fmt.Println("CPUs", runtime.NumCPU())
	fmt.Println("Goroutines", runtime.NumGoroutine())

	wg.Add(2)

	go foo()
	go bar()

	fmt.Println("Goroutines", runtime.NumGoroutine())

	wg.Wait()
}

func foo() {
	fmt.Println("Foo")
	wg.Done()
}

func bar() {
	fmt.Println("Bar")
	wg.Done()
}
