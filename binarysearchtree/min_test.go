package binarysearchtree

import "testing"

func TestMin(t *testing.T) {
	tree := createTree()
	node := tree.Min()

	if node == nil || node.value != -1 {
		t.Errorf("should be -1, got %v", node.value)
	}
}

func TestMinEmpty(t *testing.T) {
	tree := New[int]()
	node := tree.Min()

	if node != nil {
		t.Errorf("should be nil, got %v", node)
	}
}
