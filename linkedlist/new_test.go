package linkedlist

import "testing"

func TestNew(t *testing.T) {
	list := New(1)

	if length := list.Length(); length != 1 {
		t.Errorf("should have 1, got %d", length)
	}

	if list.head != 1 {
		t.Errorf("should be 1, got %v", list.head)
	}

	if list.tail != nil {
		t.Errorf("should be <nil>, got %v", list.tail)
	}
}
