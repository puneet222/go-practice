package dog

import (
	"fmt"
	"testing"
)

// BET Benchmarking Exampling Testing

func BenchmarkYears(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Years(10)
	}
}

func ExampleYears() {
	dy := Years(10)
	fmt.Println(dy)
	// Output: 70
}

func TestYears(t *testing.T) {
	dy := Years(10)
	if dy != 70 {
		t.Error("Expected 70 got", dy)
	}
}

func BenchmarkYearsTwo(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		YearsTwo(10)
	}
}

func ExampleYearsTwo() {
	dy := YearsTwo(6)
	fmt.Println(dy)
	// Output: 42
}