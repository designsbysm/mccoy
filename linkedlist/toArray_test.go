package linkedlist

import "testing"

func TestToArray(t *testing.T) {
	array := []int{31, 13, 12, 4, 18, 16, 7, 2, 3, 0, 10}
	list := FromArray(array)
	array = list.ToArray()

	idx := 3
	item := list.Index(idx)

	if array[idx] != item.head {
		t.Errorf("should have %d, got %d", array[idx], item.head)
	}
}
