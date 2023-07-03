package linkedlist

import (
	"testing"
)

func TestFilter(t *testing.T) {
	list := createList()
	filtered := list.Filter(func(i int) bool {
		return i <= 5
	})

	if size := filtered.Length(); size != 5 {
		t.Errorf("should have 5, got %d", size)
	}
}

func TestFilterEmpty(t *testing.T) {
	var list *Node[int]
	filtered := list.Filter(func(i int) bool {
		return i <= 5
	})

	if length := filtered.Length(); length != 0 {
		t.Errorf("should have 0, got %d", length)
	}
}
