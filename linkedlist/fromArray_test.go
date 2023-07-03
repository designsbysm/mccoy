package linkedlist

import (
	"testing"

	"github.com/designsbysm/mccoy"
)

func TestFromArray(t *testing.T) {
	array := []mccoy.Item{31, 13, 12, 4, 18, 16, 7, 2, 3, 0, 10}
	list := FromArray(array)

	idx := 3
	item := list.Index(idx)

	if array[idx] != item.head {
		t.Errorf("should have %d, got %d", array[idx], item.head)
	}
}
