package gigasecond

import "testing"

func Add(a int, b int) int {
	return a + b
}

func TestHelloWorld(t *testing.T) {
	sum := Add(2, 3)
	if sum != 5 {
		t.Fail()
	}
}
