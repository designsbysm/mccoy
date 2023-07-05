package binarysearchtree

import "testing"

func TestPeek(t *testing.T) {
	tree := createTree()
	node := tree.Peek()

	if node == nil || node.value != -1 {
		t.Errorf("should be -1, got %v", node.value)
	}
}
