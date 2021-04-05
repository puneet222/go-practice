package main

import "fmt"

func main() {
	fmt.Println("4*7*3 = ", multiply(4,7,3))
	fmt.Println("2*3 = ", multiply(2,3))
}

func multiply(xi ...int) int {
	pr := 1
	for _, v := range xi {
		pr *= v
	}
	return pr
}
