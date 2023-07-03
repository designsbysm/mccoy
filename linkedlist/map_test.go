package linkedlist

import (
	"testing"
)

func TestMap(t *testing.T) {
	list := createList()
	mapped := list.Map(func(i int) int {
		return i + 1
	})

	before := list.Index(4)
	if before.head != 5 {
		t.Errorf("should have 5, got %d", before.head)
	}

	after := mapped.Index(4)
	if after.head != 6 {
		t.Errorf("should have 6, got %d", after.head)
	}
}

func TestMapEmpty(t *testing.T) {
	var list *Node[int]
	mapped := list.Map(func(i int) int {
		return i + 1
	})

	if length := mapped.Length(); length != 0 {
		t.Errorf("should have 0, got %d", length)
	}
}
