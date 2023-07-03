package linkedlist

import (
	"testing"
)

func TestEmpty(t *testing.T) {
	list := Empty()

	if size := list.Length(); size != 1 {
		t.Errorf("should have 1, got %d", size)
	}

	if list.head != nil {
		t.Errorf("should be <nil>, got %v", list.head)
	}

	if list.tail != nil {
		t.Errorf("should be <nil>, got %v", list.tail)
	}
}
