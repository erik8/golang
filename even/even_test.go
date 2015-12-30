package even

import (
	"fmt"
	"testing"
)

// Example functions serving as documentation and tests
func ExampleEven() {
	if Even(2) {
		fmt.Printf("Is Even\n")
	}
	// Output:
	// Is Even
}
func TestEven(t *testing.T) {
	if Even(2) {
		t.Log("2 should be odd!")
		t.FailNow()
	}
	if Even(4) {
		t.Log("4 should be odd!")
		t.FailNow()
	}
}

func Test2Even(t *testing.T) {
	if Even(2) {
		t.Log("6 should be odd!")
		t.Fail()
	}
	if Even(4) {
		t.Log("8 should be odd!")
		t.Fail()
	}
}
