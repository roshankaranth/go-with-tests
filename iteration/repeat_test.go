package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	repeated := Repeat(10, "a")
	expected := "aaaaaaaaaa"

	if repeated != expected {
		t.Errorf("got %q , expected %q", repeated, expected)
	}
}

func ExampleRepeat() {
	repeated := Repeat(5, "a")
	fmt.Println(repeated)
	//Output: aaaaa
}

// benchmarks tests the code for b.N time. The value of b.N is picked by the compiler to give some result and figure out the op/sec.
func Benchmark(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat(10, "a")
	}
}
