package main

import (
	"fmt"
	"log"
)

type customErr struct {
	x int
	y int
}

func (ce customErr) Error() string {
	return fmt.Sprintf("Error at coordinates %v,%v", ce.x, ce.y)
}

func main() {
	err := printN()
	if err != nil {
		log.Println(err)
	}
	fmt.Println("Hello, playground")
}

func printN() error {
	for i := 0; ; i++ {
		if i > 3 {
			return customErr{i, 99}
		}
		fmt.Println("printing... ", i)
	}
}
