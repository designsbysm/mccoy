package binarysearchtree

import "testing"

func TestMin(t *testing.T) {
	tree := createTree()
	value := tree.Min()

	if value == nil || *value != -1 {
		t.Errorf("should be -1, got %d", *value)
	}
}

func TestMinEmpty(t *testing.T) {
	tree := New[int]()
	value := tree.Min()

	if value != nil {
		t.Errorf("should be nil, got %d", value)
	}
}
