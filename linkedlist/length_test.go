package linkedlist

import (
	"testing"
)

func TestLength(t *testing.T) {
	list := createList()

	if size := list.Length(); size != 9 {
		t.Errorf("should have 9, got %d", size)
	}
}
