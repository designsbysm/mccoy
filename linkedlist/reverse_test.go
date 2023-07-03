package linkedlist

import "testing"

func TestReverse(t *testing.T) {
	list := createList()
	reversed := list.Reverse()

	last := reversed.Last()
	if last.head != 1 {
		t.Errorf("should have 6, got %d", last.head)
	}
}
