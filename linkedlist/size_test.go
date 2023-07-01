package linkedlist

import (
	"testing"
)

func TestSize(t *testing.T) {
	list := createList()

	if size := list.Size(); size != 9 {
		t.Errorf("should have 9, got %d", size)
	}
}
