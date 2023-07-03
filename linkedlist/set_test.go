package linkedlist

import "testing"

func TestSet(t *testing.T) {
	list := createList()
	list.Set(4, 99)

	item := list.Index(4)
	if item.head != 99 {
		t.Errorf("should be 99, got %d", item.head)
	}
}
