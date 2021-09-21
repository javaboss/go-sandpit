package saying

import (
	"fmt"
	"testing"
)

// use go test ./... to test everything below that level
// use go test -bench . at the required level to run benchmark tests

func TestGreet(t *testing.T) {
	expected := "Welcome my dear James"

	s := Greet("James")

	if s != expected {
		t.Error("got", s, "expected", expected)
	}
}

func ExampleGreet() {
	fmt.Println(Greet("James"))
	// Output:
	// Welcome my dear James
}

func BenchmarkGreet(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Greet("James")
	}
}
