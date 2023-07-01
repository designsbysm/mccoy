package linkedlist

import (
	"testing"
)

func TestIndex(t *testing.T) {
	list := createList()
	item := list.Index(5)

	if item.head != 6 {
		t.Errorf("should be 6, got %d", item.head)
	}
}
