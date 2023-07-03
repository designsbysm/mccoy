package linkedlist

import "testing"

func TestInsert(t *testing.T) {
	list := createList()
	list.Insert(4, 99)
	item := list.Index(5)

	if item.head != 99 {
		t.Errorf("should be 99, got %v", item.head)
	}
}
