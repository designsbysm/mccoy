package binarysearchtree

import "testing"

func TestNew(t *testing.T) {
	tree := New[string]()

	if tree.root != nil {
		t.Errorf("should be nil, got %v", tree.root)
	}
}
