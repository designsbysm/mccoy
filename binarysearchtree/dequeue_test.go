package binarysearchtree

import "testing"

func TestDequeue(t *testing.T) {
	tree := createTree()

	node := tree.Dequeue()
	if node.value != -1 {
		t.Errorf("should be -1, got %d", node.value)
	}

	node = tree.Dequeue()
	if node.value != 0 {
		t.Errorf("should be 0, got %d", node.value)
	}

	node = tree.Dequeue()
	if node.value != 1 {
		t.Errorf("should be 1, got %d", node.value)
	}
}

func TestDequeueEmpty(t *testing.T) {
	tree := New[int]()

	node := tree.Dequeue()
	if node != nil {
		t.Errorf("should be nil, got %v", node)
	}
}
