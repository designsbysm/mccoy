package linkedlist

import (
	"testing"
)

func TestDelete(t *testing.T) {
	list := createList()
	list.Delete(4)

	item := list.Index(4)
	if item.head != 6 {
		t.Errorf("should be 6, got %d", item.head)
	}

	if size := list.Length(); size != 8 {
		t.Errorf("should have 8, got %d", size)
	}
}
