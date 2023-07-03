package linkedlist

import "testing"

func TestCons(t *testing.T) {
	list := New(1)
	if size := list.Length(); size != 1 {
		t.Errorf("should be 1, got %d", size)
	}

	list = list.Cons(2)
	if size := list.Length(); size != 2 {
		t.Errorf("should have 2, got %d", size)
	}

	list = list.Cons(3)
	if size := list.Length(); size != 3 {
		t.Errorf("should have 3, got %d", size)
	}

	item := list.head.(int)
	if item != 3 {
		t.Errorf("should have 3, got %d", item)
	}
}
