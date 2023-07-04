package linkedlist

import "testing"

func TestPeek(t *testing.T) {
	list := createList()
	item := list.Peek()

	if item != 1 {
		t.Errorf("should be 1, got %v", item)
	}
}
