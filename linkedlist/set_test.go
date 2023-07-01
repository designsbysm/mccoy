package linkedlist

import (
	"testing"
)

func TestSet(t *testing.T) {
	list := createList()
	list.Set(4, "e")

	item := list.Index(4)
	if item.head != "e" {
		t.Errorf("should be e, got %d", item.head)
	}
}
