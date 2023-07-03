package linkedlist

import "testing"

func TestAppend(t *testing.T) {
	list := New(1)
	list.Append(2)
	list.Append(3)

	if size := list.Length(); size != 3 {
		t.Errorf("should have 3, got %d", size)
	}

	item := list.Last().head
	if item != 3 {
		t.Errorf("should have 3, got %d", item)
	}
}
