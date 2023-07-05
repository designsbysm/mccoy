package binarysearchtree

import "testing"

func TestMax(t *testing.T) {
	tree := createTree()
	node := tree.Max()

	if node == nil || node.value != 11 {
		t.Errorf("should be 11, got %v", node.value)
	}
}

func TestMaxEmpty(t *testing.T) {
	tree := New[int]()
	node := tree.Max()

	if node != nil {
		t.Errorf("should be nil, got %v", node)
	}
}
