package word

import (
	"fmt"
	"github.com/puneet222/go-practice/testing/02/quote"
	"reflect"
	"testing"
)

func BenchmarkUseCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		UseCount(quote.SunAlso)
	}
}

func BenchmarkCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Count(quote.SunAlso)
	}
}

func ExampleCount() {
	str := "this is it, you are doing great"
	fmt.Println(Count(str))
	// Output: 7
}

func TestUseCount(t *testing.T) {
	str := "abc bbc kkk lop bbc jij kkk"
	rm := map[string]int{
		"abc": 1,
		"bbc": 2,
		"kkk": 2,
		"lop": 1,
		"jij": 1,
	}
	m := UseCount(str)
	eq := reflect.DeepEqual(m, rm)
	if !eq {
		t.Error("Map is not valid, Expected", rm, "Got", m)
	}
}

