package linkedlist

import "testing"

func TestLength(t *testing.T) {
	list := createList()

	if length := list.Length(); length != 9 {
		t.Errorf("should have 9, got %d", length)
	}
}
