package linkedlist

import "testing"

func TestEmpty(t *testing.T) {
	list := Empty[int]()
	list = list.Cons(1)
	list = list.Cons(2)
	list = list.Cons(3)

	if length := list.Length(); length != 3 {
		t.Errorf("should have 3, got %d", length)
	}
}
