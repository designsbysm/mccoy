package linkedlist

import (
	"testing"
)

func TestCons(t *testing.T) {
	list := New(1)
	list = list.Cons(2)
	list = list.Cons(3)

	if length := list.Length(); length != 3 {
		t.Errorf("should have 3, got %d", length)
	}

	item := list.head
	if item != 3 {
		t.Errorf("should have 3, got %d", item)
	}
}

func TestConsEmpty(t *testing.T) {
	var list *Node[int]
	list = list.Cons(1)
	list = list.Cons(2)
	list = list.Cons(3)

	if length := list.Length(); length != 3 {
		t.Errorf("should have 3, got %d", length)
	}

	item := list.head
	if item != 3 {
		t.Errorf("should have 3, got %d", item)
	}
}
