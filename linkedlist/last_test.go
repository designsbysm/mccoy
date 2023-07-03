package linkedlist

import "testing"

func TestLast(t *testing.T) {
	list := createList()
	item := list.Last()

	if item.head != 9 {
		t.Errorf("should be 9, got %v", item.head)
	}
}
