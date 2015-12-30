package stack

import "testing"

func TestPushPop(t *testing.T) {
	mys := new(Stack)
	mys.Push(5)
	if mys.Pop() != 5 {
		t.Log("Pop does not give 5")
		t.Fail()
	}
}
