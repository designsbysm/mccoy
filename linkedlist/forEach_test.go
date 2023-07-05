package linkedlist

import (
	"strconv"
	"testing"
)

func TestForEach(t *testing.T) {
	list := createList()

	var got string
	list.ForEach(func(node *Node[int]) {
		if node.tail == nil {
			return
		}

		got += strconv.Itoa(node.head)
	})

	if got != "12345678" {
		t.Errorf("should have 12345678, got %s", got)
	}
}
