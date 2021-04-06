package mymath

import (
	"fmt"
	"testing"
)

type data struct {
	xi []int
	result float64
}

func gen() []data {
	a := data{[]int{1, 4, 6, 8, 100}, 6}
	b := data{[]int{0, 8, 10, 1000}, 9}
	c := data{[]int{9000, 4, 10, 8, 6, 12}, 9}
	d := data{[]int{123, 744, 140, 200}, 170}
	dt := []data{a,b,c,d}
	return dt
}

func BenchmarkCenteredAvg(b *testing.B) {
	xi := []int{9000, 4, 10, 8, 6, 12}
	for i := 0; i < b.N; i++ {
		CenteredAvg(xi)
	}
}

func ExampleCenteredAvg() {
	xi := []int{1, 4, 6, 8, 100}
	r := CenteredAvg(xi)
	fmt.Println(r)
	// Output: 6
}

func TestCenteredAvg(t *testing.T) {
	dt := gen()
	for _, v := range dt {
		if CenteredAvg(v.xi) != v.result {
			t.Errorf("Expected %v but got %v", v.result, CenteredAvg(v.xi))
		}
	}
}


