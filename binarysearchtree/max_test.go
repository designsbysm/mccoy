package binarysearchtree

import "testing"

func TestMax(t *testing.T) {
	tree := createTree()
	value := tree.Max()

	if value == nil || *value != 11 {
		t.Errorf("should be 11, got %d", *value)
	}
}

func TestMaxEmpty(t *testing.T) {
	tree := New[int]()
	value := tree.Max()

	if value != nil {
		t.Errorf("should be nil, got %d", value)
	}
}
