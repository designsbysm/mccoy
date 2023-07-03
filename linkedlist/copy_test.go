package linkedlist

import (
	"testing"
)

func TestCopy(t *testing.T) {
	list1 := createList()
	list2 := list1.Copy()

	index := 3
	item1 := list1.Index(index)
	item2 := list2.Index(index)

	if item1.head != item2.head {
		t.Errorf("should have %d, got %d", item1.head, item2.head)
	}
}
