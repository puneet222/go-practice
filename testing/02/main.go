package main

import (
	"fmt"
	"github.com/puneet222/go-practice/testing/02/quote"
	"github.com/puneet222/go-practice/testing/02/word"
)

func main() {
	fmt.Println(word.Count(quote.SunAlso))

	for k, v := range word.UseCount(quote.SunAlso) {
		fmt.Println(v, k)
	}
}
