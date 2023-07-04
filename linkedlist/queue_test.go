package linkedlist

import "testing"

func TestQueue(t *testing.T) {
	list := New(1)
	list.Enqueue(2)
	list.Enqueue(3)
	list.Enqueue(4)
	item, list := list.Dequeue()

	if length := list.Length(); length != 3 {
		t.Errorf("should have 3, got %d", length)
	}

	if item != 1 {
		t.Errorf("should have 1, got %d", item)
	}

	got := toString(list)
	should := "234"
	if should != got {
		t.Errorf("should have %s, got %s", should, got)
	}
}
