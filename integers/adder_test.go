package integers

import (
	"fmt"
	"testing"
)

func TestAdder(t *testing.T) {
	sum := Add(2, 2)
	expected := 4

	if sum != expected {
		// %d is for printing integers
		t.Errorf("expected '%d' but got '%d'", expected, sum)
	}
}

/*
	Testable Examples

Will be compiled during a build, so that if ever your code changes so that the example is no longer valid, your build will fail.

Adding this code will cause the example to appear in your godoc documentation, making your code even more accessible.

//Output: 6 line indicates that the example will be executed
*/
func ExampleAdd() {
	sum := Add(1, 5)
	fmt.Println(sum)
	// Output: 6
}
