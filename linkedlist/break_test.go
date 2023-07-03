package linkedlist

import "testing"

func TestBreak(t *testing.T) {
	list := createList()
	left, right := list.Break(4)

	if length := list.Length(); length != 9 {
		t.Errorf("should have 9, got %d", length)
	}

	if length := left.Length(); length != 5 {
		t.Errorf("should have 5, got %d", length)
	}

	if length := right.Length(); length != 4 {
		t.Errorf("should have 4, got %d", length)
	}
}
