package linkedlist

import (
	"testing"
)

func TestUncons(t *testing.T) {
	list := createList()
	item, list := list.Uncons()

	if size := list.Length(); size != 8 {
		t.Errorf("should be 8, got %d", size)
	}

	item = item.(int)
	if item != 1 {
		t.Errorf("should have 1, got %d", item)
	}

	item = list.head.(int)
	if item != 2 {
		t.Errorf("should have 2, got %d", item)
	}
}
