package linkedlist

import (
	"testing"
)

func TestInsert(t *testing.T) {
	list := createList()
	list.Insert(4, "e")
	item := list.Index(5)

	if item.head != "e" {
		t.Errorf("should be e, got %v", item.head)
	}
}
