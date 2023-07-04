package linkedlist

import "testing"

func TestStack(t *testing.T) {
	list := New(1)
	list = list.Push(2)
	list = list.Push(3)
	list = list.Push(4)
	item, list := list.Pop()

	if length := list.Length(); length != 3 {
		t.Errorf("should have 3, got %d", length)
	}

	if item != 4 {
		t.Errorf("should have 4, got %d", item)
	}

	got := toString(list)
	should := "321"
	if should != got {
		t.Errorf("should have %s, got %s", should, got)
	}
}
