package binarysearchtree

import (
	"strconv"
	"testing"
)

func TestTraverse(t *testing.T) {
	tree := createTree()

	var got string
	tree.Traverse(func(n *Node[int]) {
		got += strconv.Itoa(n.value)
	})

	should := "-10123681011"
	if got != should {
		t.Errorf("should be %s, got %s", should, got)
	}
}
