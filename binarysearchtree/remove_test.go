package binarysearchtree

import (
	"strconv"
	"testing"
)

func TestRemove(t *testing.T) {
	tree := createTree()
	tree.Remove(6)

	var got string
	tree.Traverse(func(n *Node[int]) {
		got += strconv.Itoa(n.value)
	})

	should := "-1012381011"
	if got != should {
		t.Errorf("should be %s, got %s", should, got)
	}

	tree = New[int]()
	for i := 0; i < 10; i++ {
		tree.Enqueue(i, i)
	}
	tree.Remove(8)
	tree.Remove(7)

	tree = New[int]()
	for i := 9; i > 0; i-- {
		tree.Enqueue(i, i)
	}
	tree.Remove(8)
	tree.Remove(7)

	tree = New[int]()
	for i := 0; i < 100; i++ {
		tree.Enqueue(i, i)
	}
	for i := 99; i > 0; i-- {
		tree.Enqueue(i, i)
	}
	tree.Remove(8)
	tree.Remove(7)
}

func TestRemoveEmpty(t *testing.T) {
	tree := New[int]()
	tree.Remove(6)

	// should not throw error
}
